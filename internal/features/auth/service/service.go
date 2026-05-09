package auth_service

import (
	"context"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_security "github.com/Mertvyki/book-shop/internal/core/security"
)

type AuthService struct {
	authRepository AuthRepository
	userRepository UserRepository
	hasher         PasswordHasher
	jwtManager     *core_security.JWTManager
	refreshGen     *core_security.RefreshTokenService
	refreshTTL     time.Duration
}

type UserRepository interface {
	ByEmail(
		ctx context.Context,
		email string,
	) (*domain.User, error)
	ByID(
		ctx context.Context,
		id int,
	) (*domain.User, error)
}

type AuthRepository interface {
	CreateToken(
		ctx context.Context,
		token domain.RefreshToken,
	) error
	DeleteToken(
		ctx context.Context,
		id int,
	) error
	GetTokenByHash(
		ctx context.Context,
		hash string,
	) (domain.RefreshToken, error)
}

type PasswordHasher interface {
	Hash(
		password string,
	) (string, error)
	Compare(
		hashedPassword string,
		password string,
	) bool
}

func NewAuthService(
	authRepository AuthRepository,
	userRepository UserRepository,
	hasher PasswordHasher,
	jwtManager *core_security.JWTManager,
	refreshGen *core_security.RefreshTokenService,
	refreshTTL time.Duration,
) *AuthService {
	return &AuthService{
		authRepository: authRepository,
		userRepository: userRepository,
		hasher:         hasher,
		jwtManager:     jwtManager,
		refreshGen:     refreshGen,
		refreshTTL:     refreshTTL,
	}
}
