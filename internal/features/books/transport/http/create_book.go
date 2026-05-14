package books_transport_http

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type CreateBookRequest struct {
	Title         string
	Author        string
	Description   *string
	ISBN          *string
	Price         float64
	BookType      string
	StockQuantity *int
}

func (h *BooksHTTPHandler) CreateBook(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to pasre multipart form",
		)

		return
	}

	request, err := pareCreateBookRequest(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse request data",
		)

		return
	}

	coverFile, coverHeader, err := r.FormFile("cover_image")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"cover image is required",
		)

		return
	}
	defer coverFile.Close()

	var bookFile multipart.File
	var bookHeader *multipart.FileHeader

	if request.BookType == "digital" {
		bookFile, bookHeader, err = r.FormFile("book_file")
		if err != nil {
			responseHandler.ErrorResponse(
				err,
				"book file is required for digital books",
			)

			return
		}
		defer bookFile.Close()
	}

	createdBook, err := h.booksService.CreateBook(
		ctx,
		request,
		coverFile,
		coverHeader,
		bookFile,
		bookHeader,
	)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create book",
		)

		return
	}

	response := bookDTOFromDomain(createdBook)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func pareCreateBookRequest(r *http.Request) (CreateBookRequest, error) {
	price, err := strconv.ParseFloat(
		r.FormValue("price"),
		64,
	)
	if err != nil {
		return CreateBookRequest{}, fmt.Errorf("invalid price: %w", err)
	}

	request := CreateBookRequest{
		Title:    r.FormValue("title"),
		Author:   r.FormValue("author"),
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
			return CreateBookRequest{}, fmt.Errorf(
				"invalid stock quantity: %w",
				err,
			)
		}

		request.StockQuantity = &qty
	}

	return request, nil
}
