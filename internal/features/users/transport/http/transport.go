package users_transport_http

import (
	"context"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	user_service "github.com/Mertvyki/book-shop/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService    UsersService
	authMiddleware  core_http_middleware.Middleware
	adminMiddleware core_http_middleware.Middleware
}

type UsersService interface {
	CreateUser(
		ctx context.Context,
		email string,
		password string,
		fullName string,
		phoneNumber *string,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) (user_service.GetUsersResult, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	PatchUser(
		ctx context.Context,
		userID int,
		patch user_service.PatchUserPayload,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error
}

func NewUsersHTTPHandler(
	usersService UsersService,
	authMiddleware core_http_middleware.Middleware,
	adminMiddleware core_http_middleware.Middleware,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService:    usersService,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
	}
}
