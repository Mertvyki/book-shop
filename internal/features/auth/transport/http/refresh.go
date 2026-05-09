package auth_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (h *AuthHTTPHandler) Refresh(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request RefreshRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	access, refresh, err := h.authService.Refresh(ctx, request.RefreshToken)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to refresh token",
		)

		return
	}

	response := RefreshDTOResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}
