package reviews_service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
)

type UpsertReviewPayload struct {
	BookID int
	UserID int
	Rating int
	Title  *string
	Body   *string
}

func (s *ReviewService) UpsertReview(ctx context.Context, payload UpsertReviewPayload) (domain.Review, error) {
	if payload.Rating < 1 || payload.Rating > 5 {
		return domain.Review{}, fmt.Errorf("rating must be between 1 and 5: %w", core_errors.ErrInvalidArgument)
	}

	existing, err := s.reviewRepository.GetUserReview(ctx, payload.BookID, payload.UserID)
	if err != nil {
		if !errors.Is(err, core_errors.ErrNotFound) {
			return domain.Review{}, fmt.Errorf("check existing review: %w", err)
		}

		review := domain.Review{
			BookID: payload.BookID,
			UserID: payload.UserID,
			Rating: payload.Rating,
			Title:  payload.Title,
			Body:   payload.Body,
		}

		created, err := s.reviewRepository.CreateReview(ctx, review)
		if err != nil {
			return domain.Review{}, fmt.Errorf("create review: %w", err)
		}
		return created, nil
	}

	existing.Rating = payload.Rating
	existing.Title = payload.Title
	existing.Body = payload.Body

	updated, err := s.reviewRepository.UpdateReview(ctx, existing)
	if err != nil {
		return domain.Review{}, fmt.Errorf("update review: %w", err)
	}
	return updated, nil
}
