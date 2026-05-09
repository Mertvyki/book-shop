package auth_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *RefreshTokenRepository) CreateToken(
	ctx context.Context,
	token domain.RefreshToken,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(ctx, `INSERT INTO bookshop.refresh_tokens (user_id, token_hash, expires_at, created_at)
	VALUES ($1, $2, $3, $4);`, token.UserID, token.TokenHash, token.ExpiresAt, token.CreatedAt)

	if err != nil {
		return fmt.Errorf("insert refresh token: %w", err)
	}

	return nil
}
