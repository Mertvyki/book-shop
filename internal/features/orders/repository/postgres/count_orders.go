package orders_postgres_repository

import (
	"context"
	"fmt"
)

func (r *OrdersRepository) CountOrders(ctx context.Context, userID int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var total int
	err := r.pool.QueryRow(ctx, `SELECT COUNT(*) FROM bookshop.orders WHERE user_id = $1`, userID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("count orders: %w", err)
	}

	return total, nil
}
