package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) CreatePublisher(ctx context.Context, name string) (domain.Publisher, error) {
	publisher, err := s.booksRepository.CreatePublisher(ctx, name)
	if err != nil {
		return domain.Publisher{}, fmt.Errorf("create publisher: %w", err)
	}
	return publisher, nil
}

func (s *BookService) PatchPublisher(ctx context.Context, id int, name string) (domain.Publisher, error) {
	publisher, err := s.booksRepository.PatchPublisher(ctx, id, name)
	if err != nil {
		return domain.Publisher{}, fmt.Errorf("patch publisher: %w", err)
	}
	return publisher, nil
}

func (s *BookService) DeletePublisher(ctx context.Context, id int) error {
	if err := s.booksRepository.DeletePublisher(ctx, id); err != nil {
		return fmt.Errorf("delete publisher: %w", err)
	}
	return nil
}
