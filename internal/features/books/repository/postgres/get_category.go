package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) GetCategory(ctx context.Context, id int) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m CategoryModel
	err := r.pool.QueryRow(ctx, `SELECT id, name, slug, description FROM bookshop.categories WHERE id = $1`, id).Scan(
		&m.ID, &m.Name, &m.Slug, &m.Description,
	)
	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.Category{}, fmt.Errorf("category with id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.Category{}, fmt.Errorf("get category: %w", err)
	}

	return m.ToDomain(), nil
}
