package cart_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *CartService) GetCart(ctx context.Context, userID int) ([]domain.CartItem, error) {
	items, err := s.cartRepository.GetCart(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get cart from repository: %w", err)
	}

	return items, nil
}
