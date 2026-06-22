package auth_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	"github.com/Mertvyki/book-shop/internal/core/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserRepo struct{ mock.Mock }

func (m *mockUserRepo) ByEmail(ctx context.Context, email string) (*domain.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *mockUserRepo) ByID(ctx context.Context, id int) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

type mockAuthRepo struct{ mock.Mock }

func (m *mockAuthRepo) CreateToken(ctx context.Context, token domain.RefreshToken) error {
	return m.Called(ctx, token).Error(0)
}

func (m *mockAuthRepo) DeleteToken(ctx context.Context, id int) error {
	return m.Called(ctx, id).Error(0)
}

func (m *mockAuthRepo) GetTokenByHash(ctx context.Context, hash string) (domain.RefreshToken, error) {
	args := m.Called(ctx, hash)
	return args.Get(0).(domain.RefreshToken), args.Error(1)
}

type mockHasher struct{ mock.Mock }

func (m *mockHasher) Hash(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *mockHasher) Compare(hashedPassword, password string) bool {
	args := m.Called(hashedPassword, password)
	return args.Bool(0)
}

func newTestAuthService(t *testing.T) (*AuthService, *mockUserRepo, *mockAuthRepo, *mockHasher) {
	userRepo := &mockUserRepo{}
	authRepo := &mockAuthRepo{}
	hasher := &mockHasher{}
	jwtManager := core_security.NewJWTManager("test-secret", 15*time.Minute, "test")
	refreshGen := core_security.NewRefreshTokenService()
	svc := NewAuthService(authRepo, userRepo, hasher, jwtManager, refreshGen, 24*time.Hour)
	return svc, userRepo, authRepo, hasher
}

func TestLogin_Success(t *testing.T) {
	svc, userRepo, authRepo, hasher := newTestAuthService(t)

	user := &domain.User{
		ID:           1,
		Email:        "test@example.com",
		PasswordHash: "hashed-password",
		Role:         "user",
	}

	userRepo.On("ByEmail", mock.Anything, "test@example.com").Return(user, nil)
	hasher.On("Compare", "hashed-password", "correct-password").Return(true)
	authRepo.On("DeleteToken", mock.Anything, 1).Return(nil)
	authRepo.On("CreateToken", mock.Anything, mock.Anything).Return(nil)

	accessToken, refreshToken, result, err := svc.Login(context.Background(), "test@example.com", "correct-password")

	assert.NoError(t, err)
	assert.NotEmpty(t, accessToken)
	assert.NotEmpty(t, refreshToken)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Email, result.Email)
	userRepo.AssertExpectations(t)
	hasher.AssertExpectations(t)
	authRepo.AssertExpectations(t)
}

func TestLogin_UserNotFound(t *testing.T) {
	svc, userRepo, _, _ := newTestAuthService(t)

	userRepo.On("ByEmail", mock.Anything, "unknown@example.com").Return(nil, core_errors.ErrNotFound)

	_, _, _, err := svc.Login(context.Background(), "unknown@example.com", "password")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrNotFound))
}

func TestLogin_WrongPassword(t *testing.T) {
	svc, userRepo, _, hasher := newTestAuthService(t)

	user := &domain.User{
		ID:           1,
		Email:        "test@example.com",
		PasswordHash: "hashed-password",
		Role:         "user",
	}

	userRepo.On("ByEmail", mock.Anything, "test@example.com").Return(user, nil)
	hasher.On("Compare", "hashed-password", "wrong-password").Return(false)

	_, _, _, err := svc.Login(context.Background(), "test@example.com", "wrong-password")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}
