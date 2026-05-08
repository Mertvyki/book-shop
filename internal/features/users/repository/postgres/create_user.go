package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (r *UserRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO bookshop.users (email, password_hash, full_name, phone_number, role, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, version, email, password_hash, full_name, phone_number, role, created_at, updated_at;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.FullName,
		user.PhoneNumber,
		user.Role,
		user.CreatedAt,
		user.UpdatedAt,
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
