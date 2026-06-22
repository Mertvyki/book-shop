package addresses_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *AddressesRepository) CreateAddress(ctx context.Context, address domain.Address) (domain.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO bookshop.addresses (user_id, street_address, city, postal_code, country, is_default)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING
		id,
		version,
		user_id,
		street_address,
		city,
		postal_code,
		country,
		is_default,
		created_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		address.UserID,
		address.StreetAddress,
		address.City,
		address.PostalCode,
		address.Country,
		address.IsDefault,
	)

	var addressModel AddressModel

	err := row.Scan(
		&addressModel.ID,
		&addressModel.Version,
		&addressModel.UserID,
		&addressModel.StreetAddress,
		&addressModel.City,
		&addressModel.PostalCode,
		&addressModel.Country,
		&addressModel.IsDefault,
		&addressModel.CreatedAt,
	)
	if err != nil {
		return domain.Address{}, fmt.Errorf(
			"create address: %w",
			err,
		)
	}

	return addressModel.ToDomain(), nil
}
