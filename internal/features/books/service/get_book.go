package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) GetBook(ctx context.Context, id int) (domain.Book, error) {
	book, err := s.booksRepository.GetBook(ctx, id)
	if err != nil {
		return domain.Book{}, fmt.Errorf("get book from repository: %w", err)
	}

	return book, nil
}
