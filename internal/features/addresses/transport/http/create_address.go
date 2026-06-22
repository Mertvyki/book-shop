package addresses_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	addresses_service "github.com/Mertvyki/book-shop/internal/features/addresses/service"
)

type CreateAddressRequest struct {
	StreetAddress string `json:"street_address" validate:"required,max=255"`
	City          string `json:"city" validate:"required,max=100"`
	PostalCode    string `json:"postal_code" validate:"required,max=20"`
	Country       string `json:"country" validate:"omitempty,max=100"`
	IsDefault     bool   `json:"is_default"`
}

func (r CreateAddressRequest) ToService() addresses_service.CreateAddressPayload {
	return addresses_service.CreateAddressPayload{
		StreetAddress: r.StreetAddress,
		City:          r.City,
		PostalCode:    r.PostalCode,
		Country:       r.Country,
		IsDefault:     r.IsDefault,
	}
}

type CreateAddressResponse AddressDTOResponse

func (h *AddressesHTTPHandler) CreateAddress(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	pathUserID, err := core_http_request.GetIntPathValue(r, "userId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id")
		return
	}

	if err := checkOwnerOrAdmin(r, pathUserID); err != nil {
		responseHandler.ErrorResponse(err, "forbidden")
		return
	}

	var request CreateAddressRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}

	address, err := h.addressesService.CreateAddress(ctx, pathUserID, request.ToService())
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create address")
		return
	}

	response := CreateAddressResponse(addressDTOFromDomain(address))
	responseHandler.JSONResponse(response, http.StatusOK)
}
