package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) PatchCategory(ctx context.Context, id int, name *string, slug *string, description *string) (domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	sets := make([]string, 0)
	args := []any{}
	argIdx := 1

	if name != nil {
		sets = append(sets, fmt.Sprintf("name = $%d", argIdx))
		args = append(args, *name)
		argIdx++
	}
	if slug != nil {
		sets = append(sets, fmt.Sprintf("slug = $%d", argIdx))
		args = append(args, *slug)
		argIdx++
	}
	if description != nil {
		sets = append(sets, fmt.Sprintf("description = $%d", argIdx))
		args = append(args, *description)
		argIdx++
	}

	if len(sets) == 0 {
		return r.GetCategory(ctx, id)
	}

	args = append(args, id)
	query := fmt.Sprintf(
		`UPDATE bookshop.categories SET %s WHERE id = $%d RETURNING id, name, slug, description`,
		strings.Join(sets, ", "),
		argIdx,
	)

	var m CategoryModel
	err := r.pool.QueryRow(ctx, query, args...).Scan(
		&m.ID, &m.Name, &m.Slug, &m.Description,
	)
	if err != nil {
		return domain.Category{}, fmt.Errorf("patch category: %w", err)
	}

	return m.ToDomain(), nil
}
