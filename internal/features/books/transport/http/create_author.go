package books_transport_http

import (
	"fmt"
	"net/http"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

const maxAuthorNameLen = 100

type CreateAuthorRequest struct {
	Name      string  `json:"name"`
	Bio       *string `json:"bio"`
	BirthYear *int    `json:"birth_year"`
}

func (h *BooksHTTPHandler) CreateAuthor(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var req CreateAuthorRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	if req.Name == "" {
		responseHandler.ErrorResponse(fmt.Errorf("name is required"), "name is required")
		return
	}

	if len(req.Name) > maxAuthorNameLen {
		responseHandler.ErrorResponse(fmt.Errorf("name too long (max %d characters): %w", maxAuthorNameLen, core_errors.ErrInvalidArgument), "name too long")
		return
	}

	author, err := h.booksService.CreateAuthor(ctx, req.Name, req.Bio, req.BirthYear)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create author")
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
