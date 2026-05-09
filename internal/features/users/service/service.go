package user_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type UserService struct {
	usersRepository UserRepository
	hasher          PasswordHash
	tokenManager    TokenManager
}

type PatchUserPayload struct {
	Email       domain.Nullable[string]
	FullName    domain.Nullable[string]
	PhoneNumber domain.Nullable[string]
	Password    domain.Nullable[string]
	OldPassword domain.Nullable[string]
}

type UserRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)

	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)

	PatchUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)

	DeleteUser(
		ctx context.Context,
		id int,
	) error

	ByEmail(
		ctx context.Context,
		email string,
	) (*domain.User, error)
	ByID(
		ctx context.Context,
		id int,
	) (*domain.User, error)
}

type PasswordHash interface {
	Hash(password string) (string, error)
	Compare(hashedPassword string, password string) bool
}

type TokenManager interface {
	GenerateAccessToken(userID int, role string) (string, error)
}

func NewUserService(userRepository UserRepository, hasher PasswordHash, tokenManager TokenManager) *UserService {
	return &UserService{
		usersRepository: userRepository,
		hasher:          hasher,
		tokenManager:    tokenManager,
	}
}
