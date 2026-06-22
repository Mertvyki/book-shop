package cart_postgres_repository

import (
	"context"
	"fmt"
)

func (r *CartRepository) RemoveItem(ctx context.Context, itemID, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(ctx, `DELETE FROM bookshop.cart_items WHERE id = $1 AND user_id = $2`, itemID, userID)
	if err != nil {
		return fmt.Errorf("delete cart item: %w", err)
	}

	return nil
}
