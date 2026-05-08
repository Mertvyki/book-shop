package users_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Email       string  `json:"email" validate:"required"`
	Password    string  `json:"password" validate:"required,min=6"`
	FullName    string  `json:"full_name" validate:"required,min=15,max=50"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	createdUser, err := h.usersService.CreateUser(ctx, request.Email, request.Password, request.FullName, request.PhoneNumber)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create user",
		)

		return
	}

	response := CreateUserResponse(userDTOFromDomain(createdUser))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
