package addresses_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *AddressesService) GetAddresses(ctx context.Context, userID int) ([]domain.Address, error) {
	addresses, err := s.addressesRepository.GetAddresses(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get addresses from repository: %w", err)
	}

	return addresses, nil
}
