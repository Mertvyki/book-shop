package cart_postgres_repository

import (
	"context"
	"fmt"
)

func (r *CartRepository) ClearCart(ctx context.Context, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(ctx, `DELETE FROM bookshop.cart_items WHERE user_id = $1`, userID)
	if err != nil {
		return fmt.Errorf("clear cart: %w", err)
	}

	return nil
}
