package orders_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (r *OrdersRepository) UpdateOrderStatus(ctx context.Context, orderID, version int, status string) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tag, err := r.pool.Exec(ctx,
		`UPDATE bookshop.orders SET status = $1, version = version + 1, updated_at = NOW() WHERE id = $2 AND version = $3`,
		status, orderID, version,
	)
	if err != nil {
		return fmt.Errorf("update order status: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("order with id=%d version=%d: %w", orderID, version, core_errors.ErrNotFound)
	}

	return nil
}
