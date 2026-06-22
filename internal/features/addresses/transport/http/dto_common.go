package addresses_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type AddressDTOResponse struct {
	ID            int       `json:"id"`
	Version       int       `json:"version"`
	UserID        int       `json:"user_id"`
	StreetAddress string    `json:"street_address"`
	City          string    `json:"city"`
	PostalCode    string    `json:"postal_code"`
	Country       string    `json:"country"`
	IsDefault     bool      `json:"is_default"`
	CreatedAt     time.Time `json:"created_at"`
}

func addressDTOFromDomain(
	address domain.Address,
) AddressDTOResponse {

	return AddressDTOResponse{
		ID:            address.ID,
		Version:       address.Version,
		UserID:        address.UserID,
		StreetAddress: address.StreetAddress,
		City:          address.City,
		PostalCode:    address.PostalCode,
		Country:       address.Country,
		IsDefault:     address.IsDefault,
		CreatedAt:     address.CreatedAt,
	}
}

func addressesDTOFromDomains(addresses []domain.Address) []AddressDTOResponse {
	addressDTO := make([]AddressDTOResponse, len(addresses))
	for i, address := range addresses {
		addressDTO[i] = addressDTOFromDomain(address)
	}

	return addressDTO
}
