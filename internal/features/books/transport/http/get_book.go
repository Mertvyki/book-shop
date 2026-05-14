package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type GetBookResponse BookDTOResponse

func (h *BooksHTTPHandler) GetBook(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	bookID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get bookID path value",
		)

		return
	}

	book, err := h.booksService.GetBook(ctx, bookID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get book",
		)

		return
	}

	response := GetBookResponse(bookDTOFromDomain(book))
	responseHandler.JSONResponse(response, http.StatusOK)
}
