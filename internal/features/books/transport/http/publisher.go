package books_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *BooksHTTPHandler) CreatePublisher(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var req struct {
		Name string `json:"name"`
	}
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}
	if req.Name == "" {
		responseHandler.ErrorResponse(fmt.Errorf("name is required"), "name is required")
		return
	}

	publisher, err := h.booksService.CreatePublisher(ctx, req.Name)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create publisher")
		return
	}

	responseHandler.JSONResponse(PublisherDTOResponse{ID: publisher.ID, Name: publisher.Name}, http.StatusOK)
}

func (h *BooksHTTPHandler) PatchPublisher(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	publisherID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get publisher id")
		return
	}

	var req struct {
		Name string `json:"name"`
	}
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode request")
		return
	}
	if req.Name == "" {
		responseHandler.ErrorResponse(fmt.Errorf("name is required"), "name is required")
		return
	}

	publisher, err := h.booksService.PatchPublisher(ctx, publisherID, req.Name)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch publisher")
		return
	}

	responseHandler.JSONResponse(PublisherDTOResponse{ID: publisher.ID, Name: publisher.Name}, http.StatusOK)
}

func (h *BooksHTTPHandler) DeletePublisher(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	publisherID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get publisher id")
		return
	}

	if err := h.booksService.DeletePublisher(ctx, publisherID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete publisher")
		return
	}

	responseHandler.NoContentResponse()
}
