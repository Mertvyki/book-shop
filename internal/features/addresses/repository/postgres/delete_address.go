package addresses_postgres_repository

import (
	"context"
	"fmt"
)

func (r *AddressesRepository) DeleteAddress(ctx context.Context, userID, addrID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	_, err := r.pool.Exec(ctx, `DELETE FROM bookshop.addresses WHERE id = $1 AND user_id = $2`, addrID, userID)
	if err != nil {
		return fmt.Errorf("delete address: %w", err)
	}

	return nil
}
