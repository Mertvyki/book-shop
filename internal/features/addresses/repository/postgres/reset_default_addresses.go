package addresses_postgres_repository

import "context"

func (r *AddressesRepository) ResetDefaultAddresses(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE bookshop.addresses
	SET is_default = false
	WHERE user_id = $1
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		userID,
	)

	return err
}
