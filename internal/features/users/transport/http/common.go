package users_transport_http

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
)

func checkOwnerOrAdmin(r *http.Request, pathUserID int) error {
	jwtUserID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		return fmt.Errorf("user id not found in context")
	}

	role, ok := r.Context().Value(core_http_middleware.UserRoleKey).(string)
	if !ok {
		return fmt.Errorf("user role not found in context")
	}

	if jwtUserID != pathUserID && role != "admin" {
		return fmt.Errorf("insufficient permissions")
	}

	return nil
}
