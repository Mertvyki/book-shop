package cart_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *CartRepository) UpdateItem(ctx context.Context, itemID, userID, quantity int) (domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE bookshop.cart_items
	SET quantity = $1, version = version + 1
	WHERE id = $2 AND user_id = $3
	RETURNING id, version, user_id, book_id, quantity, added_at
	`

	row := r.pool.QueryRow(ctx, query, quantity, itemID, userID)

	var model CartItemModel

	err := row.Scan(
		&model.ID,
		&model.Version,
		&model.UserID,
		&model.BookID,
		&model.Quantity,
		&model.AddedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.CartItem{}, fmt.Errorf("cart item with id=%d: %w", itemID, core_errors.ErrNotFound)
		}

		return domain.CartItem{}, fmt.Errorf("update cart item: %w", err)
	}

	return model.ToDomain(), nil
}
