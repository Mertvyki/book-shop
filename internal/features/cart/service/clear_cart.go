package cart_service

import (
	"context"
	"fmt"
)

func (s *CartService) ClearCart(ctx context.Context, userID int) error {
	if err := s.cartRepository.ClearCart(ctx, userID); err != nil {
		return fmt.Errorf("clear cart in repository: %w", err)
	}

	return nil
}
