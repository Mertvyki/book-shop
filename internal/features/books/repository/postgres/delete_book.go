package book_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (r *BooksRepository) DeleteBook(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	DELETE FROM bookshop.books
	WHERE id=$1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("book with id=%d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
