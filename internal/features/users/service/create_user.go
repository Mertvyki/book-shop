package user_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	email string,
	password string,
	fullName string,
	phoneNumber *string,
) (domain.User, error) {
	existing, err := s.usersRepository.ByEmail(ctx, email)
	if err != nil {
		return domain.User{}, fmt.Errorf("validate email: %w", err)
	}
	if existing != nil {
		return domain.User{}, fmt.Errorf("email already exists: %w", core_errors.ErrConflict)
	}

	hashedPassword, err := s.hasher.Hash(password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	userDomain := domain.NewUserUninitialized(
		email,
		hashedPassword,
		fullName,
		phoneNumber,
	)

	user, err := s.usersRepository.CreateUser(ctx, userDomain)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}
