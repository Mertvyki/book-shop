package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *BooksRepository) CreatePublisher(ctx context.Context, name string) (domain.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m PublisherModel
	err := r.pool.QueryRow(ctx, `
		INSERT INTO bookshop.publishers (name) VALUES ($1)
		RETURNING id, name
	`, name).Scan(&m.ID, &m.Name)
	if err != nil {
		return domain.Publisher{}, fmt.Errorf("create publisher: %w", err)
	}

	return m.ToDomain(), nil
}

func (r *BooksRepository) PatchPublisher(ctx context.Context, id int, name string) (domain.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m PublisherModel
	err := r.pool.QueryRow(ctx, `
		UPDATE bookshop.publishers SET name = $1 WHERE id = $2
		RETURNING id, name
	`, name, id).Scan(&m.ID, &m.Name)
	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.Publisher{}, fmt.Errorf("publisher with id=%d: %w", id, core_errors.ErrNotFound)
		}
		return domain.Publisher{}, fmt.Errorf("patch publisher: %w", err)
	}

	return m.ToDomain(), nil
}

func (r *BooksRepository) DeletePublisher(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tag, err := r.pool.Exec(ctx, `DELETE FROM bookshop.publishers WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete publisher: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("publisher with id=%d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
