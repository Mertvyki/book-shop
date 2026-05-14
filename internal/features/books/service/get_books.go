package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	books_transport_http "github.com/Mertvyki/book-shop/internal/features/books/transport/http"
)

func (s *BookService) GetBooks(
	ctx context.Context,
	queryParams books_transport_http.GetBooksQueryParams,
) ([]domain.Book, error) {
	books, err := s.booksRepository.GetBooks(ctx, queryParams)
	if err != nil {
		return nil, fmt.Errorf("get books from repository: %w", err)
	}

	return books, nil
}
