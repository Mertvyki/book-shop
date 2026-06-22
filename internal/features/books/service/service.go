package books_service

import (
	"context"
	"io"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type GetBooksQueryParams struct {
	Type        *string
	AuthorID    *int
	CategoryID  *int
	PublisherID *int
	Search      *string
	MinPrice    *float64
	MaxPrice    *float64
	Sort        *string
	Page        int
	Limit       int
}

type CreateBookPayload struct {
	Title         string
	Description   *string
	ISBN          *string
	Price         float64
	BookType      string
	StockQuantity *int
	PublisherID   *int
	AuthorIDs     []int
	CategoryIDs   []int
}

type PatchBookPayload struct {
	Title         *string
	Description   *string
	ISBN          *string
	Price         *float64
	BookType      *string
	StockQuantity *int
	PublisherID   *int
	AuthorIDs     []int
	CategoryIDs   []int
}

type BookService struct {
	booksRepository BooksRepository
	fileStorage     FileStorage
}

type BooksRepository interface {
	CreateBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error)
	GetBook(ctx context.Context, id int) (domain.Book, error)
	GetBooks(ctx context.Context, queryParams GetBooksQueryParams) ([]domain.Book, int, error)
	DeleteBook(ctx context.Context, id int) error
	PatchBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error)
	ListAuthors(ctx context.Context) ([]domain.Author, error)
	ListCategories(ctx context.Context) ([]domain.Category, error)
	ListPublishers(ctx context.Context) ([]domain.Publisher, error)
	GetPublisher(ctx context.Context, id int) (domain.Publisher, error)
	GetAuthor(ctx context.Context, id int) (domain.Author, error)
	GetCategory(ctx context.Context, id int) (domain.Category, error)
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

type FileStorage interface {
	Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error)
	DeleteObject(ctx context.Context, objectName string) error
}

func NewBookService(booksRepository BooksRepository, fileStorage FileStorage) *BookService {
	return &BookService{
		booksRepository: booksRepository,
		fileStorage:     fileStorage,
	}
}
