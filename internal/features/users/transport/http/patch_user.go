package users_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	core_http_types "github.com/Mertvyki/book-shop/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	Email       core_http_types.Nullable[string] `json:"email"`
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
	Password    core_http_types.Nullable[string] `json:"password"`
	OldPassword core_http_types.Nullable[string] `json:"old_password"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get userID path value",
		)

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate HTTP request",
		)

		return
	}

	patchedUser, err := h.usersService.PatchUser(ctx, userID, request.ToDomain())
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)

		return
	}

	response := PatchUserResponse(userDTOFromDomain(patchedUser))
	responseHandler.JSONResponse(response, http.StatusOK)
}

// userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
// if !ok {
//     resp.ErrorResponse(fmt.Errorf("unauthorized"), "user id not found in context")
//     return
// }
