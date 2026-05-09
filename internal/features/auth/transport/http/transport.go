package auth_transport_http

import (
	"context"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
)

type AuthHTTPHandler struct {
	authService AuthService
}

type AuthService interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (string, string, domain.User, error)
	Refresh(
		ctx context.Context,
		oldRefreshToken string,
	) (string, string, error)
}

func NewAuthHTTPHandler(
	authService AuthService,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
	}
}

func (h *AuthHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/refresh",
			Handler: h.Refresh,
		},
	}
}
