package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) CreateCategory(ctx context.Context, name, slug string, description *string) (domain.Category, error) {
	category, err := s.booksRepository.CreateCategory(ctx, name, slug, description)
	if err != nil {
		return domain.Category{}, fmt.Errorf("create category: %w", err)
	}

	return category, nil
}
