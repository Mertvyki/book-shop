package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *UserRepository) ByID(ctx context.Context, id int) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var userModel UserModel
	err := r.pool.QueryRow(ctx,
		`SELECT id, version, email, password_hash, full_name, phone_number, role, created_at, updated_at
		 FROM bookshop.users WHERE id = $1`, id,
	).Scan(
		&userModel.ID, &userModel.Version, &userModel.Email, &userModel.PasswordHash,
		&userModel.FullName, &userModel.PhoneNumber, &userModel.Role,
		&userModel.CreatedAt, &userModel.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("query user by id: %w", err)
	}
	user := domain.NewUser(
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
	return &user, nil
}
