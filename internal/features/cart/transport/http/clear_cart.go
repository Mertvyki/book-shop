package cart_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

var unauthorizedErr = fmt.Errorf("unauthorized")

func (h *CartHTTPHandler) ClearCart(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	if err := h.cartService.ClearCart(ctx, userID); err != nil {
		responseHandler.ErrorResponse(err, "failed to clear cart")
		return
	}

	responseHandler.NoContentResponse()
}
