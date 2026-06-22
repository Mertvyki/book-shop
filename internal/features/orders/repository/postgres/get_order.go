package orders_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

type OrderWithItems struct {
	Order domain.Order
	Items []domain.OrderItem
}

func (r *OrdersRepository) GetOrder(ctx context.Context, orderID, userID int) (OrderWithItems, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var orderModel OrderModel
	err := r.pool.QueryRow(ctx, `
		SELECT id, version, user_id, status, total_amount, shipping_address_id, payment_method, created_at, updated_at
		FROM bookshop.orders
		WHERE id = $1 AND user_id = $2
	`, orderID, userID).Scan(
		&orderModel.ID, &orderModel.Version, &orderModel.UserID, &orderModel.Status,
		&orderModel.TotalAmount, &orderModel.ShippingAddressID, &orderModel.PaymentMethod,
		&orderModel.CreatedAt, &orderModel.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return OrderWithItems{}, fmt.Errorf("order with id=%d: %w", orderID, core_errors.ErrNotFound)
		}

		return OrderWithItems{}, fmt.Errorf("get order: %w", err)
	}

	rows, err := r.pool.Query(ctx, `
		SELECT oi.id, oi.version, oi.order_id, oi.book_id, oi.quantity, oi.unit_price, oi.item_type,
			COALESCE(b.title, ''), b.file_key
		FROM bookshop.order_items oi
		JOIN bookshop.books b ON b.id = oi.book_id
		WHERE oi.order_id = $1
	`, orderID)
	if err != nil {
		return OrderWithItems{}, fmt.Errorf("query order items: %w", err)
	}
	defer rows.Close()

	items := make([]domain.OrderItem, 0)

	for rows.Next() {
		var itemModel OrderItemModel
		err = rows.Scan(
			&itemModel.ID, &itemModel.Version, &itemModel.OrderID,
			&itemModel.BookID, &itemModel.Quantity, &itemModel.UnitPrice, &itemModel.ItemType,
			&itemModel.Title, &itemModel.FileKey,
		)
		if err != nil {
			return OrderWithItems{}, fmt.Errorf("scan order item: %w", err)
		}

		items = append(items, itemModel.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return OrderWithItems{}, fmt.Errorf("iterate order items: %w", err)
	}

	return OrderWithItems{
		Order: orderModel.ToDomain(),
		Items: items,
	}, nil
}
