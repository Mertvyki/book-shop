package book_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (r *BooksRepository) DeleteCategory(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tag, err := r.pool.Exec(ctx, `DELETE FROM bookshop.categories WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete category: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("category with id=%d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
