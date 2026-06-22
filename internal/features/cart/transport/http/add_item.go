package cart_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type AddItemResponse CartItemDTOResponse

type AddItemRequest struct {
	BookID   int `json:"book_id" validate:"required,min=1"`
	Quantity int `json:"quantity" validate:"required,min=1"`
}

func (h *CartHTTPHandler) AddItem(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	var request AddItemRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	item, err := h.cartService.AddItem(ctx, userID, request.BookID, request.Quantity)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to add item to cart")
		return
	}

	response := AddItemResponse(cartItemDTOFromDomain(item))
	responseHandler.JSONResponse(response, http.StatusOK)
}
