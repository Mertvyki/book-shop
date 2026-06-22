package books_transport_http

import (
	"fmt"
	"net/http"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

const maxCategoryNameLen = 100

type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
}

func (h *BooksHTTPHandler) CreateCategory(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var req CreateCategoryRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	if req.Name == "" {
		responseHandler.ErrorResponse(fmt.Errorf("name is required"), "name is required")
		return
	}

	if len(req.Name) > maxCategoryNameLen {
		responseHandler.ErrorResponse(fmt.Errorf("name too long (max %d characters): %w", maxCategoryNameLen, core_errors.ErrInvalidArgument), "name too long")
		return
	}

	category, err := h.booksService.CreateCategory(ctx, req.Name, req.Slug, req.Description)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create category")
		return
	}

	responseHandler.JSONResponse(CategoryDTOResponse{
		ID:          category.ID,
		Name:        category.Name,
		Slug:        category.Slug,
		Description: category.Description,
	}, http.StatusOK)
}
