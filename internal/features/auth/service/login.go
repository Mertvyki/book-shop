package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *AuthService) Login(ctx context.Context, email string, password string) (string, string, domain.User, error) {
	user, err := s.userRepository.ByEmail(ctx, email)
	if err != nil {
		return "", "", domain.User{}, fmt.Errorf("get user: %w", err)
	}
	if user == nil {
		return "", "", domain.User{}, fmt.Errorf("invalid email: %w", core_errors.ErrNotFound)
	}

	if !s.hasher.Compare(user.PasswordHash, password) {
		return "", "", domain.User{}, fmt.Errorf("invalid password: %w", core_errors.ErrInvalidArgument)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", "", domain.User{}, fmt.Errorf("generate access token: %w", err)
	}

	refreshToken, err := s.refreshGen.Generate()
	if err != nil {
		return "", "", domain.User{}, fmt.Errorf("generate refresh token: %w", err)
	}

	hash, err := s.refreshGen.Hash(refreshToken)
	if err != nil {
		return "", "", domain.User{}, fmt.Errorf("hash refresh token: %w", err)
	}

	if err := s.authRepository.DeleteToken(ctx, user.ID); err != nil {
		return "", "", domain.User{}, fmt.Errorf("delete old refresh tokens: %w", err)
	}

	refreshDomain := domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: hash,
		ExpiresAt: time.Now().Add(s.refreshTTL),
		CreatedAt: time.Now(),
	}
	if err := s.authRepository.CreateToken(ctx, refreshDomain); err != nil {
		return "", "", domain.User{}, fmt.Errorf("store refresh token: %w", err)
	}

	return accessToken, refreshToken, *user, nil
}
