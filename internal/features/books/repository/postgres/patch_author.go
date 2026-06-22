package book_postgres_repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *BooksRepository) PatchAuthor(ctx context.Context, id int, name *string, bio *string, birthYear *int) (domain.Author, error) {
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
	if bio != nil {
		sets = append(sets, fmt.Sprintf("bio = $%d", argIdx))
		args = append(args, *bio)
		argIdx++
	}
	if birthYear != nil {
		sets = append(sets, fmt.Sprintf("birth_year = $%d", argIdx))
		args = append(args, *birthYear)
		argIdx++
	}

	if len(sets) == 0 {
		return r.GetAuthor(ctx, id)
	}

	args = append(args, id)
	query := fmt.Sprintf(
		`UPDATE bookshop.authors SET %s WHERE id = $%d RETURNING id, name, bio, birth_year, created_at`,
		strings.Join(sets, ", "),
		argIdx,
	)

	var m AuthorModel
	err := r.pool.QueryRow(ctx, query, args...).Scan(
		&m.ID, &m.Name, &m.Bio, &m.BirthYear, &m.CreatedAt,
	)
	if err != nil {
		return domain.Author{}, fmt.Errorf("patch author: %w", err)
	}

	return m.ToDomain(), nil
}
