package orders_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type CreateOrderResult struct {
	Order        domain.Order
	Items        []domain.OrderItem
	DeliveryCost float64
}

func (r *OrdersRepository) CreateOrder(ctx context.Context, userID int, shippingAddressID *int, items []domain.CartItem, deliveryCost float64) (CreateOrderResult, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return CreateOrderResult{}, fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Batch fetch book prices
	type bookPriceInfo struct {
		Price    float64
		BookType string
	}

	bookPrices := make(map[int]bookPriceInfo)
	for _, item := range items {
		var info bookPriceInfo
		err := tx.QueryRow(ctx, `SELECT price, book_type FROM bookshop.books WHERE id = $1 AND deleted_at IS NULL`, item.BookID).Scan(&info.Price, &info.BookType)
		if err != nil {
			return CreateOrderResult{}, fmt.Errorf("get book price for id %d: %w", item.BookID, err)
		}
		bookPrices[item.BookID] = info
	}

	var totalAmount float64
	for _, item := range items {
		info := bookPrices[item.BookID]
		totalAmount += info.Price * float64(item.Quantity)
	}

	totalAmount += deliveryCost

	var order domain.Order
	err = tx.QueryRow(ctx, `
		INSERT INTO bookshop.orders (user_id, status, total_amount, shipping_address_id, created_at, updated_at)
		VALUES ($1, 'paid', $2, $3, NOW(), NOW())
		RETURNING id, version, user_id, status, total_amount, shipping_address_id, payment_method, created_at, updated_at
	`, userID, totalAmount, shippingAddressID).Scan(
		&order.ID, &order.Version, &order.UserID, &order.Status, &order.TotalAmount,
		&order.ShippingAddressID, &order.PaymentMethod, &order.CreatedAt, &order.UpdatedAt,
	)
	if err != nil {
		return CreateOrderResult{}, fmt.Errorf("insert order: %w", err)
	}

	orderItems := make([]domain.OrderItem, 0, len(items))

	for _, item := range items {
		info := bookPrices[item.BookID]

		var orderItem domain.OrderItem
		err = tx.QueryRow(ctx, `
			INSERT INTO bookshop.order_items (order_id, book_id, quantity, unit_price, item_type)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id, version, order_id, book_id, quantity, unit_price, item_type
		`, order.ID, item.BookID, item.Quantity, info.Price, info.BookType).Scan(
			&orderItem.ID, &orderItem.Version, &orderItem.OrderID,
			&orderItem.BookID, &orderItem.Quantity, &orderItem.UnitPrice, &orderItem.ItemType,
		)
		if err != nil {
			return CreateOrderResult{}, fmt.Errorf("insert order item: %w", err)
		}

		orderItems = append(orderItems, orderItem)
	}

	_, err = tx.Exec(ctx, `DELETE FROM bookshop.cart_items WHERE user_id = $1`, userID)
	if err != nil {
		return CreateOrderResult{}, fmt.Errorf("clear cart: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return CreateOrderResult{}, fmt.Errorf("commit transaction: %w", err)
	}

	return CreateOrderResult{Order: order, Items: orderItems, DeliveryCost: deliveryCost}, nil
}
