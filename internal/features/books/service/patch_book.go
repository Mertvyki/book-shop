package books_service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	books_transport_http "github.com/Mertvyki/book-shop/internal/features/books/transport/http"
	"github.com/google/uuid"
)

func (s *BookService) PatchBook(
	ctx context.Context,
	bookID int,
	request books_transport_http.PatchBookRequest,
	coverFile multipart.File,
	coverHeader *multipart.FileHeader,
	bookFile multipart.File,
	bookHeader *multipart.FileHeader,
) (domain.Book, error) {
	book, err := s.booksRepository.GetBook(ctx, bookID)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	if request.Title.Set {
		if request.Title.Value == nil {
			return domain.Book{}, fmt.Errorf(
				"title cannot be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		book.Title = *request.Title.Value
	}

	if request.Author.Set {
		if request.Author.Value == nil {
			return domain.Book{}, fmt.Errorf(
				"author cannot be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		book.Author = *request.Author.Value
	}

	if request.Description.Set {
		if request.Description.Value == nil {
			book.Description = nil
		} else {
			book.Description = request.Description.Value
		}
	}

	if request.ISBN.Set {
		if request.ISBN.Value == nil {
			book.ISBN = nil
		} else {
			book.ISBN = request.ISBN.Value
		}
	}

	if request.Price.Set {
		if request.Price.Value == nil {
			return domain.Book{}, fmt.Errorf(
				"price cannot be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		book.Price = *request.Price.Value
	}

	if request.StockQuantity.Set {
		if request.StockQuantity.Value == nil {
			book.StockQuantity = nil
		} else {
			book.StockQuantity = request.StockQuantity.Value
		}
	}

	if book.BookType != "digital" && book.BookType != "physical" {
		return domain.Book{}, fmt.Errorf("invalid book type: %w", core_errors.ErrInvalidArgument)
	}

	if book.BookType == "digital" {
		book.StockQuantity = nil
	}

	if book.BookType == "physical" && book.StockQuantity == nil {
		return domain.Book{}, fmt.Errorf("stock required for physical book: %w", core_errors.ErrInvalidArgument)
	}

	oldCover := book.CoverImageURL
	oldFile := book.FileURL
	storageUUID := extractStorageUUID(*book.FileURL)

	if coverFile != nil && coverHeader != nil {
		objectKey := fmt.Sprintf("books/%s/cover%s", storageUUID, filepath.Ext(coverHeader.Filename))

		uploadedKey, err := s.fileStorage.Upload(ctx, objectKey, coverFile, coverHeader.Size, coverHeader.Header.Get("Content-Type"))
		if err != nil {
			return domain.Book{}, err
		}

		book.CoverImageURL = &uploadedKey
	}

	if bookFile != nil && bookHeader != nil {
		if book.BookType != "digital" {
			return domain.Book{}, fmt.Errorf("only digital books can have file: %w", core_errors.ErrInvalidArgument)
		}

		objectKey := fmt.Sprintf("books/%s/book%s", storageUUID, filepath.Ext(bookHeader.Filename))

		uploadedKey, err := s.fileStorage.Upload(ctx, objectKey, bookFile, bookHeader.Size, bookHeader.Header.Get("Content-Type"))
		if err != nil {
			return domain.Book{}, err
		}

		book.FileURL = &uploadedKey
	}

	updatedBook, err := s.booksRepository.PatchBook(ctx, book)
	if err != nil {
		return domain.Book{}, err
	}

	if oldCover != nil && book.CoverImageURL != nil && *oldCover != *book.CoverImageURL {
		_ = s.fileStorage.DeleteObject(ctx, *oldCover)
	}

	if oldFile != nil && book.FileURL != nil && *oldFile != *book.FileURL {
		_ = s.fileStorage.DeleteObject(ctx, *oldFile)
	}

	return updatedBook, nil
}

func extractStorageUUID(
	objectKey string,
) string {

	parts := strings.Split(objectKey, "/")

	if len(parts) < 2 {
		return uuid.NewString()
	}

	return parts[1]
}
