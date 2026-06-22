package books_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type PatchAuthorRequest struct {
	Name      *string `json:"name"`
	Bio       *string `json:"bio"`
	BirthYear *int    `json:"birth_year"`
}

func (h *BooksHTTPHandler) PatchAuthor(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	authorID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get author id")
		return
	}

	var req PatchAuthorRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	author, err := h.booksService.PatchAuthor(ctx, authorID, req.Name, req.Bio, req.BirthYear)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch author")
		return
	}

	responseHandler.JSONResponse(AuthorDTOResponse{
		ID:        author.ID,
		Name:      author.Name,
		Bio:       author.Bio,
		BirthYear: author.BirthYear,
		CreatedAt: author.CreatedAt,
	}, http.StatusOK)
}
