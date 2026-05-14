package books_service

import (
	"context"
	"io"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	books_transport_http "github.com/Mertvyki/book-shop/internal/features/books/transport/http"
)

type BookService struct {
	booksRepository BooksRepository
	fileStorage     FileStorage
}

type BooksRepository interface {
	CreateBook(
		ctx context.Context,
		book domain.Book,
	) (domain.Book, error)

	GetBook(
		ctx context.Context,
		id int,
	) (domain.Book, error)

	GetBooks(
		ctx context.Context,
		queryParams books_transport_http.GetBooksQueryParams,
	) ([]domain.Book, error)

	DeleteBook(
		ctx context.Context,
		id int,
	) error

	PatchBook(
		ctx context.Context,
		book domain.Book,
	) (domain.Book, error)
}

type FileStorage interface {
	Upload(
		ctx context.Context,
		objectName string,
		reader io.Reader,
		size int64,
		contentType string,
	) (string, error)

	DeleteObject(
		ctx context.Context,
		objectName string,
	) error
}

func NewBookService(booksRepository BooksRepository, fileStorage FileStorage) *BookService {
	return &BookService{
		booksRepository: booksRepository,
		fileStorage:     fileStorage,
	}
}
