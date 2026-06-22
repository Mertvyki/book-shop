package book_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) PatchBook(ctx context.Context, book domain.Book, authorIDs, categoryIDs []int) (domain.Book, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Book{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	var model BookModel
	err = tx.QueryRow(ctx, `
		UPDATE bookshop.books
		SET title = $1, description = $2, isbn = $3, price = $4, book_type = $5,
			stock_quantity = $6, file_key = $7, cover_image_key = $8,
			publisher_id = $9, version = version + 1
		WHERE id = $10 AND version = $11
		RETURNING id, version, title, description, isbn, price, book_type,
			stock_quantity, file_key, cover_image_key, publisher_id, created_at
	`, book.Title, book.Description, book.ISBN, book.Price, book.BookType,
		book.StockQuantity, book.FileURL, book.CoverImageURL, book.PublisherID,
		book.ID, book.Version,
	).Scan(
		&model.ID, &model.Version, &model.Title, &model.Description,
		&model.ISBN, &model.Price, &model.BookType, &model.StockQuantity,
		&model.FileURL, &model.CoverImageURL, &model.PublisherID, &model.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Book{}, fmt.Errorf("patch book with id=%d: %w", book.ID, core_errors.ErrNotFound)
		}

		return domain.Book{}, fmt.Errorf("update book: %w", err)
	}

	if authorIDs != nil {
		_, err = tx.Exec(ctx, `DELETE FROM bookshop.book_authors WHERE book_id = $1`, model.ID)
		if err != nil {
			return domain.Book{}, fmt.Errorf("delete book_authors: %w", err)
		}

		if len(authorIDs) > 0 {
			query, args := buildMultiInsert("bookshop.book_authors", []string{"book_id", "author_id"}, model.ID, authorIDs)
			_, err = tx.Exec(ctx, query, args...)
			if err != nil {
				return domain.Book{}, fmt.Errorf("insert book_authors: %w", err)
			}
		}
	}

	if categoryIDs != nil {
		_, err = tx.Exec(ctx, `DELETE FROM bookshop.book_categories WHERE book_id = $1`, model.ID)
		if err != nil {
			return domain.Book{}, fmt.Errorf("delete book_categories: %w", err)
		}

		if len(categoryIDs) > 0 {
			query, args := buildMultiInsert("bookshop.book_categories", []string{"book_id", "category_id"}, model.ID, categoryIDs)
			_, err = tx.Exec(ctx, query, args...)
			if err != nil {
				return domain.Book{}, fmt.Errorf("insert book_categories: %w", err)
			}
		}
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Book{}, fmt.Errorf("commit transaction: %w", err)
	}

	return model.ToDomain(), nil
}
