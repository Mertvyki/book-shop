package cart_postgres_repository

import core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"

type CartRepository struct {
	pool core_postgres_pool.Pool
}

func NewCartRepository(pool core_postgres_pool.Pool) *CartRepository {
	return &CartRepository{
		pool: pool,
	}
}
