package reviews_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type ReviewModel struct {
	ID        int
	Version   int
	BookID    int
	UserID    int
	Rating    int
	Title     *string
	Body      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m ReviewModel) ToDomain() domain.Review {
	return domain.Review{
		ID:        m.ID,
		Version:   m.Version,
		BookID:    m.BookID,
		UserID:    m.UserID,
		Rating:    m.Rating,
		Title:     m.Title,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
