package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *UserRepository) ByEmail(ctx context.Context, email string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, version, email, password_hash, full_name, phone_number, role, created_at, updated_at
	FROM bookshop.users
	WHERE email=$1;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		email,
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
			return nil, nil
		}

		return &domain.User{}, fmt.Errorf("scan error: %w", err)
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

	return &userDomain, nil
}
