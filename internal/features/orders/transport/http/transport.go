package orders_transport_http

import (
	"context"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	orders_postgres_repository "github.com/Mertvyki/book-shop/internal/features/orders/repository/postgres"
	orders_service "github.com/Mertvyki/book-shop/internal/features/orders/service"
)

type OrdersHTTPHandler struct {
	orderService    OrderService
	authMiddleware  core_http_middleware.Middleware
	adminMiddleware core_http_middleware.Middleware
}

type OrderService interface {
	CreateOrder(ctx context.Context, userID int, shippingAddressID *int) (domain.Order, []domain.OrderItem, float64, error)
	GetOrder(ctx context.Context, userID, orderID int) (orders_postgres_repository.OrderWithItems, error)
	GetOrders(ctx context.Context, userID int, limit, offset int) (orders_service.GetOrdersResult, error)
	GetAllOrders(ctx context.Context, limit, offset int) (orders_service.GetAllOrdersResult, error)
	PatchOrderStatus(ctx context.Context, orderID, version int, status string) error
}

func NewOrdersHTTPHandler(
	orderService OrderService,
	authMiddleware core_http_middleware.Middleware,
	adminMiddleware core_http_middleware.Middleware,
) *OrdersHTTPHandler {
	return &OrdersHTTPHandler{
		orderService:    orderService,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (h *OrdersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/orders",
			Handler: h.CreateOrder,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/orders/{orderId}",
			Handler: h.GetOrder,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/orders",
			Handler: h.GetOrders,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/orders/all",
			Handler: h.GetAllOrders,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/orders/{id}/status",
			Handler: h.PatchOrderStatus,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
	}
}
