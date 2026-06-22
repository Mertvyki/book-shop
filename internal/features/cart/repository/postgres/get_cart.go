package cart_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *CartRepository) GetCart(ctx context.Context, userID int) ([]domain.CartItem, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT ci.id, ci.version, ci.user_id, ci.book_id, ci.quantity, ci.added_at,
	       b.title, b.price, b.cover_image_key, b.book_type
	FROM bookshop.cart_items ci
	LEFT JOIN bookshop.books b ON b.id = ci.book_id AND b.deleted_at IS NULL
	WHERE ci.user_id = $1
	ORDER BY ci.added_at ASC
	`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query cart items: %w", err)
	}
	defer rows.Close()

	items := make([]domain.CartItem, 0)

	for rows.Next() {
		var model CartItemModel

		err = rows.Scan(
			&model.ID,
			&model.Version,
			&model.UserID,
			&model.BookID,
			&model.Quantity,
			&model.AddedAt,
			&model.Title,
			&model.Price,
			&model.CoverImageKey,
			&model.BookType,
		)
		if err != nil {
			return nil, fmt.Errorf("scan cart item: %w", err)
		}

		items = append(items, model.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate cart items: %w", err)
	}

	return items, nil
}
