package users_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type UserDTOResponse struct {
	ID          int       `json:"id"`
	Version     int       `json:"version"`
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	PhoneNumber *string   `json:"phone_number"`
	Role        string    `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		Email:       user.Email,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	userDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		userDTO[i] = userDTOFromDomain(user)
	}

	return userDTO
}
