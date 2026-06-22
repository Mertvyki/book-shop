package reviews_transport_http

import (
	"net/http"

	core_logger "github.com/Mertvyki/book-shop/internal/core/logger"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_request "github.com/Mertvyki/book-shop/internal/core/transport/http/request"
	core_http_response "github.com/Mertvyki/book-shop/internal/core/transport/http/response"
)

func (h *ReviewsHTTPHandler) GetBookReviews(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	bookID, err := core_http_request.GetIntPathValue(r, "bookId")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get bookId path value")
		return
	}

	reviews, err := h.reviewService.GetBookReviews(ctx, bookID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get reviews")
		return
	}

	responseHandler.JSONResponse(reviewsDTOFromDomains(reviews), http.StatusOK)
}

func (h *ReviewsHTTPHandler) GetUserReview(rw http.ResponseWriter, r *http.Request) {
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

	review, err := h.reviewService.GetUserReview(ctx, bookID, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user review")
		return
	}

	responseHandler.JSONResponse(reviewDTOFromDomain(review), http.StatusOK)
}
