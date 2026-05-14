package book_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type BookModel struct {
	ID            int
	Version       int
	Title         string
	Author        string
	Description   *string
	ISBN          *string
	Price         float64
	BookType      string
	StockQuantity *int
	FileURL       *string
	CoverImageURL *string
	CreatedAt     time.Time
}

func (m BookModel) ToDomain() domain.Book {
	return domain.NewBook(
		m.ID,
		m.Version,
		m.Title,
		m.Author,
		m.Description,
		m.ISBN,
		m.Price,
		m.BookType,
		m.StockQuantity,
		m.FileURL,
		m.CoverImageURL,
		m.CreatedAt,
	)
}
