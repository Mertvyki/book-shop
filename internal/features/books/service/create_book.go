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

func (s *BookService) CreateBook(
	ctx context.Context,
	request books_transport_http.CreateBookRequest,
	coverFile multipart.File,
	coverHeader *multipart.FileHeader,
	bookFile multipart.File,
	bookHeader *multipart.FileHeader,
) (domain.Book, error) {
	if request.BookType != "digital" &&
		request.BookType != "physical" {
		return domain.Book{}, fmt.Errorf(
			"invalid book type: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if request.BookType == "physical" &&
		request.StockQuantity == nil {
		return domain.Book{}, fmt.Errorf(
			"stock quantity required for physical book: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if request.BookType == "digital" &&
		bookHeader == nil {
		return domain.Book{}, fmt.Errorf(
			"book file required for digital book: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	bookUUID := uuid.NewString()

	coverExt := filepath.Ext(coverHeader.Filename)

	coverObjectKey := fmt.Sprintf(
		"books/%s/cover%s",
		bookUUID,
		coverExt,
	)

	uploadedCoverKey, err := s.fileStorage.Upload(
		ctx,
		coverObjectKey,
		coverFile,
		coverHeader.Size,
		coverHeader.Header.Get("Content-Type"),
	)
	if err != nil {
		return domain.Book{}, fmt.Errorf(
			"upload cover image: %w",
			err,
		)
	}

	var uploadedBookKey *string
	if request.BookType == "digital" {
		bookExt := filepath.Ext(bookHeader.Filename)

		bookObjectKey := fmt.Sprintf(
			"books/%s/book%s",
			bookUUID,
			strings.ToLower(bookExt),
		)

		uploadedKey, err := s.fileStorage.Upload(
			ctx,
			bookObjectKey,
			bookFile,
			bookHeader.Size,
			bookHeader.Header.Get("Content-Type"),
		)
		if err != nil {
			return domain.Book{}, fmt.Errorf(
				"upload book file: %w",
				err,
			)
		}

		uploadedBookKey = &uploadedKey
	}

	book := domain.NewBookUninitialized(
		request.Title,
		request.Author,
		request.Description,
		request.ISBN,
		request.Price,
		request.BookType,
		request.StockQuantity,
		uploadedBookKey,
		&uploadedCoverKey,
	)

	createdBook, err := s.booksRepository.CreateBook(ctx, book)
	if err != nil {
		return domain.Book{}, fmt.Errorf("create book in repository: %w", err)
	}

	return createdBook, nil
}
