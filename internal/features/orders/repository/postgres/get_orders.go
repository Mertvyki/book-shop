package orders_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *OrdersRepository) GetOrders(ctx context.Context, userID int, limit, offset int) ([]domain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, user_id, status, total_amount, shipping_address_id, payment_method, created_at, updated_at
	FROM bookshop.orders
	WHERE user_id = $1
	ORDER BY created_at DESC
	LIMIT $2 OFFSET $3
	`

	rows, err := r.pool.Query(ctx, query, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("query orders: %w", err)
	}
	defer rows.Close()

	orders := make([]domain.Order, 0)

	for rows.Next() {
		var model OrderModel
		err = rows.Scan(
			&model.ID, &model.Version, &model.UserID, &model.Status,
			&model.TotalAmount, &model.ShippingAddressID, &model.PaymentMethod,
			&model.CreatedAt, &model.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan order: %w", err)
		}

		orders = append(orders, model.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate orders: %w", err)
	}

	return orders, nil
}
