package orders_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type PatchOrderStatusRequest struct {
	Status  string `json:"status" validate:"required"`
	Version int    `json:"version" validate:"required,min=1"`
}

func (h *OrdersHTTPHandler) PatchOrderStatus(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	orderID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get order ID")
		return
	}

	var req PatchOrderStatusRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "invalid request body")
		return
	}

	if err := h.orderService.PatchOrderStatus(ctx, orderID, req.Version, req.Status); err != nil {
		responseHandler.ErrorResponse(err, "failed to update order status")
		return
	}

	responseHandler.NoContentResponse()
}
