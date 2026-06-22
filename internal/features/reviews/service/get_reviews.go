package reviews_service

import (
	"context"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

func (s *ReviewService) GetBookReviews(ctx context.Context, bookID int) ([]domain.Review, error) {
	reviews, err := s.reviewRepository.GetBookReviews(ctx, bookID)
	if err != nil {
		return nil, fmt.Errorf("get book reviews: %w", err)
	}
	return reviews, nil
}

func (s *ReviewService) GetUserReview(ctx context.Context, bookID, userID int) (domain.Review, error) {
	review, err := s.reviewRepository.GetUserReview(ctx, bookID, userID)
	if err != nil {
		return domain.Review{}, fmt.Errorf("get user review: %w", err)
	}
	return review, nil
}
