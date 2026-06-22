package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) CreateAuthor(ctx context.Context, name string, bio *string, birthYear *int) (domain.Author, error) {
	author, err := s.booksRepository.CreateAuthor(ctx, name, bio, birthYear)
	if err != nil {
		return domain.Author{}, fmt.Errorf("create author: %w", err)
	}

	return author, nil
}
