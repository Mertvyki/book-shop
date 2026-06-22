package cart_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *CartHTTPHandler) RemoveItem(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	itemID, err := core_http_request.GetIntPathValue(r, "itemId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get itemId path value")
		return
	}

	if err := h.cartService.RemoveItem(ctx, userID, itemID); err != nil {
		responseHandler.ErrorResponse(err, "failed to remove cart item")
		return
	}

	responseHandler.NoContentResponse()
}
