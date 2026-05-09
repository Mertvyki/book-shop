package auth_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type RefreshTokenModel struct {
	ID        int
	UserID    int
	TokenHash string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func (m *RefreshTokenModel) ToDomain() domain.RefreshToken {
	return domain.RefreshToken{
		ID:        m.ID,
		UserID:    m.UserID,
		TokenHash: m.TokenHash,
		ExpiresAt: m.ExpiresAt,
		CreatedAt: m.CreatedAt,
	}
}
