package addresses_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	core_http_types "github.com/Mertvyki/book-shop/internal/core/transport/http/types"
	addresses_service "github.com/Mertvyki/book-shop/internal/features/addresses/service"
)

type PatchAddressRequest struct {
	StreetAddress core_http_types.Nullable[string] `json:"street_address"`
	City          core_http_types.Nullable[string] `json:"city"`
	PostalCode    core_http_types.Nullable[string] `json:"postal_code"`
	Country       core_http_types.Nullable[string] `json:"country"`
	IsDefault     core_http_types.Nullable[bool]   `json:"is_default"`
}

func (r PatchAddressRequest) ToService() addresses_service.PatchAddressPayload {
	return addresses_service.PatchAddressPayload{
		StreetAddress: r.StreetAddress.ToDomain().Value,
		City:          r.City.ToDomain().Value,
		PostalCode:    r.PostalCode.ToDomain().Value,
		Country:       r.Country.ToDomain().Value,
		IsDefault:     r.IsDefault.ToDomain().Value,
	}
}

type PatchAddressResponse AddressDTOResponse

func (h *AddressesHTTPHandler) PatchAddress(rw http.ResponseWriter, r *http.Request) {
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

	var request PatchAddressRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	address, err := h.addressesService.PatchAddress(ctx, userID, addrID, request.ToService())
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch address")
		return
	}

	response := PatchAddressResponse(addressDTOFromDomain(address))
	responseHandler.JSONResponse(response, http.StatusOK)
}
