package addresses_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *AddressesService) GetAddress(
	ctx context.Context,
	userID int,
	addrID int,
) (domain.Address, error) {
	address, err := s.addressesRepository.GetAddress(ctx, userID, addrID)
	if err != nil {
		return domain.Address{}, fmt.Errorf("get address from repository: %w", err)
	}

	return address, nil
}
