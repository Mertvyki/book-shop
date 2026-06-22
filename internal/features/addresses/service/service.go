package addresses_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type AddressesService struct {
	addressesRepository AddressesRepository
}

type AddressesRepository interface {
	CreateAddress(ctx context.Context, address domain.Address) (domain.Address, error)
	ResetDefaultAddresses(ctx context.Context, userID int) error
	GetAddresses(ctx context.Context, userID int) ([]domain.Address, error)
	GetAddress(ctx context.Context, userID int, addrID int) (domain.Address, error)
	PatchAddress(ctx context.Context, address domain.Address) (domain.Address, error)
	DeleteAddress(ctx context.Context, userID, addrID int) error
}

type CreateAddressPayload struct {
	StreetAddress string
	City          string
	PostalCode    string
	Country       string
	IsDefault     bool
}

type PatchAddressPayload struct {
	StreetAddress *string
	City          *string
	PostalCode    *string
	Country       *string
	IsDefault     *bool
}

func NewAddressesService(
	addressesRepository AddressesRepository,
) *AddressesService {

	return &AddressesService{
		addressesRepository: addressesRepository,
	}
}
