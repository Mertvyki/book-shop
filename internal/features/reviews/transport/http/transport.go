package reviews_transport_http

import (
	"context"
	"errors"
	"net/http"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_http_middleware "github.com/Mertvyki/book-shop/internal/core/transport/http/middleware"
	core_http_server "github.com/Mertvyki/book-shop/internal/core/transport/http/server"
	reviews_service "github.com/Mertvyki/book-shop/internal/features/reviews/service"
)

var unauthorizedErr = errors.New("unauthorized")

type ReviewsHTTPHandler struct {
	reviewService  ReviewService
	authMiddleware core_http_middleware.Middleware
}

type ReviewService interface {
	UpsertReview(ctx context.Context, payload reviews_service.UpsertReviewPayload) (domain.Review, error)
	GetBookReviews(ctx context.Context, bookID int) ([]domain.Review, error)
	GetUserReview(ctx context.Context, bookID, userID int) (domain.Review, error)
}

func NewReviewsHTTPHandler(
	reviewService ReviewService,
	authMiddleware core_http_middleware.Middleware,
) *ReviewsHTTPHandler {
	return &ReviewsHTTPHandler{
		reviewService:  reviewService,
		authMiddleware: authMiddleware,
	}
}

func (h *ReviewsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/books/{bookId}/reviews",
			Handler: h.GetBookReviews,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/{bookId}/reviews/mine",
			Handler: h.GetUserReview,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
		{
			Method:  http.MethodPut,
			Path:    "/books/{bookId}/reviews",
			Handler: h.UpsertReview,
			Middleware: []core_http_middleware.Middleware{
				h.authMiddleware,
			},
		},
	}
}
