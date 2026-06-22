package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) CreateCategory(ctx context.Context, name, slug string, description *string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(name, " ", "-"))
	}

	var m CategoryModel
	err := r.pool.QueryRow(ctx, `
		INSERT INTO bookshop.categories (name, slug, description)
		VALUES ($1, $2, $3)
		RETURNING id, name, slug, description
	`, name, slug, description).Scan(
		&m.ID, &m.Name, &m.Slug, &m.Description,
	)
	if err != nil {
		return domain.Category{}, fmt.Errorf("create category: %w", err)
	}

	return m.ToDomain(), nil
}
