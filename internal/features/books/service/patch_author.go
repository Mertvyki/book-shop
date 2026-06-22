package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) PatchAuthor(ctx context.Context, id int, name *string, bio *string, birthYear *int) (domain.Author, error) {
	author, err := s.booksRepository.PatchAuthor(ctx, id, name, bio, birthYear)
	if err != nil {
		return domain.Author{}, fmt.Errorf("patch author: %w", err)
	}

	return author, nil
}
