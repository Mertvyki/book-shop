package addresses_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *AddressesHTTPHandler) DeleteAddress(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "userId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	if err := checkOwnerOrAdmin(r, userID); err != nil {
		responseHandler.ErrorResponse(err, "forbidden")
		return
	}

	addrID, err := core_http_request.GetIntPathValue(r, "addrId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get addrId path value")
		return
	}

	if err := h.addressesService.DeleteAddress(ctx, userID, addrID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete address")
		return
	}

	responseHandler.NoContentResponse()
}
