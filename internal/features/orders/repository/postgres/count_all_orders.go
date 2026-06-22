package orders_postgres_repository

import (
	"context"
	"fmt"
)

func (r *OrdersRepository) CountAllOrders(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM bookshop.orders`).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count all orders: %w", err)
	}

	return total, nil
}
