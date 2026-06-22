package books_transport_http

import (
	"mime/multipart"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	core_http_types "github.com/Mertvyki/book-shop/internal/core/transport/http/types"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
)

type PatchBookRequest struct {
	Title         core_http_types.Nullable[string]  `json:"title"`
	Description   core_http_types.Nullable[string]  `json:"description"`
	ISBN          core_http_types.Nullable[string]  `json:"isbn"`
	Price         core_http_types.Nullable[float64] `json:"price"`
	BookType      core_http_types.Nullable[string]  `json:"book_type"`
	StockQuantity core_http_types.Nullable[int]     `json:"stock_quantity"`
	PublisherID   core_http_types.Nullable[int]     `json:"publisher_id"`
	AuthorIDs     []int                             `json:"author_ids"`
	CategoryIDs   []int                             `json:"category_ids"`
}

func (r PatchBookRequest) ToService() books_service.PatchBookPayload {
	payload := books_service.PatchBookPayload{
		Title:       r.Title.ToDomain().Value,
		Description: r.Description.ToDomain().Value,
		ISBN:        r.ISBN.ToDomain().Value,
		Price:       r.Price.ToDomain().Value,
		BookType:    r.BookType.ToDomain().Value,
		AuthorIDs:   r.AuthorIDs,
		CategoryIDs: r.CategoryIDs,
	}

	pub := r.PublisherID.ToDomain()
	if pub.Set {
		payload.PublisherID = pub.Value
	}

	stock := r.StockQuantity.ToDomain()
	if stock.Set {
		payload.StockQuantity = stock.Value
	}

	return payload
}

func (h *BooksHTTPHandler) PatchBook(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	bookID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bookID path value")
		return
	}

	err = r.ParseMultipartForm(32 << 20)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to parse multipart form")
		return
	}

	var request PatchBookRequest
	err = core_http_request.DecodeAndValidateMultipartJSONField(r, "request", &request)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	var coverFile multipart.File
	var coverHeader *multipart.FileHeader

	coverFile, coverHeader, err = r.FormFile("cover_image")
	if err != nil {
		coverFile = nil
		coverHeader = nil
	}

	var bookFile multipart.File
	var bookHeader *multipart.FileHeader

	bookFile, bookHeader, err = r.FormFile("book_file")
	if err != nil {
		bookFile = nil
		bookHeader = nil
	}

	book, err := h.booksService.PatchBook(
		ctx,
		bookID,
		request.ToService(),
		coverFile,
		coverHeader,
		bookFile,
		bookHeader,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch book")
		return
	}

	response := GetBookResponse(bookDTOFromDomain(book))
	responseHandler.JSONResponse(response, http.StatusOK)
}
