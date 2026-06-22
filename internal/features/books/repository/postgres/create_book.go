package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) CreateBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Book{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var model BookModel
	err = tx.QueryRow(ctx, `
		INSERT INTO bookshop.books
			(title, description, isbn, price, book_type, stock_quantity, file_key, cover_image_key, publisher_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
		RETURNING id, version, title, description, isbn, price, book_type,
			stock_quantity, file_key, cover_image_key, publisher_id, created_at
	`, book.Title, book.Description, book.ISBN, book.Price, book.BookType,
		book.StockQuantity, book.FileURL, book.CoverImageURL, book.PublisherID,
	).Scan(
		&model.ID, &model.Version, &model.Title, &model.Description,
		&model.ISBN, &model.Price, &model.BookType, &model.StockQuantity,
		&model.FileURL, &model.CoverImageURL, &model.PublisherID, &model.CreatedAt,
	)
	if err != nil {
		return domain.Book{}, fmt.Errorf("insert book: %w", err)
	}

	if len(authorIDs) > 0 {
		query, args := buildMultiInsert("bookshop.book_authors", []string{"book_id", "author_id"}, model.ID, authorIDs)
		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return domain.Book{}, fmt.Errorf("insert book_authors: %w", err)
		}
	}

	if len(categoryIDs) > 0 {
		query, args := buildMultiInsert("bookshop.book_categories", []string{"book_id", "category_id"}, model.ID, categoryIDs)
		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			return domain.Book{}, fmt.Errorf("insert book_categories: %w", err)
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Book{}, fmt.Errorf("commit transaction: %w", err)
	}

	return model.ToDomain(), nil
}

func buildMultiInsert(table string, columns []string, constID int, ids []int) (string, []any) {
	colNames := make([]string, len(columns))
	for i, c := range columns {
		colNames[i] = table + "." + c
	}

	var sb strings.Builder
	sb.WriteString("INSERT INTO ")
	sb.WriteString(table)
	sb.WriteString(" (")
	for i, col := range columns {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(col)
	}
	sb.WriteString(") VALUES ")

	args := make([]any, 0, 1+len(ids))
	args = append(args, constID)

	for i, id := range ids {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("($1, $%d)", i+2))
		args = append(args, id)
	}

	return sb.String(), args
}
