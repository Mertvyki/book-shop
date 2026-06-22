package books_service

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBooksRepo struct{ mock.Mock }

func (m *mockBooksRepo) CreateBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error) {
	args := m.Called(ctx, book, authorIDs, categoryIDs)
	return args.Get(0).(domain.Book), args.Error(1)
}

func (m *mockBooksRepo) GetBook(ctx context.Context, id int) (domain.Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Book), args.Error(1)
}

func (m *mockBooksRepo) GetBooks(ctx context.Context, qp GetBooksQueryParams) ([]domain.Book, int, error) {
	args := m.Called(ctx, qp)
	return args.Get(0).([]domain.Book), args.Int(1), args.Error(2)
}

func (m *mockBooksRepo) DeleteBook(ctx context.Context, id int) error {
	return m.Called(ctx, id).Error(0)
}

func (m *mockBooksRepo) PatchBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error) {
	args := m.Called(ctx, book, authorIDs, categoryIDs)
	return args.Get(0).(domain.Book), args.Error(1)
}

func (m *mockBooksRepo) ListAuthors(ctx context.Context) ([]domain.Author, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Author), args.Error(1)
}

func (m *mockBooksRepo) ListCategories(ctx context.Context) ([]domain.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Category), args.Error(1)
}

func (m *mockBooksRepo) ListPublishers(ctx context.Context) ([]domain.Publisher, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.Publisher), args.Error(1)
}

func (m *mockBooksRepo) GetPublisher(ctx context.Context, id int) (domain.Publisher, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Publisher), args.Error(1)
}

func (m *mockBooksRepo) GetAuthor(ctx context.Context, id int) (domain.Author, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Author), args.Error(1)
}

func (m *mockBooksRepo) GetCategory(ctx context.Context, id int) (domain.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Category), args.Error(1)
}

func (m *mockBooksRepo) CreateAuthor(ctx context.Context, name string, bio *string, birthYear *int) (domain.Author, error) {
	args := m.Called(ctx, name, bio, birthYear)
	return args.Get(0).(domain.Author), args.Error(1)
}

func (m *mockBooksRepo) PatchAuthor(ctx context.Context, id int, name *string, bio *string, birthYear *int) (domain.Author, error) {
	args := m.Called(ctx, id, name, bio, birthYear)
	return args.Get(0).(domain.Author), args.Error(1)
}

func (m *mockBooksRepo) DeleteAuthor(ctx context.Context, id int) error {
	return m.Called(ctx, id).Error(0)
}

func (m *mockBooksRepo) CreateCategory(ctx context.Context, name, slug string, description *string) (domain.Category, error) {
	args := m.Called(ctx, name, slug, description)
	return args.Get(0).(domain.Category), args.Error(1)
}

func (m *mockBooksRepo) PatchCategory(ctx context.Context, id int, name *string, slug *string, description *string) (domain.Category, error) {
	args := m.Called(ctx, id, name, slug, description)
	return args.Get(0).(domain.Category), args.Error(1)
}

func (m *mockBooksRepo) DeleteCategory(ctx context.Context, id int) error {
	return m.Called(ctx, id).Error(0)
}

func (m *mockBooksRepo) CreatePublisher(ctx context.Context, name string) (domain.Publisher, error) {
	args := m.Called(ctx, name)
	return args.Get(0).(domain.Publisher), args.Error(1)
}

func (m *mockBooksRepo) PatchPublisher(ctx context.Context, id int, name string) (domain.Publisher, error) {
	args := m.Called(ctx, id, name)
	return args.Get(0).(domain.Publisher), args.Error(1)
}

func (m *mockBooksRepo) DeletePublisher(ctx context.Context, id int) error {
	return m.Called(ctx, id).Error(0)
}

type mockFileStorage struct{ mock.Mock }

func (m *mockFileStorage) Upload(ctx context.Context, objectName string, reader io.Reader, size int64, contentType string) (string, error) {
	args := m.Called(ctx, objectName, reader, size, contentType)
	return args.String(0), args.Error(1)
}

func (m *mockFileStorage) DeleteObject(ctx context.Context, objectName string) error {
	return m.Called(ctx, objectName).Error(0)
}

func TestCreateBook_InvalidType(t *testing.T) {
	repo := &mockBooksRepo{}
	storage := &mockFileStorage{}
	svc := NewBookService(repo, storage)

	payload := CreateBookPayload{
		Title:    "Test Book",
		Price:    10.0,
		BookType: "invalid-type",
	}

	_, err := svc.CreateBook(context.Background(), payload, nil, nil, nil, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestCreateBook_PhysicalMissingStock(t *testing.T) {
	repo := &mockBooksRepo{}
	storage := &mockFileStorage{}
	svc := NewBookService(repo, storage)

	payload := CreateBookPayload{
		Title:    "Test Book",
		Price:    10.0,
		BookType: "physical",
	}

	_, err := svc.CreateBook(context.Background(), payload, nil, nil, nil, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestCreateBook_DigitalMissingFile(t *testing.T) {
	repo := &mockBooksRepo{}
	storage := &mockFileStorage{}
	svc := NewBookService(repo, storage)

	payload := CreateBookPayload{
		Title:    "Test Book",
		Price:    10.0,
		BookType: "digital",
	}

	_, err := svc.CreateBook(context.Background(), payload, nil, nil, nil, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}
