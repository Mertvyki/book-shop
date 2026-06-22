package orders_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type CreateOrderRequest struct {
	ShippingAddressID *int `json:"shipping_address_id"`
}

// @Summary Create order
// @Description Create an order from the current user's cart
// @Tags orders
// @Accept json
// @Produce json
// @Param request body CreateOrderRequest true "Order details"
// @Success 200 {object} OrderDTOResponse
// @Failure 400 {object} map[string]string
// @Router /orders [post]
// @Security bearerAuth
func (h *OrdersHTTPHandler) CreateOrder(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	var request CreateOrderRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	order, items, deliveryCost, err := h.orderService.CreateOrder(ctx, userID, request.ShippingAddressID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create order")
		return
	}

	response := orderWithItemsDTO(order, items, deliveryCost)
	responseHandler.JSONResponse(response, http.StatusOK)
}
