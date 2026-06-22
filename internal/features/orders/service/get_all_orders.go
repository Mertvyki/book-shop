package orders_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type GetAllOrdersResult struct {
	Orders []domain.Order
	Total  int
}

func (s *OrderService) GetAllOrders(ctx context.Context, limit, offset int) (GetAllOrdersResult, error) {
	orders, err := s.orderRepository.GetAllOrders(ctx, limit, offset)
	if err != nil {
		return GetAllOrdersResult{}, fmt.Errorf("get all orders: %w", err)
	}

	total, err := s.orderRepository.CountAllOrders(ctx)
	if err != nil {
		return GetAllOrdersResult{}, fmt.Errorf("count all orders: %w", err)
	}

	return GetAllOrdersResult{
		Orders: orders,
		Total:  total,
	}, nil
}
