package auth_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

func (s *AuthService) Refresh(
	ctx context.Context,
	oldRefreshToken string,
) (string, string, error) {
	hash, err := s.refreshGen.Hash(oldRefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("hash token: %w", err)
	}

	refreshToken, err := s.authRepository.GetTokenByHash(ctx, hash)
	if err != nil {
		return "", "", fmt.Errorf("get refresh token: %w", core_errors.ErrNotFound)
	}

	if time.Now().After(refreshToken.ExpiresAt) {
		return "", "", fmt.Errorf("refresh token expired: %w", core_errors.ErrInvalidArgument)
	}

	user, err := s.userRepository.ByID(ctx, refreshToken.UserID)
	if err != nil {
		return "", "", fmt.Errorf("user not found: %w", err)
	}

	if err := s.authRepository.DeleteToken(ctx, user.ID); err != nil {
		return "", "", fmt.Errorf("delete old token: %w", err)
	}

	accessToken, err := s.jwtManager.GenerateAccessToken(user.ID, user.Role)
	if err != nil {
		return "", "", fmt.Errorf("generate access: %w", err)
	}

	newRefreshToken, err := s.refreshGen.Generate()
	if err != nil {
		return "", "", fmt.Errorf("generate refresh: %w", err)
	}

	newHash, err := s.refreshGen.Hash(newRefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("hash new refresh: %w", err)
	}

	refreshDomain := domain.RefreshToken{
		UserID:    user.ID,
		TokenHash: newHash,
		ExpiresAt: time.Now().Add(s.refreshTTL),
		CreatedAt: time.Now(),
	}
	if err := s.authRepository.CreateToken(ctx, refreshDomain); err != nil {
		return "", "", fmt.Errorf("store new refresh: %w", err)
	}

	return accessToken, newRefreshToken, nil
}
