package auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResponse LoginDTOResponse

// @Summary Login
// @Description Authenticate user by email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginUserRequest true "Credentials"
// @Success 200 {object} LoginDTOResponse
// @Failure 400 {object} map[string]string
// @Router /auth/login [post]
func (h *AuthHTTPHandler) Login(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request LoginUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	access, refresh, user, err := h.authService.Login(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to login user",
		)

		return
	}

	response := LoginUserResponse(loginDTOFromDomain(access, refresh, user.ID, user.Email, user.FullName, user.Role))

	responseHandler.JSONResponse(response, http.StatusOK)
}
