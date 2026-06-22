package orders_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *OrdersHTTPHandler) GetAllOrders(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to parse query params")
		return
	}

	result, err := h.orderService.GetAllOrders(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get all orders")
		return
	}

	dtoOrders := ordersDTOFromDomains(result.Orders)
	page := offset/limit + 1
	response := core_http_response.NewPaginatedResponse(dtoOrders, result.Total, page, limit)
	responseHandler.JSONResponse(response, http.StatusOK)
}
