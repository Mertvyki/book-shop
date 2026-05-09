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

	var model UserModel
	err := r.pool.QueryRow(ctx,
		`SELECT id, version, email, password_hash, full_name, phone_number, role, created_at, updated_at
		 FROM bookshop.users WHERE id = $1`, id,
	).Scan(
		&model.ID, &model.Version, &model.Email, &model.PasswordHash,
		&model.FullName, &model.PhoneNumber, &model.Role,
		&model.CreatedAt, &model.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("query user by id: %w", err)
	}
	user := domain.NewUser(
		model.ID,
		model.Version,
		model.Email,
		model.PasswordHash,
		model.FullName,
		model.PhoneNumber,
		model.Role,
		model.CreatedAt,
		model.UpdatedAt,
	)
	return &user, nil
}
