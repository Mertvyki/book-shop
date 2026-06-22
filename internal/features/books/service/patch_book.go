package books_service

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	"github.com/google/uuid"
)

func (s *BookService) PatchBook(
	ctx context.Context,
	bookID int,
	payload PatchBookPayload,
	coverFile multipart.File,
	coverHeader *multipart.FileHeader,
	bookFile multipart.File,
	bookHeader *multipart.FileHeader,
) (domain.Book, error) {
	book, err := s.booksRepository.GetBook(ctx, bookID)
	if err != nil {
		return domain.Book{}, fmt.Errorf("failed to get book: %w", err)
	}

	if payload.Title != nil {
		book.Title = *payload.Title
	}

	if payload.Description != nil {
		if *payload.Description == "" {
			book.Description = nil
		} else {
			book.Description = payload.Description
		}
	}

	if payload.ISBN != nil {
		if *payload.ISBN == "" {
			book.ISBN = nil
		} else {
			book.ISBN = payload.ISBN
		}
	}

	if payload.Price != nil {
		book.Price = *payload.Price
	}

	if payload.StockQuantity != nil || (payload.StockQuantity == nil && payload.Price != nil) {
		book.StockQuantity = payload.StockQuantity
	}

	bookType := book.BookType
	if payload.BookType != nil {
		bookType = *payload.BookType
	}

	if bookType != "digital" && bookType != "physical" {
		return domain.Book{}, fmt.Errorf("invalid book type: %w", core_errors.ErrInvalidArgument)
	}

	if bookType == "digital" {
		book.StockQuantity = nil
	}

	if bookType == "physical" && book.StockQuantity == nil {
		return domain.Book{}, fmt.Errorf("stock required for physical book: %w", core_errors.ErrInvalidArgument)
	}

	book.BookType = bookType

	if payload.PublisherID != nil {
		book.PublisherID = payload.PublisherID
	}

	oldCover := book.CoverImageURL
	oldFile := book.FileURL

	storageUUID := extractStorageUUID(book.CoverImageURL)

	if coverFile != nil && coverHeader != nil {
		objectKey := fmt.Sprintf("books/%s/cover%s", storageUUID, filepath.Ext(coverHeader.Filename))

		uploadedKey, err := s.fileStorage.Upload(ctx, objectKey, coverFile, coverHeader.Size, coverHeader.Header.Get("Content-Type"))
		if err != nil {
			return domain.Book{}, err
		}

		book.CoverImageURL = &uploadedKey
	}

	if bookFile != nil && bookHeader != nil {
		if bookType != "digital" {
			return domain.Book{}, fmt.Errorf("only digital books can have file: %w", core_errors.ErrInvalidArgument)
		}

		objectKey := fmt.Sprintf("books/%s/book%s", storageUUID, filepath.Ext(bookHeader.Filename))

		uploadedKey, err := s.fileStorage.Upload(ctx, objectKey, bookFile, bookHeader.Size, bookHeader.Header.Get("Content-Type"))
		if err != nil {
			return domain.Book{}, err
		}

		book.FileURL = &uploadedKey
	}

	var authorIDs []int
	if payload.AuthorIDs != nil {
		authorIDs = payload.AuthorIDs
	} else {
		authorIDs = nil
	}

	var categoryIDs []int
	if payload.CategoryIDs != nil {
		categoryIDs = payload.CategoryIDs
	} else {
		categoryIDs = nil
	}

	updatedBook, err := s.booksRepository.PatchBook(ctx, book, authorIDs, categoryIDs)
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

func extractStorageUUID(objectKey *string) string {
	if objectKey == nil {
		return uuid.NewString()
	}

	parts := strings.Split(*objectKey, "/")
	if len(parts) < 2 {
		return uuid.NewString()
	}

	return parts[1]
}
