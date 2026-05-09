package auth_postgres_repository

import core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"

type RefreshTokenRepository struct {
	pool core_postgres_pool.Pool
}

func NewRefreshTokenRepository(pool core_postgres_pool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		pool: pool,
	}
}
