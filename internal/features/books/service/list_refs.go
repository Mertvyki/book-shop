package books_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *BookService) ListAuthors(ctx context.Context) ([]domain.Author, error) {
	authors, err := s.booksRepository.ListAuthors(ctx)
	if err != nil {
		return nil, fmt.Errorf("list authors: %w", err)
	}

	return authors, nil
}

func (s *BookService) ListCategories(ctx context.Context) ([]domain.Category, error) {
	categories, err := s.booksRepository.ListCategories(ctx)
	if err != nil {
		return nil, fmt.Errorf("list categories: %w", err)
	}

	return categories, nil
}

func (s *BookService) ListPublishers(ctx context.Context) ([]domain.Publisher, error) {
	publishers, err := s.booksRepository.ListPublishers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list publishers: %w", err)
	}

	return publishers, nil
}
