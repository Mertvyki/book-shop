package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) GetCategories(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	categories, err := h.booksService.ListCategories(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to list categories")
		return
	}

	response := make([]CategoryDTOResponse, len(categories))
	for i, c := range categories {
		response[i] = CategoryDTOResponse{
			ID:          c.ID,
			Name:        c.Name,
			Slug:        c.Slug,
			Description: c.Description,
		}
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
