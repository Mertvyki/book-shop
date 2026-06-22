package reviews_transport_http

import (
	"errors"
	"net/http"

	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
	reviews_service "github.com/Mertvyki/book-shop/internal/features/reviews/service"
)

func (h *ReviewsHTTPHandler) UpsertReview(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, ok := r.Context().Value(core_http_middleware.UserIDKey).(int)
	if !ok {
		responseHandler.ErrorResponse(unauthorizedErr, "unauthorized")
		return
	}

	bookID, err := core_http_request.GetIntPathValue(r, "bookId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bookId path value")
		return
	}

	var req UpsertReviewRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &req); err != nil {
		responseHandler.ErrorResponse(err, "invalid request body")
		return
	}

	payload := reviews_service.UpsertReviewPayload{
		BookID: bookID,
		UserID: userID,
		Rating: req.Rating,
		Title:  req.Title,
		Body:   req.Body,
	}

	review, err := h.reviewService.UpsertReview(ctx, payload)
	if err != nil {
		if errors.Is(err, core_errors.ErrInvalidArgument) {
			responseHandler.ErrorResponse(err, "invalid rating")
			return
		}
		responseHandler.ErrorResponse(err, "failed to save review")
		return
	}

	responseHandler.JSONResponse(reviewDTOFromDomain(review), http.StatusOK)
}
