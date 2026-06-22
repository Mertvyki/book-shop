package user_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *UserService) PatchUser(
	ctx context.Context,
	userID int,
	patch PatchUserPayload,
) (domain.User, error) {
	user, err := s.GetUser(ctx, userID)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user: %w", err)
	}

	if patch.Password.Set {
		if !patch.OldPassword.Set || patch.OldPassword.Value == nil {
			return domain.User{}, fmt.Errorf("old password is required to change password: %w", core_errors.ErrInvalidArgument)
		}
		if !s.hasher.Compare(user.PasswordHash, *patch.OldPassword.Value) {
			return domain.User{}, fmt.Errorf("old password is incorrect: %w", core_errors.ErrInvalidArgument)
		}

		if patch.Password.Value == nil || len(*patch.Password.Value) < 6 {
			return domain.User{}, fmt.Errorf("new password must be at least 6 characters: %w", core_errors.ErrInvalidArgument)
		}
		hashed, err := s.hasher.Hash(*patch.Password.Value)
		if err != nil {
			return domain.User{}, fmt.Errorf("hash password: %w", err)
		}
		user.PasswordHash = hashed
	}
	if patch.Email.Set {
		if patch.Email.Value == nil || *patch.Email.Value == "" {
			return domain.User{}, fmt.Errorf("email cannot be empty: %w", core_errors.ErrInvalidArgument)
		}
		if *patch.Email.Value != user.Email {
			existing, err := s.usersRepository.ByEmail(ctx, *patch.Email.Value)
			if err != nil && !errors.Is(err, core_errors.ErrNotFound) {
				return domain.User{}, err
			}
			if existing != nil {
				return domain.User{}, fmt.Errorf("email already taken: %w", core_errors.ErrConflict)
			}
			user.Email = *patch.Email.Value
		}
	}
	if patch.FullName.Set {
		if patch.FullName.Value == nil || len(*patch.FullName.Value) < 15 {
			return domain.User{}, fmt.Errorf("full name to short: %w", core_errors.ErrInvalidArgument)
		}
		user.FullName = *patch.FullName.Value
	}
	if patch.PhoneNumber.Set {
		user.PhoneNumber = patch.PhoneNumber.Value
	}
	if patch.Role.Set {
		if patch.Role.Value == nil || (*patch.Role.Value != "customer" && *patch.Role.Value != "employee") {
			return domain.User{}, fmt.Errorf("invalid role: must be 'customer' or 'employee': %w", core_errors.ErrInvalidArgument)
		}
		user.Role = *patch.Role.Value
	}
	//user.UpdatedAt = time.Now().UTC()
	patchedUser, err := s.usersRepository.PatchUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("patched user: %w", err)
	}
	return patchedUser, nil
}
