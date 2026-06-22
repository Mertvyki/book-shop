package books_service

import (
	"context"
	"fmt"
)

func (s *BookService) DeleteCategory(ctx context.Context, id int) error {
	if err := s.booksRepository.DeleteCategory(ctx, id); err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	return nil
}
