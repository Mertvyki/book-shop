package auth_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *RefreshTokenRepository) GetTokenByHash(
	ctx context.Context,
	hash string,
) (domain.RefreshToken, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m RefreshTokenModel
	err := r.pool.QueryRow(ctx, `SELECT id, user_id, token_hash, expires_at, created_at FROM bookshop.refresh_tokens WHERE token_hash=$1;`, hash).Scan(&m.ID, &m.UserID, &m.TokenHash, &m.ExpiresAt, &m.CreatedAt)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.RefreshToken{}, fmt.Errorf("token not found: %w", core_postgres_pool.ErrNoRows)
		}

		return domain.RefreshToken{}, fmt.Errorf("find token by hash: %w", err)
	}

	return m.ToDomain(), nil
}
