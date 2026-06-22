package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UserRepository) CountUsers(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM bookshop.users`).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count users: %w", err)
	}

	return total, nil
}
