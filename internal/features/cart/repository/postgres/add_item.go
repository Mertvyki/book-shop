package cart_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *CartRepository) AddItem(ctx context.Context, item domain.CartItem) (domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO bookshop.cart_items (user_id, book_id, quantity, added_at)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (user_id, book_id)
	DO UPDATE SET quantity = bookshop.cart_items.quantity + EXCLUDED.quantity
	RETURNING id, version, user_id, book_id, quantity, added_at
	`

	row := r.pool.QueryRow(ctx, query, item.UserID, item.BookID, item.Quantity, item.AddedAt)

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
		return domain.CartItem{}, fmt.Errorf("add cart item: %w", err)
	}

	return model.ToDomain(), nil
}
