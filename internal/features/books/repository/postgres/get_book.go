package book_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) GetBook(
	ctx context.Context,
	id int,
) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, title, author, description, isbn, price, book_type, stock_quantity, file_key, cover_image_key, created_at
	FROM bookshop.books
	WHERE id=$1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var bookModel BookModel

	err := row.Scan(
		&bookModel.ID,
		&bookModel.Version,
		&bookModel.Title,
		&bookModel.Author,
		&bookModel.Description,
		&bookModel.ISBN,
		&bookModel.Price,
		&bookModel.BookType,
		&bookModel.StockQuantity,
		&bookModel.FileURL,
		&bookModel.CoverImageURL,
		&bookModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Book{}, fmt.Errorf("book with id=%d: %w", id, core_errors.ErrNotFound)
		}

		return domain.Book{}, fmt.Errorf("scan error: %w", err)
	}

	bookDomain := domain.NewBook(
		bookModel.ID,
		bookModel.Version,
		bookModel.Title,
		bookModel.Author,
		bookModel.Description,
		bookModel.ISBN,
		bookModel.Price,
		bookModel.BookType,
		bookModel.StockQuantity,
		bookModel.FileURL,
		bookModel.CoverImageURL,
		bookModel.CreatedAt,
	)

	return bookDomain, nil
}
