package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type GetBooksResult struct {
	Books []domain.Book
	Total int
}

func (s *BookService) GetBooks(ctx context.Context, queryParams GetBooksQueryParams) (GetBooksResult, error) {
	books, total, err := s.booksRepository.GetBooks(ctx, queryParams)
	if err != nil {
		return GetBooksResult{}, fmt.Errorf("get books from repository: %w", err)
	}

	return GetBooksResult{
		Books: books,
		Total: total,
	}, nil
}
