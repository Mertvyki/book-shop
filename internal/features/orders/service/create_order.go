package orders_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *OrderService) CreateOrder(ctx context.Context, userID int, shippingAddressID *int) (domain.Order, []domain.OrderItem, float64, error) {
	cartItems, err := s.cartRepository.GetCart(ctx, userID)
	if err != nil {
		return domain.Order{}, nil, 0, fmt.Errorf("get cart: %w", err)
	}

	if len(cartItems) == 0 {
		return domain.Order{}, nil, 0, fmt.Errorf("cart is empty: %w", core_errors.ErrInvalidArgument)
	}

	hasPhysical := false
	for _, item := range cartItems {
		if item.BookType == "physical" {
			hasPhysical = true
			break
		}
	}

	if hasPhysical && shippingAddressID == nil {
		return domain.Order{}, nil, 0, fmt.Errorf("shipping address is required for physical books: %w", core_errors.ErrInvalidArgument)
	}

	if shippingAddressID != nil {
		_, err := s.addressRepository.GetAddress(ctx, userID, *shippingAddressID)
		if err != nil {
			return domain.Order{}, nil, 0, fmt.Errorf("validate shipping address: %w", err)
		}
	}

	deliveryCost := 0.0
	if hasPhysical {
		deliveryCost = 250.0
	}

	result, err := s.orderRepository.CreateOrder(ctx, userID, shippingAddressID, cartItems, deliveryCost)
	if err != nil {
		return domain.Order{}, nil, 0, fmt.Errorf("create order: %w", err)
	}

	return result.Order, result.Items, deliveryCost, nil
}
