package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	books_transport_http "github.com/Mertvyki/book-shop/internal/features/books/transport/http"
)

func (r *BooksRepository) GetBooks(
	ctx context.Context,
	queryParams books_transport_http.GetBooksQueryParams,
) ([]domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var query strings.Builder

	query.WriteString(`
	SELECT
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
	FROM bookshop.books
	WHERE 1=1
	`)

	args := []any{}
	argPos := 1

	if queryParams.Type != nil {
		query.WriteString(fmt.Sprintf(" AND book_type = $%d", argPos))
		args = append(args, *queryParams.Type)
		argPos++
	}

	if queryParams.Author != nil {
		query.WriteString(fmt.Sprintf(" AND author ILIKE $%d", argPos))
		args = append(args, "%"+*queryParams.Author+"%")
		argPos++
	}

	if queryParams.Search != nil {
		query.WriteString(fmt.Sprintf(`
		AND (
			title ILIKE $%d
			OR author ILIKE $%d
		)
		`, argPos, argPos))
		args = append(args, "%"+*queryParams.Search+"%")
		argPos++
	}

	if queryParams.MinPrice != nil {
		query.WriteString(
			fmt.Sprintf(
				" AND price >= $%d",
				argPos,
			),
		)

		args = append(args, *queryParams.MinPrice)

		argPos++
	}

	if queryParams.MaxPrice != nil {
		query.WriteString(
			fmt.Sprintf(
				" AND price <= $%d",
				argPos,
			),
		)

		args = append(args, *queryParams.MaxPrice)

		argPos++
	}

	offset := (queryParams.Page - 1) *
		queryParams.Limit

	query.WriteString(
		fmt.Sprintf(`
		ORDER BY created_at DESC
		LIMIT $%d
		OFFSET $%d
		`,
			argPos,
			argPos+1,
		),
	)

	args = append(
		args,
		queryParams.Limit,
		offset,
	)

	rows, err := r.pool.Query(
		ctx,
		query.String(),
		args...,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"query books: %w",
			err,
		)
	}
	defer rows.Close()

	books := make([]domain.Book, 0)

	for rows.Next() {
		var model BookModel

		err := rows.Scan(
			&model.ID,
			&model.Version,
			&model.Title,
			&model.Author,
			&model.Description,
			&model.ISBN,
			&model.Price,
			&model.BookType,
			&model.StockQuantity,
			&model.FileURL,
			&model.CoverImageURL,
			&model.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf(
				"scan book: %w",
				err,
			)
		}

		books = append(
			books,
			model.ToDomain(),
		)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"rows iteration: %w",
			err,
		)
	}

	return books, nil
}
