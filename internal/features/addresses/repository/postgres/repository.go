package addresses_postgres_repository

import core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"

type AddressesRepository struct {
	pool core_postgres_pool.Pool
}

func NewAddressesRepository(pool core_postgres_pool.Pool) *AddressesRepository {
	return &AddressesRepository{
		pool: pool,
	}
}
