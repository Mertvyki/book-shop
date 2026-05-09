package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_config "github.com/Mertvyki/book-shop/internal/core/config"
	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_pgx_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool/pgx"
	core_security "github.com/Mertvyki/book-shop/internal/core/security"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	auth_postgres_repository "github.com/Mertvyki/book-shop/internal/features/auth/repository/postgres"
	auth_service "github.com/Mertvyki/book-shop/internal/features/auth/service"
	auth_transport_http "github.com/Mertvyki/book-shop/internal/features/auth/transport/http"
	users_postgres_repository "github.com/Mertvyki/book-shop/internal/features/users/repository/postgres"
	user_service "github.com/Mertvyki/book-shop/internal/features/users/service"
	users_transport_http "github.com/Mertvyki/book-shop/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	cfg := core_config.NewConfigMust()
	jwtCfg := core_security.NewJWTConfigMust()
	time.Local = cfg.TimeZone
	hasher := core_security.NewBcryptHasher(12)
	tokenManager := core_security.NewJWTManager(jwtCfg.Secret, jwtCfg.AccessTokenExpiry, jwtCfg.Issuer)
	refreshToken := core_security.NewRefreshTokenService()
	authMiddleware := core_http_middleware.Authenticate(tokenManager)
	adminMiddleware := core_http_middleware.RequireRole("admin")

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

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouterV1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRouters(usersTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRouters(authTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
