package auth_transport_http

import (
	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type LoginDTOResponse struct {
	AccessToken  string
	RefreshToken string
	User         domain.User
}

type RefreshDTOResponse struct {
	AccessToken  string
	RefreshToken string
}

func RefreshTokenDTOFromDomain(accessToken string, refreshToken string) RefreshDTOResponse {
	return RefreshDTOResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func loginDTOFromDomain(accessToken string, refreshToken string, user domain.User) LoginDTOResponse {
	return LoginDTOResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}
}
