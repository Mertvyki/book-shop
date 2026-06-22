package addresses_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type AddressModel struct {
	ID            int
	Version       int
	UserID        int
	StreetAddress string
	City          string
	PostalCode    string
	Country       string
	IsDefault     bool
	CreatedAt     time.Time
}

func (m AddressModel) ToDomain() domain.Address {
	return domain.Address{
		ID:            m.ID,
		Version:       m.Version,
		UserID:        m.UserID,
		StreetAddress: m.StreetAddress,
		City:          m.City,
		PostalCode:    m.PostalCode,
		Country:       m.Country,
		IsDefault:     m.IsDefault,
		CreatedAt:     m.CreatedAt,
	}
}
