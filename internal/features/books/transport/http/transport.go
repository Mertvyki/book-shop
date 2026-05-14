package books_transport_http

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
)

type BooksHTTPHandler struct {
	booksService BooksService
}

type BooksService interface {
	CreateBook(
		ctx context.Context,
		request CreateBookRequest,
		coverFile multipart.File,
		coverHeader *multipart.FileHeader,
		bookFile multipart.File,
		bookHeader *multipart.FileHeader,
	) (domain.Book, error)

	GetBook(
		ctx context.Context,
		id int,
	) (domain.Book, error)

	GetBooks(
		ctx context.Context,
		queryParams GetBooksQueryParams,
	) ([]domain.Book, error)

	DeleteBook(
		ctx context.Context,
		id int,
	) error

	PatchBook(
		ctx context.Context,
		bookID int,
		request PatchBookRequest,
		coverFile multipart.File,
		coverHeader *multipart.FileHeader,
		bookFile multipart.File,
		bookHeader *multipart.FileHeader,
	) (domain.Book, error)
}

func NewBooksHTTPHandler(booksService BooksService) *BooksHTTPHandler {
	return &BooksHTTPHandler{
		booksService: booksService,
	}
}

func (h *BooksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/books",
			Handler: h.CreateBook,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/{id}",
			Handler: h.GetBook,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books",
			Handler: h.GetBooks,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/books/{id}",
			Handler: h.DeleteBook,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/books/{id}",
			Handler: h.PatchBook,
		},
	}
}
