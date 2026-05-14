package books_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type GetBooksResponse []BookDTOResponse

type GetBooksQueryParams struct {
	Type     *string
	Author   *string
	Search   *string
	MinPrice *float64
	MaxPrice *float64
	Page     int
	Limit    int
}

func (h *BooksHTTPHandler) GetBooks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	queryParams, err := getBooksQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"invalid query params",
		)

		return
	}

	books, err := h.booksService.GetBooks(ctx, queryParams)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get books",
		)

		return
	}

	response := GetBooksResponse(booksDTOFromDomains(books))
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getBooksQueryParams(r *http.Request) (GetBooksQueryParams, error) {
	bookType, err := core_http_request.GetStringQueryParam(r, "type")
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get type query param: %w", err)
	}
	author, err := core_http_request.GetStringQueryParam(r, "author")
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get author query param: %w", err)
	}
	search, err := core_http_request.GetStringQueryParam(r, "search")
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get search query param: %w", err)
	}
	minPrice, err := core_http_request.GetFloatQueryParam(r, "min_price")
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get min_price query param: %w", err)
	}
	maxPrice, err := core_http_request.GetFloatQueryParam(r, "max_price")
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get max_price query param: %w", err)
	}
	pagePtr, err := core_http_request.GetIntQueryParam(
		r,
		"page",
	)
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get page query param: %w", err)
	}

	limitPtr, err := core_http_request.GetIntQueryParam(
		r,
		"limit",
	)
	if err != nil {
		return GetBooksQueryParams{}, fmt.Errorf("get limit query param: %w", err)
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

	return GetBooksQueryParams{
		Type:     bookType,
		Author:   author,
		Search:   search,
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		Page:     page,
		Limit:    limit,
	}, nil
}
