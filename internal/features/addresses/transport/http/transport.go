package addresses_transport_http

import (
	"context"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	addresses_service "github.com/Mertvyki/book-shop/internal/features/addresses/service"
)

type AddressesHTTPHandler struct {
	addressesService AddressesService
	authMiddleware   core_http_middleware.Middleware
}

type AddressesService interface {
	CreateAddress(ctx context.Context, userID int, request addresses_service.CreateAddressPayload) (domain.Address, error)
	GetAddresses(ctx context.Context, userID int) ([]domain.Address, error)
	GetAddress(ctx context.Context, userID int, addrID int) (domain.Address, error)
	PatchAddress(ctx context.Context, userID, addrID int, patch addresses_service.PatchAddressPayload) (domain.Address, error)
	DeleteAddress(ctx context.Context, userID, addrID int) error
}

func NewAddressesHTTPHandler(
	addressesService AddressesService,
	authMiddleware core_http_middleware.Middleware,
) *AddressesHTTPHandler {
	return &AddressesHTTPHandler{
		addressesService: addressesService,
		authMiddleware:   authMiddleware,
	}
}

func (h *AddressesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users/{userId}/addresses",
			Handler: h.CreateAddress,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{userId}/addresses",
			Handler: h.GetAddresses,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{userId}/addresses/{addrId}",
			Handler: h.GetAddress,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{userId}/addresses/{addrId}",
			Handler: h.PatchAddress,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{userId}/addresses/{addrId}",
			Handler: h.DeleteAddress,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
	}
}
