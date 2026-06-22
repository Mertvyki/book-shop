package book_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) ListAuthors(ctx context.Context) ([]domain.Author, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	rows, err := r.pool.Query(ctx, `SELECT id, name, bio, birth_year, created_at FROM bookshop.authors ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("query authors: %w", err)
	}
	defer rows.Close()

	authors := make([]domain.Author, 0)
	for rows.Next() {
		var m AuthorModel
		if err := rows.Scan(&m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("scan author: %w", err)
		}
		authors = append(authors, m.ToDomain())
	}

	return authors, rows.Err()
}

func (r *BooksRepository) ListCategories(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	rows, err := r.pool.Query(ctx, `SELECT id, name, slug, description FROM bookshop.categories ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("query categories: %w", err)
	}
	defer rows.Close()

	categories := make([]domain.Category, 0)
	for rows.Next() {
		var m CategoryModel
		if err := rows.Scan(&m.ID, &m.Name, &m.Slug, &m.Description); err != nil {
			return nil, fmt.Errorf("scan category: %w", err)
		}
		categories = append(categories, m.ToDomain())
	}

	return categories, rows.Err()
}

func (r *BooksRepository) ListPublishers(ctx context.Context) ([]domain.Publisher, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	rows, err := r.pool.Query(ctx, `SELECT id, name FROM bookshop.publishers ORDER BY name`)
	if err != nil {
		return nil, fmt.Errorf("query publishers: %w", err)
	}
	defer rows.Close()

	publishers := make([]domain.Publisher, 0)
	for rows.Next() {
		var m PublisherModel
		if err := rows.Scan(&m.ID, &m.Name); err != nil {
			return nil, fmt.Errorf("scan publisher: %w", err)
		}
		publishers = append(publishers, m.ToDomain())
	}

	return publishers, rows.Err()
}
