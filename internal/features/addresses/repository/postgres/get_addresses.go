package addresses_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *AddressesRepository) GetAddresses(ctx context.Context, userID int) ([]domain.Address, error) {
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
	WHERE user_id = $1
	ORDER BY
		is_default DESC,
		id ASC
	`

	rows, err := r.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query addresses: %w", err)
	}
	defer rows.Close()

	addresses := make([]domain.Address, 0)

	for rows.Next() {

		var addressModel AddressModel

		err = rows.Scan(
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
			return nil, fmt.Errorf(
				"scan address: %w",
				err,
			)
		}

		addresses = append(
			addresses,
			addressModel.ToDomain(),
		)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf(
			"iterate addresses: %w",
			err,
		)
	}

	return addresses, nil
}
