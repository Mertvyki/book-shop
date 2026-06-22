package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) PatchCategory(ctx context.Context, id int, name *string, slug *string, description *string) (domain.Category, error) {
	category, err := s.booksRepository.PatchCategory(ctx, id, name, slug, description)
	if err != nil {
		return domain.Category{}, fmt.Errorf("patch category: %w", err)
	}

	return category, nil
}
