package books_transport_http

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
)

type CreateBookRequest struct {
	Title         string
	Price         float64
	BookType      string
	Description   *string
	ISBN          *string
	StockQuantity *int
	PublisherID   *int
	AuthorIDs     []int
	CategoryIDs   []int
}

func (r CreateBookRequest) ToService() books_service.CreateBookPayload {
	return books_service.CreateBookPayload{
		Title:         r.Title,
		Description:   r.Description,
		ISBN:          r.ISBN,
		Price:         r.Price,
		BookType:      r.BookType,
		StockQuantity: r.StockQuantity,
		PublisherID:   r.PublisherID,
		AuthorIDs:     r.AuthorIDs,
		CategoryIDs:   r.CategoryIDs,
	}
}

// @Summary Create book
// @Description Add a new book (admin only, multipart form)
// @Tags books
// @Accept mpfd
// @Produce json
// @Param title formData string true "Book title"
// @Param price formData number true "Price"
// @Param book_type formData string true "Book type (physical/digital)"
// @Param description formData string false "Description"
// @Param isbn formData string false "ISBN"
// @Param stock_quantity formData int false "Stock quantity (required for physical)"
// @Param publisher_id formData int false "Publisher ID"
// @Param author_ids formData string false "Comma-separated author IDs"
// @Param category_ids formData string false "Comma-separated category IDs"
// @Param cover_image formData file true "Cover image file"
// @Param book_file formData file false "Book file (required for digital)"
// @Success 200 {object} BookDTOResponse
// @Failure 400 {object} map[string]string
// @Router /books [post]
// @Security bearerAuth
func (h *BooksHTTPHandler) CreateBook(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to parse multipart form")
		return
	}

	request, err := parseCreateBookRequest(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to parse request data")
		return
	}

	coverFile, coverHeader, err := r.FormFile("cover_image")
	if err != nil {
		responseHandler.ErrorResponse(err, "cover image is required")
		return
	}
	defer coverFile.Close()

	var bookFile multipart.File
	var bookHeader *multipart.FileHeader

	if request.BookType == "digital" {
		bookFile, bookHeader, err = r.FormFile("book_file")
		if err != nil {
			responseHandler.ErrorResponse(err, "book file is required for digital books")
			return
		}
		defer bookFile.Close()
	}

	createdBook, err := h.booksService.CreateBook(
		ctx,
		request.ToService(),
		coverFile,
		coverHeader,
		bookFile,
		bookHeader,
	)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create book")
		return
	}

	response := bookDTOFromDomain(createdBook)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func parseIntSlice(s string) ([]int, error) {
	if s == "" {
		return nil, nil
	}

	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		v, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, fmt.Errorf("invalid int in list: %w", err)
		}
		result = append(result, v)
	}

	return result, nil
}

func parseCreateBookRequest(r *http.Request) (CreateBookRequest, error) {
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		return CreateBookRequest{}, fmt.Errorf("invalid price: %w", err)
	}

	request := CreateBookRequest{
		Title:    r.FormValue("title"),
		BookType: r.FormValue("book_type"),
		Price:    price,
	}

	description := r.FormValue("description")
	if description != "" {
		request.Description = &description
	}

	isbn := r.FormValue("isbn")
	if isbn != "" {
		request.ISBN = &isbn
	}

	stockQuantity := r.FormValue("stock_quantity")
	if stockQuantity != "" {
		qty, err := strconv.Atoi(stockQuantity)
		if err != nil {
			return CreateBookRequest{}, fmt.Errorf("invalid stock quantity: %w", err)
		}

		request.StockQuantity = &qty
	}

	publisherID := r.FormValue("publisher_id")
	if publisherID != "" {
		pid, err := strconv.Atoi(publisherID)
		if err != nil {
			return CreateBookRequest{}, fmt.Errorf("invalid publisher_id: %w", err)
		}

		request.PublisherID = &pid
	}

	authorIDs := r.FormValue("author_ids")
	request.AuthorIDs, err = parseIntSlice(authorIDs)
	if err != nil {
		return CreateBookRequest{}, fmt.Errorf("invalid author_ids: %w", err)
	}

	categoryIDs := r.FormValue("category_ids")
	request.CategoryIDs, err = parseIntSlice(categoryIDs)
	if err != nil {
		return CreateBookRequest{}, fmt.Errorf("invalid category_ids: %w", err)
	}

	return request, nil
}
