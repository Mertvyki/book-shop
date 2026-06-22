package orders_service

import (
	"context"
	"fmt"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

var validStatuses = map[string]bool{
	"pending":   true,
	"paid":      true,
	"shipped":   true,
	"delivered": true,
	"cancelled": true,
}

func (s *OrderService) PatchOrderStatus(ctx context.Context, orderID, version int, status string) error {
	if !validStatuses[status] {
		return fmt.Errorf("invalid status %q: %w", status, core_errors.ErrInvalidArgument)
	}

	if err := s.orderRepository.UpdateOrderStatus(ctx, orderID, version, status); err != nil {
		return fmt.Errorf("update order status: %w", err)
	}

	return nil
}
