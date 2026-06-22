package books_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
)

// @Summary List books
// @Description Get a paginated list of books with filters
// @Tags books
// @Produce json
// @Param type query string false "Filter by book type (physical/digital)"
// @Param author_id query int false "Filter by author ID"
// @Param category_id query int false "Filter by category ID"
// @Param search query string false "Search by title or author name"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 20, max 100)"
// @Success 200 {object} core_http_response.PaginatedResponse
// @Router /books [get]
func (h *BooksHTTPHandler) GetBooks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	queryParams, err := getBooksQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "invalid query params")
		return
	}

	result, err := h.booksService.GetBooks(ctx, queryParams)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get books")
		return
	}

	dtoBooks := booksDTOFromDomains(result.Books)
	response := core_http_response.NewPaginatedResponse(
		dtoBooks, result.Total, queryParams.Page, queryParams.Limit,
	)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getBooksQueryParams(r *http.Request) (books_service.GetBooksQueryParams, error) {
	bookType, err := core_http_request.GetStringQueryParam(r, "type")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get type query param: %w", err)
	}

	authorID, err := core_http_request.GetIntQueryParam(r, "author_id")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get author_id query param: %w", err)
	}

	categoryID, err := core_http_request.GetIntQueryParam(r, "category_id")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get category_id query param: %w", err)
	}

	publisherID, err := core_http_request.GetIntQueryParam(r, "publisher_id")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get publisher_id query param: %w", err)
	}

	search, err := core_http_request.GetStringQueryParam(r, "search")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get search query param: %w", err)
	}

	sort, err := core_http_request.GetStringQueryParam(r, "sort")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get sort query param: %w", err)
	}

	minPrice, err := core_http_request.GetFloatQueryParam(r, "min_price")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get min_price query param: %w", err)
	}

	maxPrice, err := core_http_request.GetFloatQueryParam(r, "max_price")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get max_price query param: %w", err)
	}

	pagePtr, err := core_http_request.GetIntQueryParam(r, "page")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get page query param: %w", err)
	}

	limitPtr, err := core_http_request.GetIntQueryParam(r, "limit")
	if err != nil {
		return books_service.GetBooksQueryParams{}, fmt.Errorf("get limit query param: %w", err)
	}

	page := 1
	if pagePtr != nil {
		page = *pagePtr
	}

	limit := 20
	if limitPtr != nil {
		limit = *limitPtr
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 20
	}

	if limit > 100 {
		limit = 100
	}

	return books_service.GetBooksQueryParams{
		Type:        bookType,
		AuthorID:    authorID,
		CategoryID:  categoryID,
		PublisherID: publisherID,
		Search:      search,
		MinPrice:    minPrice,
		MaxPrice:    maxPrice,
		Sort:        sort,
		Page:        page,
		Limit:       limit,
	}, nil
}
