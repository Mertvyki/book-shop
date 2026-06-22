package orders_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *OrdersHTTPHandler) GetOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	orderID, err := core_http_request.GetIntPathValue(r, "orderId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get orderId path value")
		return
	}

	result, err := h.orderService.GetOrder(ctx, userID, orderID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get order")
		return
	}

	deliveryCost := 0.0
	for _, item := range result.Items {
		if item.ItemType == "physical" {
			deliveryCost = 250.0
			break
		}
	}
	response := orderWithItemsDTO(result.Order, result.Items, deliveryCost)
	responseHandler.JSONResponse(response, http.StatusOK)
}
