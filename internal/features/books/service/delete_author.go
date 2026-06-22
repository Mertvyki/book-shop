package books_service

import (
	"context"
	"fmt"
)

func (s *BookService) DeleteAuthor(ctx context.Context, id int) error {
	if err := s.booksRepository.DeleteAuthor(ctx, id); err != nil {
		return fmt.Errorf("delete author: %w", err)
	}

	return nil
}
