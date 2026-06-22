package addresses_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *AddressesService) CreateAddress(ctx context.Context, userID int, request CreateAddressPayload) (domain.Address, error) {
	country := request.Country
	if country == "" {
		country = "Россия"
	}

	address := domain.NewAddressUninitialized(
		userID,
		request.StreetAddress,
		request.City,
		request.PostalCode,
		country,
		request.IsDefault,
	)

	if address.IsDefault {
		if err := s.addressesRepository.ResetDefaultAddresses(ctx, userID); err != nil {
			return domain.Address{}, err
		}
	}

	createdAddress, err := s.addressesRepository.CreateAddress(ctx, address)
	if err != nil {
		return domain.Address{}, err
	}

	return createdAddress, nil
}
