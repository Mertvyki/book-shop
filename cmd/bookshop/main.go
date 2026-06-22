//	Book Shop API
//
//	REST API for an online book store.
//
//	Schemes: http
//	Host: localhost:5050
//	BasePath: /api/v1
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//	- multipart/form-data
//
//	Produces:
//	- application/json
//
//	SecurityDefinitions:
//	bearerAuth:
//	  type: apiKey
//	  name: Authorization
//	  in: header
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Mertvyki/book-shop/internal/core/config"
	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_pgx_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool/pgx"
	core_security "github.com/Mertvyki/book-shop/internal/core/security"
	core_minio "github.com/Mertvyki/book-shop/internal/core/storage/minio"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	addresses_postgres_repository "github.com/Mertvyki/book-shop/internal/features/addresses/repository/postgres"
	addresses_service "github.com/Mertvyki/book-shop/internal/features/addresses/service"
	addresses_transport_http "github.com/Mertvyki/book-shop/internal/features/addresses/transport/http"
	auth_postgres_repository "github.com/Mertvyki/book-shop/internal/features/auth/repository/postgres"
	auth_service "github.com/Mertvyki/book-shop/internal/features/auth/service"
	auth_transport_http "github.com/Mertvyki/book-shop/internal/features/auth/transport/http"
	book_postgres_repository "github.com/Mertvyki/book-shop/internal/features/books/repository/postgres"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
	books_transport_http "github.com/Mertvyki/book-shop/internal/features/books/transport/http"
	cart_postgres_repository "github.com/Mertvyki/book-shop/internal/features/cart/repository/postgres"
	cart_service "github.com/Mertvyki/book-shop/internal/features/cart/service"
	cart_transport_http "github.com/Mertvyki/book-shop/internal/features/cart/transport/http"
	orders_postgres_repository "github.com/Mertvyki/book-shop/internal/features/orders/repository/postgres"
	orders_service "github.com/Mertvyki/book-shop/internal/features/orders/service"
	orders_transport_http "github.com/Mertvyki/book-shop/internal/features/orders/transport/http"
	reviews_postgres_repository "github.com/Mertvyki/book-shop/internal/features/reviews/repository/postgres"
	reviews_service "github.com/Mertvyki/book-shop/internal/features/reviews/service"
	reviews_transport_http "github.com/Mertvyki/book-shop/internal/features/reviews/transport/http"
	users_postgres_repository "github.com/Mertvyki/book-shop/internal/features/users/repository/postgres"
	user_service "github.com/Mertvyki/book-shop/internal/features/users/service"
	users_transport_http "github.com/Mertvyki/book-shop/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	jwtCfg := core_security.NewJWTConfigMust()
	minioCfg := core_minio.NewConfigMust()
	time.Local = cfg.TimeZone
	hasher := core_security.NewBcryptHasher(12)
	tokenManager := core_security.NewJWTManager(jwtCfg.Secret, jwtCfg.AccessTokenExpiry, jwtCfg.Issuer)
	refreshToken := core_security.NewRefreshTokenService()
	authMiddleware := core_http_middleware.Authenticate(tokenManager)
	adminMiddleware := core_http_middleware.RequireRole("admin")
	staffMiddleware := core_http_middleware.RequireRole("admin", "employee")

	minioClient, err := core_minio.NewClient(context.Background(), minioCfg)
	if err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("application time zone", zap.Any("zone", time.Local))

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := user_service.NewUserService(usersRepository, hasher, tokenManager)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService, authMiddleware, adminMiddleware)

	logger.Debug("initializing feature", zap.String("feature", "auth"))
	authRepository := auth_postgres_repository.NewRefreshTokenRepository(pool)
	authService := auth_service.NewAuthService(authRepository, usersRepository, hasher, tokenManager, refreshToken, 7*24*time.Hour)
	authTransportHTTP := auth_transport_http.NewAuthHTTPHandler(authService)

	logger.Debug("initializing feature", zap.String("feature", "books"))
	booksRepository := book_postgres_repository.NewBooksRepository(pool)

	logger.Debug("initializing feature", zap.String("feature", "addresses"))
	addressesRepository := addresses_postgres_repository.NewAddressesRepository(pool)
	addressesService := addresses_service.NewAddressesService(addressesRepository)
	addressesTransportHTTP := addresses_transport_http.NewAddressesHTTPHandler(addressesService, authMiddleware)

	logger.Debug("initializing feature", zap.String("feature", "cart"))
	cartRepository := cart_postgres_repository.NewCartRepository(pool)
	cartService := cart_service.NewCartService(cartRepository, booksRepository)
	cartTransportHTTP := cart_transport_http.NewCartHTTPHandler(cartService, authMiddleware)

	logger.Debug("initializing feature", zap.String("feature", "orders"))
	ordersRepository := orders_postgres_repository.NewOrdersRepository(pool)
	ordersService := orders_service.NewOrderService(ordersRepository, cartRepository, addressesRepository)
	ordersTransportHTTP := orders_transport_http.NewOrdersHTTPHandler(ordersService, authMiddleware, staffMiddleware)

	logger.Debug("initializing feature", zap.String("feature", "reviews"))
	reviewsRepository := reviews_postgres_repository.NewReviewsRepository(pool)
	reviewsService := reviews_service.NewReviewService(reviewsRepository)
	reviewsTransportHTTP := reviews_transport_http.NewReviewsHTTPHandler(reviewsService, authMiddleware)

	booksService := books_service.NewBookService(booksRepository, minioClient)
	booksHTTPHandler := books_transport_http.NewBooksHTTPHandler(booksService, ordersService, authMiddleware, staffMiddleware)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
		core_http_middleware.CORS(),
	)

	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRouters(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(authTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(booksHTTPHandler.Routes()...)
	apiVersionRouterV1.RegisterRouters(addressesTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(cartTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(ordersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(reviewsTransportHTTP.Routes()...)

	healthHandler := core_http_server.NewHealthHandler(pool)
	httpServer.RegisterHandler(http.MethodGet, "/health", healthHandler)

	httpServer.RegisterHandler(http.MethodGet, "/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/swagger.json")
	}))

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	httpServer.RegisterHandler("GET", "/files/{key...}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("key")
		if key == "" {
			http.Error(w, "missing file key", http.StatusBadRequest)
			return
		}

		presignedURL, err := minioClient.PresignedGetObject(r.Context(), key, 15*time.Minute)
		if err != nil {
			logger.Warn("failed to generate presigned URL", zap.String("key", key), zap.Error(err))
			http.Error(w, "file not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, presignedURL, http.StatusFound)
	}))

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
