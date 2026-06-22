package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) DeleteAuthor(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	authorID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get author id")
		return
	}

	if err := h.booksService.DeleteAuthor(ctx, authorID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete author")
		return
	}

	responseHandler.NoContentResponse()
}
