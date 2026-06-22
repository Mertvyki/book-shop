package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) GetPublishers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	publishers, err := h.booksService.ListPublishers(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to list publishers")
		return
	}

	response := make([]PublisherDTOResponse, len(publishers))
	for i, p := range publishers {
		response[i] = PublisherDTOResponse{
			ID:   p.ID,
			Name: p.Name,
		}
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
