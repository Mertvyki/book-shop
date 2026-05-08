package core_http_middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_security "github.com/Mertvyki/book-shop/internal/core/security"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type contextKey string

const (
	UserIDKey   contextKey = "userID"
	UserRoleKey contextKey = "userRole"
)

func Authenticate(jwtManager *core_security.JWTManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log := core_logger.FromContext(r.Context())
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				responseHandler.ErrorResponse(fmt.Errorf("missing or invalid authorization header"), "unauthorized")
				return
			}
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := jwtManager.Parse(tokenString)
			if err != nil {
				responseHandler.ErrorResponse(fmt.Errorf("invalid token: %w", err), "unauthorized")
				return
			}

			userIDFloat, ok := claims["user_id"].(float64)
			if !ok {
				responseHandler.ErrorResponse(fmt.Errorf("invalid token claims"), "unauthorized")
				return
			}
			userID := int(userIDFloat)

			role, ok := claims["role"].(string)
			if !ok {
				responseHandler.ErrorResponse(fmt.Errorf("invalid token claims"), "unauthorized")
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, userID)
			ctx = context.WithValue(ctx, UserRoleKey, role)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func RequireRole(allowedRoles ...string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			role, ok := r.Context().Value(UserRoleKey).(string)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			for _, allowed := range allowedRoles {
				if role == allowed {
					next.ServeHTTP(w, r)
					return
				}
			}

			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"message":"forbidden","error":"insufficient permissions"}`))
		})
	}
}
