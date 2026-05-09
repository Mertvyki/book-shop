package auth_postgres_repository

import (
	"context"
	"fmt"
)

func (r *RefreshTokenRepository) DeleteToken(
	ctx context.Context,
	id int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(ctx, `DELETE FROM bookshop.refresh_tokens WHERE user_id=$1`, id)
	if err != nil {
		return fmt.Errorf("delete refresh tokens for user %d: %w", id, err)
	}

	return nil
}
