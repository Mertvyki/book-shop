package cart_transport_http

import (
	"context"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
)

type CartHTTPHandler struct {
	cartService    CartService
	authMiddleware core_http_middleware.Middleware
}

type CartService interface {
	GetCart(ctx context.Context, userID int) ([]domain.CartItem, error)
	AddItem(ctx context.Context, userID, bookID, quantity int) (domain.CartItem, error)
	UpdateItem(ctx context.Context, userID, itemID, quantity int) (domain.CartItem, error)
	RemoveItem(ctx context.Context, userID, itemID int) error
	ClearCart(ctx context.Context, userID int) error
}

func NewCartHTTPHandler(cartService CartService, authMiddleware core_http_middleware.Middleware) *CartHTTPHandler {
	return &CartHTTPHandler{
		cartService:    cartService,
		authMiddleware: authMiddleware,
	}
}

func (h *CartHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/cart",
			Handler: h.GetCart,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/cart/items",
			Handler: h.AddItem,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/cart/items/{itemId}",
			Handler: h.UpdateItem,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/cart/items/{itemId}",
			Handler: h.RemoveItem,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/cart",
			Handler: h.ClearCart,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
	}
}
