package reviews_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type ReviewDTOResponse struct {
	ID        int       `json:"id"`
	Version   int       `json:"version"`
	BookID    int       `json:"book_id"`
	UserID    int       `json:"user_id"`
	Rating    int       `json:"rating"`
	Title     *string   `json:"title"`
	Body      *string   `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserName  string    `json:"user_name"`
}

func reviewDTOFromDomain(review domain.Review) ReviewDTOResponse {
	return ReviewDTOResponse{
		ID:        review.ID,
		Version:   review.Version,
		BookID:    review.BookID,
		UserID:    review.UserID,
		Rating:    review.Rating,
		Title:     review.Title,
		Body:      review.Body,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		UserName:  review.UserName,
	}
}

func reviewsDTOFromDomains(reviews []domain.Review) []ReviewDTOResponse {
	dtos := make([]ReviewDTOResponse, len(reviews))
	for i, r := range reviews {
		dtos[i] = reviewDTOFromDomain(r)
	}
	return dtos
}

type UpsertReviewRequest struct {
	Rating int     `json:"rating" validate:"required,min=1,max=5"`
	Title  *string `json:"title"`
	Body   *string `json:"body"`
}
