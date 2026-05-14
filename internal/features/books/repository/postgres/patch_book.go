package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) PatchBook(
	ctx context.Context,
	book domain.Book,
) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE bookshop.books
	SET
		version = version + 1,
		title = $1,
		author = $2,
		description = $3,
		isbn = $4,
		price = $5,
		book_type = $6,
		stock_quantity = $7,
		file_key = $8,
		cover_image_key = $9
	WHERE id = $10
	AND version = $11
	RETURNING
		id,
		version,
		title,
		author,
		description,
		isbn,
		price,
		book_type,
		stock_quantity,
		file_key,
		cover_image_key,
		created_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		book.Title,
		book.Author,
		book.Description,
		book.ISBN,
		book.Price,
		book.BookType,
		book.StockQuantity,
		book.FileURL,
		book.CoverImageURL,
		book.ID,
		book.Version,
	)

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
		return domain.Book{}, fmt.Errorf(
			"patch book: %w",
			err,
		)
	}

	return bookModel.ToDomain(), nil
}
