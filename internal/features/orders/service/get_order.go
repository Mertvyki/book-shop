package orders_service

import (
	"context"
	"fmt"

	orders_postgres_repository "github.com/Mertvyki/book-shop/internal/features/orders/repository/postgres"
)

func (s *OrderService) GetOrder(ctx context.Context, userID, orderID int) (orders_postgres_repository.OrderWithItems, error) {
	result, err := s.orderRepository.GetOrder(ctx, orderID, userID)
	if err != nil {
		return orders_postgres_repository.OrderWithItems{}, fmt.Errorf("get order from repository: %w", err)
	}

	return result, nil
}
