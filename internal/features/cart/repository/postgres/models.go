package cart_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type CartItemModel struct {
	ID            int
	Version       int
	UserID        int
	BookID        int
	Quantity      int
	AddedAt       time.Time
	Title         string
	Price         float64
	CoverImageKey *string
	BookType      string
}

func (m CartItemModel) ToDomain() domain.CartItem {
	return domain.CartItem{
		ID:            m.ID,
		Version:       m.Version,
		UserID:        m.UserID,
		BookID:        m.BookID,
		Quantity:      m.Quantity,
		AddedAt:       m.AddedAt,
		Title:         m.Title,
		Price:         m.Price,
		CoverImageKey: m.CoverImageKey,
		BookType:      m.BookType,
	}
}
