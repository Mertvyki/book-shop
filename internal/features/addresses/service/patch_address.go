package addresses_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *AddressesService) PatchAddress(ctx context.Context, userID, addrID int, patch PatchAddressPayload) (domain.Address, error) {
	existing, err := s.addressesRepository.GetAddress(ctx, userID, addrID)
	if err != nil {
		return domain.Address{}, fmt.Errorf("get address for patch: %w", err)
	}

	if patch.StreetAddress != nil {
		existing.StreetAddress = *patch.StreetAddress
	}

	if patch.City != nil {
		existing.City = *patch.City
	}

	if patch.PostalCode != nil {
		existing.PostalCode = *patch.PostalCode
	}

	if patch.Country != nil {
		existing.Country = *patch.Country
	}

	if patch.IsDefault != nil {
		existing.IsDefault = *patch.IsDefault
	}

	if existing.IsDefault {
		if err := s.addressesRepository.ResetDefaultAddresses(ctx, userID); err != nil {
			return domain.Address{}, fmt.Errorf("reset default addresses: %w", err)
		}
	}

	updatedAddress, err := s.addressesRepository.PatchAddress(ctx, existing)
	if err != nil {
		return domain.Address{}, fmt.Errorf("patch address in repository: %w", err)
	}

	return updatedAddress, nil
}
