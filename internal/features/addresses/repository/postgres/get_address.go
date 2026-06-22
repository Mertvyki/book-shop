package addresses_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *AddressesRepository) GetAddress(
	ctx context.Context,
	userID int,
	addrID int,
) (domain.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT
		id,
		version,
		user_id,
		street_address,
		city,
		postal_code,
		country,
		is_default,
		created_at
	FROM bookshop.addresses
	WHERE user_id = $1 AND id = $2
	`

	row := r.pool.QueryRow(ctx, query, userID, addrID)

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
			return domain.Address{}, fmt.Errorf("address with id=%d: user with id=%d: %w", addrID, userID, core_errors.ErrNotFound)
		}

		return domain.Address{}, fmt.Errorf("scan error: %w", err)
	}

	addressDomain := domain.NewAddress(
		addressModel.ID,
		addressModel.Version,
		addressModel.UserID,
		addressModel.StreetAddress,
		addressModel.City,
		addressModel.PostalCode,
		addressModel.Country,
		addressModel.IsDefault,
		addressModel.CreatedAt,
	)

	return addressDomain, nil
}
