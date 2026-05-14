package book_postgres_repository

import core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"

type BooksRepository struct {
	pool core_postgres_pool.Pool
}

func NewBooksRepository(pool core_postgres_pool.Pool) *BooksRepository {
	return &BooksRepository{
		pool: pool,
	}
}
