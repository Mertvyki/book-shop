package addresses_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

type GetAddressesResponse []AddressDTOResponse

func (h *AddressesHTTPHandler) GetAddresses(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "userId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id")
		return
	}

	if err := checkOwnerOrAdmin(r, userID); err != nil {
		responseHandler.ErrorResponse(err, "forbidden")
		return
	}

	addresses, err := h.addressesService.GetAddresses(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get addresses")
		return
	}

	response := GetAddressesResponse(addressesDTOFromDomains(addresses))
	responseHandler.JSONResponse(response, http.StatusOK)
}
