package cart_service

import (
	"context"
	"fmt"
)

func (s *CartService) RemoveItem(ctx context.Context, userID, itemID int) error {
	if err := s.cartRepository.RemoveItem(ctx, itemID, userID); err != nil {
		return fmt.Errorf("remove item from repository: %w", err)
	}

	return nil
}
