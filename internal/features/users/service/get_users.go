package user_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

type GetUsersResult struct {
	Users []domain.User
	Total int
}

func (s *UserService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) (GetUsersResult, error) {
	if limit != nil && *limit < 0 {
		return GetUsersResult{}, fmt.Errorf("limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return GetUsersResult{}, fmt.Errorf("offset must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	users, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return GetUsersResult{}, fmt.Errorf("get users from repository: %w", err)
	}

	total, err := s.usersRepository.CountUsers(ctx)
	if err != nil {
		return GetUsersResult{}, fmt.Errorf("count users: %w", err)
	}

	return GetUsersResult{
		Users: users,
		Total: total,
	}, nil
}
