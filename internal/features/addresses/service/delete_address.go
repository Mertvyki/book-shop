package addresses_service

import (
	"context"
	"fmt"
)

func (s *AddressesService) DeleteAddress(ctx context.Context, userID, addrID int) error {
	if err := s.addressesRepository.DeleteAddress(ctx, userID, addrID); err != nil {
		return fmt.Errorf("delete address from repository: %w", err)
	}

	return nil
}
