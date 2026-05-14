package books_service

import (
	"context"
	"fmt"
)

func (s *BookService) DeleteBook(ctx context.Context, id int) error {
	book, err := s.booksRepository.GetBook(ctx, id)
	if err != nil {
		return fmt.Errorf("get book by id: %w", err)
	}

	if err := s.booksRepository.DeleteBook(ctx, id); err != nil {
		return fmt.Errorf("delete book; %w", err)
	}

	if book.CoverImageURL != nil {
		err = s.fileStorage.DeleteObject(ctx, *book.CoverImageURL)
		if err != nil {
			return fmt.Errorf("failed to delete image: %w", err)
		}
	}

	if book.FileURL != nil {
		err = s.fileStorage.DeleteObject(ctx, *book.FileURL)
		if err != nil {
			return fmt.Errorf("failed to delete file: %w", err)
		}
	}

	return nil
}
