package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *UserRepository) PatchUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE bookshop.users
	SET
		email = $1,
        password_hash = $2,
        full_name = $3,
        phone_number = $4,
        role = $5,
        updated_at = NOW(),
        version = version + 1
	WHERE id=$6 AND version=$7
	RETURNING
		id,
        version,
        email,
        password_hash,
        full_name,
        phone_number,
        role,
        created_at,
        updated_at;
	`
	row := r.pool.QueryRow(ctx, query,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.PhoneNumber,
		user.Role,
		user.ID,
		user.Version,
	)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Email,
		&userModel.PasswordHash,
		&userModel.FullName,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.CreatedAt,
		&userModel.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id=%d concurrently accessed: %w", user.ID, core_errors.ErrConflict)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Email,
		userModel.PasswordHash,
		userModel.FullName,
		userModel.PhoneNumber,
		userModel.Role,
		userModel.CreatedAt,
		userModel.UpdatedAt,
	)

	return userDomain, nil
}
