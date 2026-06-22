package book_postgres_repository

import (
	"context"
	"fmt"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) CreateAuthor(ctx context.Context, name string, bio *string, birthYear *int) (domain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m AuthorModel
	err := r.pool.QueryRow(ctx, `
		INSERT INTO bookshop.authors (name, bio, birth_year, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, bio, birth_year, created_at
	`, name, bio, birthYear, time.Now().UTC()).Scan(
		&m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt,
	)
	if err != nil {
		return domain.Author{}, fmt.Errorf("create author: %w", err)
	}

	return m.ToDomain(), nil
}
