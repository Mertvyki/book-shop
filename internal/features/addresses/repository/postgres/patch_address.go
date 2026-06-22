package addresses_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *AddressesRepository) PatchAddress(ctx context.Context, address domain.Address) (domain.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE bookshop.addresses
	SET
		street_address = $1,
		city = $2,
		postal_code = $3,
		country = $4,
		is_default = $5,
		version = version + 1
	WHERE id = $6 AND user_id = $7 AND version = $8
	RETURNING id, version, user_id, street_address, city, postal_code, country, is_default, created_at
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		address.StreetAddress,
		address.City,
		address.PostalCode,
		address.Country,
		address.IsDefault,
		address.ID,
		address.UserID,
		address.Version,
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
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Address{}, fmt.Errorf("patch address with id=%d: %w", address.ID, core_errors.ErrNotFound)
		}

		return domain.Address{}, fmt.Errorf("patch address: %w", err)
	}

	return addressModel.ToDomain(), nil
}
