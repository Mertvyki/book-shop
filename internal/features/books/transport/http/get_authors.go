package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) GetAuthors(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	authors, err := h.booksService.ListAuthors(ctx)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to list authors")
		return
	}

	response := make([]AuthorDTOResponse, len(authors))
	for i, a := range authors {
		response[i] = AuthorDTOResponse{
			ID:        a.ID,
			Name:      a.Name,
			Bio:       a.Bio,
			BirthYear: a.BirthYear,
			CreatedAt: a.CreatedAt,
		}
	}

	responseHandler.JSONResponse(response, http.StatusOK)
}
