package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) CreateBook(ctx context.Context, book domain.Book) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO bookshop.books (
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
	)
	VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
		created_at;
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
		book.CreatedAt,
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
			"scan created book: %w",
			err,
		)
	}

	return bookModel.ToDomain(), nil
}
