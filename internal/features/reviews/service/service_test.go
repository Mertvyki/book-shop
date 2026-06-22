package reviews_service

import (
	"context"
	"errors"
	"testing"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockReviewRepo struct{ mock.Mock }

func (m *mockReviewRepo) CreateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	args := m.Called(ctx, review)
	return args.Get(0).(domain.Review), args.Error(1)
}

func (m *mockReviewRepo) GetUserReview(ctx context.Context, bookID, userID int) (domain.Review, error) {
	args := m.Called(ctx, bookID, userID)
	return args.Get(0).(domain.Review), args.Error(1)
}

func (m *mockReviewRepo) UpdateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	args := m.Called(ctx, review)
	return args.Get(0).(domain.Review), args.Error(1)
}

func (m *mockReviewRepo) GetBookReviews(ctx context.Context, bookID int) ([]domain.Review, error) {
	args := m.Called(ctx, bookID)
	return args.Get(0).([]domain.Review), args.Error(1)
}

func (m *mockReviewRepo) DeleteReview(ctx context.Context, reviewID int) error {
	args := m.Called(ctx, reviewID)
	return args.Error(0)
}

func newTestReviewService(t *testing.T) (*ReviewService, *mockReviewRepo) {
	t.Helper()
	repo := &mockReviewRepo{}
	svc := NewReviewService(repo)
	return svc, repo
}

func TestUpsertReview_Create(t *testing.T) {
	svc, repo := newTestReviewService(t)

	payload := UpsertReviewPayload{
		BookID: 1,
		UserID: 2,
		Rating: 4,
		Title:  strPtr("Good book"),
		Body:   strPtr("Really enjoyed it"),
	}

	// No existing review → returns ErrNotFound
	repo.On("GetUserReview", mock.Anything, 1, 2).Return(domain.Review{}, core_errors.ErrNotFound)

	expected := domain.Review{
		ID:     1,
		BookID: 1,
		UserID: 2,
		Rating: 4,
		Title:  strPtr("Good book"),
		Body:   strPtr("Really enjoyed it"),
	}
	repo.On("CreateReview", mock.Anything, mock.MatchedBy(func(r domain.Review) bool {
		return r.BookID == 1 && r.UserID == 2 && r.Rating == 4
	})).Return(expected, nil)

	result, err := svc.UpsertReview(context.Background(), payload)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}

func TestUpsertReview_Update(t *testing.T) {
	svc, repo := newTestReviewService(t)

	payload := UpsertReviewPayload{
		BookID: 1,
		UserID: 2,
		Rating: 3,
		Body:   strPtr("Updated review"),
	}

	existing := domain.Review{
		ID:      1,
		Version: 1,
		BookID:  1,
		UserID:  2,
		Rating:  5,
		Title:   strPtr("Original"),
		Body:    strPtr("Original text"),
	}

	// Existing review found → update path
	repo.On("GetUserReview", mock.Anything, 1, 2).Return(existing, nil)

	updated := domain.Review{
		ID:      1,
		Version: 2,
		BookID:  1,
		UserID:  2,
		Rating:  3,
		Body:    strPtr("Updated review"),
	}
	repo.On("UpdateReview", mock.Anything, mock.MatchedBy(func(r domain.Review) bool {
		return r.ID == 1 && r.Rating == 3 && r.Body != nil && *r.Body == "Updated review"
	})).Return(updated, nil)

	result, err := svc.UpsertReview(context.Background(), payload)

	assert.NoError(t, err)
	assert.Equal(t, updated, result)
	repo.AssertExpectations(t)
}

func TestUpsertReview_InvalidRating(t *testing.T) {
	svc, _ := newTestReviewService(t)

	payload := UpsertReviewPayload{
		BookID: 1,
		UserID: 2,
		Rating: 6,
	}

	_, err := svc.UpsertReview(context.Background(), payload)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestUpsertReview_RatingBelowMin(t *testing.T) {
	svc, _ := newTestReviewService(t)

	payload := UpsertReviewPayload{
		BookID: 1,
		UserID: 2,
		Rating: 0,
	}

	_, err := svc.UpsertReview(context.Background(), payload)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestGetBookReviews(t *testing.T) {
	svc, repo := newTestReviewService(t)

	expected := []domain.Review{
		{ID: 1, BookID: 1, UserID: 2, Rating: 5, UserName: "Alice"},
		{ID: 2, BookID: 1, UserID: 3, Rating: 4, UserName: "Bob"},
	}

	repo.On("GetBookReviews", mock.Anything, 1).Return(expected, nil)

	result, err := svc.GetBookReviews(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}

func TestGetBookReviews_Empty(t *testing.T) {
	svc, repo := newTestReviewService(t)

	repo.On("GetBookReviews", mock.Anything, 1).Return([]domain.Review{}, nil)

	result, err := svc.GetBookReviews(context.Background(), 1)

	assert.NoError(t, err)
	assert.Empty(t, result)
	repo.AssertExpectations(t)
}

func TestGetUserReview_Success(t *testing.T) {
	svc, repo := newTestReviewService(t)

	expected := domain.Review{
		ID:     1,
		BookID: 1,
		UserID: 2,
		Rating: 5,
	}

	repo.On("GetUserReview", mock.Anything, 1, 2).Return(expected, nil)

	result, err := svc.GetUserReview(context.Background(), 1, 2)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	repo.AssertExpectations(t)
}

func TestGetUserReview_NotFound(t *testing.T) {
	svc, repo := newTestReviewService(t)

	repo.On("GetUserReview", mock.Anything, 1, 2).Return(domain.Review{}, core_errors.ErrNotFound)

	_, err := svc.GetUserReview(context.Background(), 1, 2)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrNotFound))
	repo.AssertExpectations(t)
}

func strPtr(s string) *string {
	return &s
}
