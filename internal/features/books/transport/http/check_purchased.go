package books_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) CheckPurchased(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(fmt.Errorf("unauthorized"), "unauthorized")
		return
	}

	bookID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bookID path value")
		return
	}

	purchased, err := h.ordersService.HasUserPurchasedBook(ctx, userID, bookID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to check purchase")
		return
	}

	responseHandler.JSONResponse(map[string]bool{"purchased": purchased}, http.StatusOK)
}
