package orders_postgres_repository

import (
	"context"
	"fmt"
)

func (r *OrdersRepository) HasUserPurchasedBook(ctx context.Context, userID, bookID int) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var count int
	err := r.pool.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM bookshop.orders o
		JOIN bookshop.order_items oi ON oi.order_id = o.id
		WHERE o.user_id = $1 AND oi.book_id = $2 AND o.status = 'paid'
	`, userID, bookID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("check purchase: %w", err)
	}

	return count > 0, nil
}
