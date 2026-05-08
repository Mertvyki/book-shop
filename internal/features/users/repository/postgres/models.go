package users_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type UserModel struct {
	ID      int
	Version int

	Email        string
	PasswordHash string
	FullName     string
	PhoneNumber  *string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func userDomainsFromModels(users []UserModel) []domain.User {
	userDomains := make([]domain.User, len(users))

	for i, user := range users {
		userDomains[i] = domain.NewUser(
			user.ID,
			user.Version,
			user.Email,
			user.PasswordHash,
			user.FullName,
			user.PhoneNumber,
			user.Role,
			user.CreatedAt,
			user.UpdatedAt,
		)
	}

	return userDomains
}
