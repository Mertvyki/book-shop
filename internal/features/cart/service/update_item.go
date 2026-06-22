package cart_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *CartService) UpdateItem(ctx context.Context, userID, itemID, quantity int) (domain.CartItem, error) {
	if quantity < 1 {
		return domain.CartItem{}, fmt.Errorf("quantity must be positive: %w", core_errors.ErrInvalidArgument)
	}

	item, err := s.cartRepository.UpdateItem(ctx, itemID, userID, quantity)
	if err != nil {
		return domain.CartItem{}, fmt.Errorf("update item in repository: %w", err)
	}

	return item, nil
}
