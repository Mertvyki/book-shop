package orders_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type GetOrdersResult struct {
	Orders []domain.Order
	Total  int
}

func (s *OrderService) GetOrders(ctx context.Context, userID int, limit, offset int) (GetOrdersResult, error) {
	orders, err := s.orderRepository.GetOrders(ctx, userID, limit, offset)
	if err != nil {
		return GetOrdersResult{}, fmt.Errorf("get orders from repository: %w", err)
	}

	total, err := s.orderRepository.CountOrders(ctx, userID)
	if err != nil {
		return GetOrdersResult{}, fmt.Errorf("count orders: %w", err)
	}

	return GetOrdersResult{
		Orders: orders,
		Total:  total,
	}, nil
}
