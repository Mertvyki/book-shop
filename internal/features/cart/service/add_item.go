package cart_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *CartService) AddItem(ctx context.Context, userID, bookID, quantity int) (domain.CartItem, error) {
	_, err := s.bookRepository.GetBook(ctx, bookID)
	if err != nil {
		return domain.CartItem{}, fmt.Errorf("validate book: %w", core_errors.ErrInvalidArgument)
	}

	item := domain.NewCartItemUninitialized(userID, bookID, quantity)

	createdItem, err := s.cartRepository.AddItem(ctx, item)
	if err != nil {
		return domain.CartItem{}, fmt.Errorf("add item to repository: %w", err)
	}

	return createdItem, nil
}
