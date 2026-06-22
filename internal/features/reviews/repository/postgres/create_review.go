package reviews_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	core_postgres_pool "github.com/Mertvyki/book-shop/internal/core/repository/postgres/pool"
)

func (r *ReviewsRepository) CreateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m ReviewModel
	err := r.pool.QueryRow(ctx, `
		INSERT INTO bookshop.reviews (book_id, user_id, rating, title, body)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, version, book_id, user_id, rating, title, body, created_at, updated_at
	`, review.BookID, review.UserID, review.Rating, review.Title, review.Body).Scan(
		&m.ID, &m.Version, &m.BookID, &m.UserID, &m.Rating, &m.Title, &m.Body, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrUniqueViolation) {
			return domain.Review{}, fmt.Errorf("create review: %w", core_errors.ErrConflict)
		}
		return domain.Review{}, fmt.Errorf("insert review: %w", err)
	}

	return m.ToDomain(), nil
}

func (r *ReviewsRepository) GetUserReview(ctx context.Context, bookID, userID int) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m ReviewModel
	err := r.pool.QueryRow(ctx, `
		SELECT id, version, book_id, user_id, rating, title, body, created_at, updated_at
		FROM bookshop.reviews
		WHERE book_id = $1 AND user_id = $2
	`, bookID, userID).Scan(
		&m.ID, &m.Version, &m.BookID, &m.UserID, &m.Rating, &m.Title, &m.Body, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Review{}, fmt.Errorf("get user review: %w", core_errors.ErrNotFound)
		}
		return domain.Review{}, fmt.Errorf("scan user review: %w", err)
	}

	return m.ToDomain(), nil
}

func (r *ReviewsRepository) UpdateReview(ctx context.Context, review domain.Review) (domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	var m ReviewModel
	err := r.pool.QueryRow(ctx, `
		UPDATE bookshop.reviews
		SET rating = $1, title = $2, body = $3, version = version + 1, updated_at = NOW()
		WHERE id = $4 AND version = $5
		RETURNING id, version, book_id, user_id, rating, title, body, created_at, updated_at
	`, review.Rating, review.Title, review.Body, review.ID, review.Version).Scan(
		&m.ID, &m.Version, &m.BookID, &m.UserID, &m.Rating, &m.Title, &m.Body, &m.CreatedAt, &m.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Review{}, fmt.Errorf("update review: %w", core_errors.ErrNotFound)
		}
		return domain.Review{}, fmt.Errorf("update review: %w", err)
	}

	return m.ToDomain(), nil
}

func (r *ReviewsRepository) GetBookReviews(ctx context.Context, bookID int) ([]domain.Review, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	rows, err := r.pool.Query(ctx, `
		SELECT r.id, r.version, r.book_id, r.user_id, r.rating, r.title, r.body, r.created_at, r.updated_at,
		       u.full_name
		FROM bookshop.reviews r
		JOIN bookshop.users u ON u.id = r.user_id
		WHERE r.book_id = $1
		ORDER BY r.created_at DESC
	`, bookID)
	if err != nil {
		return nil, fmt.Errorf("query reviews: %w", err)
	}
	defer rows.Close()

	reviews := make([]domain.Review, 0)
	for rows.Next() {
		var m ReviewModel
		var userName string
		err = rows.Scan(
			&m.ID, &m.Version, &m.BookID, &m.UserID, &m.Rating, &m.Title, &m.Body, &m.CreatedAt, &m.UpdatedAt,
			&userName,
		)
		if err != nil {
			return nil, fmt.Errorf("scan review: %w", err)
		}
		review := m.ToDomain()
		review.UserName = userName
		reviews = append(reviews, review)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate reviews: %w", err)
	}

	return reviews, nil
}

func (r *ReviewsRepository) DeleteReview(ctx context.Context, reviewID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	tag, err := r.pool.Exec(ctx, `DELETE FROM bookshop.reviews WHERE id = $1`, reviewID)
	if err != nil {
		return fmt.Errorf("delete review: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("delete review: %w", core_errors.ErrNotFound)
	}
	return nil
}
