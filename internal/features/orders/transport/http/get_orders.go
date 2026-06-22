package orders_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

var unauthorizedErr = fmt.Errorf("unauthorized")

// @Summary List user orders
// @Description Get paginated list of current user's orders
// @Tags orders
// @Produce json
// @Param limit query int false "Items per page (default 20)"
// @Param offset query int false "Offset (default 0)"
// @Success 200 {object} core_http_response.PaginatedResponse
// @Router /orders [get]
// @Security bearerAuth
func (h *OrdersHTTPHandler) GetOrders(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to parse query params")
		return
	}

	result, err := h.orderService.GetOrders(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get orders")
		return
	}

	dtoOrders := ordersDTOFromDomains(result.Orders)
	page := offset/limit + 1
	response := core_http_response.NewPaginatedResponse(dtoOrders, result.Total, page, limit)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (int, int, error) {
	limitPtr, err := core_http_request.GetIntQueryParam(r, "limit")
	if err != nil {
		return 0, 0, fmt.Errorf("get limit query param: %w", err)
	}

	offsetPtr, err := core_http_request.GetIntQueryParam(r, "offset")
	if err != nil {
		return 0, 0, fmt.Errorf("get offset query param: %w", err)
	}

	limit := 20
	if limitPtr != nil {
		limit = *limitPtr
	}

	offset := 0
	if offsetPtr != nil {
		offset = *offsetPtr
	}

	if limit < 1 {
		limit = 1
	}

	if limit > 100 {
		limit = 100
	}

	if offset < 0 {
		offset = 0
	}

	return limit, offset, nil
}
