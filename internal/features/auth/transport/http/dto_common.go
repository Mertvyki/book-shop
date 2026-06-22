package auth_transport_http

type LoginDTOResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		FullName string `json:"full_name"`
		Role     string `json:"role"`
	} `json:"user"`
}

type RefreshDTOResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func RefreshTokenDTOFromDomain(accessToken string, refreshToken string) RefreshDTOResponse {
	return RefreshDTOResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func loginDTOFromDomain(accessToken string, refreshToken string, userID int, email, fullName, role string) LoginDTOResponse {
	return LoginDTOResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: struct {
			ID       int    `json:"id"`
			Email    string `json:"email"`
			FullName string `json:"full_name"`
			Role     string `json:"role"`
		}{
			ID:       userID,
			Email:    email,
			FullName: fullName,
			Role:     role,
		},
	}
}
