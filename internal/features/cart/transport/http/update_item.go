package cart_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type UpdateItemResponse CartItemDTOResponse

type UpdateItemRequest struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

func (h *CartHTTPHandler) UpdateItem(rw http.ResponseWriter, r *http.Request) {
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

	var request UpdateItemRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	item, err := h.cartService.UpdateItem(ctx, userID, itemID, request.Quantity)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to update cart item")
		return
	}

	response := UpdateItemResponse(cartItemDTOFromDomain(item))
	responseHandler.JSONResponse(response, http.StatusOK)
}
