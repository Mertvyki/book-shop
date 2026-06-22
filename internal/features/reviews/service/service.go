package reviews_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type ReviewService struct {
	reviewRepository ReviewRepository
}

type ReviewRepository interface {
	CreateReview(ctx context.Context, review domain.Review) (domain.Review, error)
	GetUserReview(ctx context.Context, bookID, userID int) (domain.Review, error)
	UpdateReview(ctx context.Context, review domain.Review) (domain.Review, error)
	GetBookReviews(ctx context.Context, bookID int) ([]domain.Review, error)
	DeleteReview(ctx context.Context, reviewID int) error
}

func NewReviewService(reviewRepository ReviewRepository) *ReviewService {
	return &ReviewService{
		reviewRepository: reviewRepository,
	}
}
