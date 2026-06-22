package books_transport_http

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	books_service "github.com/Mertvyki/book-shop/internal/features/books/service"
	orders_service "github.com/Mertvyki/book-shop/internal/features/orders/service"
)

type BooksHTTPHandler struct {
	booksService   BooksService
	ordersService  *orders_service.OrderService
	authMiddleware  core_http_middleware.Middleware
	adminMiddleware core_http_middleware.Middleware
}

type BooksService interface {
	CreateBook(
		ctx context.Context,
		payload books_service.CreateBookPayload,
		coverFile multipart.File,
		coverHeader *multipart.FileHeader,
		bookFile multipart.File,
		bookHeader *multipart.FileHeader,
	) (domain.Book, error)

	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, queryParams books_service.GetBooksQueryParams) (books_service.GetBooksResult, error)
	DeleteBook(ctx context.Context, id int) error

	PatchBook(
		ctx context.Context,
		bookID int,
		payload books_service.PatchBookPayload,
		coverFile multipart.File,
		coverHeader *multipart.FileHeader,
		bookFile multipart.File,
		bookHeader *multipart.FileHeader,
	) (domain.Book, error)

	ListAuthors(ctx context.Context) ([]domain.Author, error)
	ListCategories(ctx context.Context) ([]domain.Category, error)
	ListPublishers(ctx context.Context) ([]domain.Publisher, error)

	CreateAuthor(ctx context.Context, name string, bio *string, birthYear *int) (domain.Author, error)
	PatchAuthor(ctx context.Context, id int, name *string, bio *string, birthYear *int) (domain.Author, error)
	DeleteAuthor(ctx context.Context, id int) error
	CreateCategory(ctx context.Context, name, slug string, description *string) (domain.Category, error)
	PatchCategory(ctx context.Context, id int, name *string, slug *string, description *string) (domain.Category, error)
	DeleteCategory(ctx context.Context, id int) error
	CreatePublisher(ctx context.Context, name string) (domain.Publisher, error)
	PatchPublisher(ctx context.Context, id int, name string) (domain.Publisher, error)
	DeletePublisher(ctx context.Context, id int) error
}

func NewBooksHTTPHandler(
	booksService BooksService,
	ordersService *orders_service.OrderService,
	authMiddleware core_http_middleware.Middleware,
	adminMiddleware core_http_middleware.Middleware,
) *BooksHTTPHandler {
	return &BooksHTTPHandler{
		booksService:   booksService,
		ordersService:  ordersService,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (h *BooksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/books",
			Handler: h.CreateBook,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
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
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/books/{id}",
			Handler: h.PatchBook,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/{id}/download",
			Handler: h.DownloadBook,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/{id}/purchased",
			Handler: h.CheckPurchased,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/authors",
			Handler: h.GetAuthors,
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories",
			Handler: h.GetCategories,
		},
		{
			Method:  http.MethodGet,
			Path:    "/publishers",
			Handler: h.GetPublishers,
		},
		{
			Method:  http.MethodPost,
			Path:    "/authors",
			Handler: h.CreateAuthor,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/authors/{id}",
			Handler: h.PatchAuthor,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/authors/{id}",
			Handler: h.DeleteAuthor,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: h.CreateCategory,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/categories/{id}",
			Handler: h.PatchCategory,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/categories/{id}",
			Handler: h.DeleteCategory,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/publishers",
			Handler: h.CreatePublisher,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodPatch,
			Path:    "/publishers/{id}",
			Handler: h.PatchPublisher,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
		{
			Method:  http.MethodDelete,
			Path:    "/publishers/{id}",
			Handler: h.DeletePublisher,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
				h.adminMiddleware,
			},
		},
	}
}
