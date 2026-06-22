package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) GetAuthor(ctx context.Context, id int) (domain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m AuthorModel
	err := r.pool.QueryRow(ctx, `SELECT id, name, bio, birth_year, created_at FROM bookshop.authors WHERE id = $1`, id).Scan(
		&m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt,
	)
	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.Author{}, fmt.Errorf("author with id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.Author{}, fmt.Errorf("get author: %w", err)
	}

	return m.ToDomain(), nil
}
