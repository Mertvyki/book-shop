package orders_service

import (
	"context"
	"fmt"
)

func (s *OrderService) HasUserPurchasedBook(ctx context.Context, userID, bookID int) (bool, error) {
	purchased, err := s.orderRepository.HasUserPurchasedBook(ctx, userID, bookID)
	if err != nil {
		return false, fmt.Errorf("has user purchased book: %w", err)
	}

	return purchased, nil
}
