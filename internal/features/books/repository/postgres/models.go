package book_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type BookModel struct {
	ID            int
	Version       int
	Title         string
	Description   *string
	ISBN          *string
	Price         float64
	BookType      string
	StockQuantity *int
	FileURL       *string
	CoverImageURL *string
	PublisherID   *int
	AvgRating     float64
	CreatedAt     time.Time
}

func (m BookModel) ToDomain() domain.Book {
	return domain.Book{
		ID:            m.ID,
		Version:       m.Version,
		Title:         m.Title,
		Description:   m.Description,
		ISBN:          m.ISBN,
		Price:         m.Price,
		BookType:      m.BookType,
		StockQuantity: m.StockQuantity,
		FileURL:       m.FileURL,
		CoverImageURL: m.CoverImageURL,
		PublisherID:   m.PublisherID,
		AvgRating:     m.AvgRating,
		CreatedAt:     m.CreatedAt,
	}
}

type AuthorModel struct {
	ID        int
	Name      string
	Bio       *string
	BirthYear *int
	CreatedAt time.Time
}

func (m AuthorModel) ToDomain() domain.Author {
	return domain.Author{
		ID:        m.ID,
		Name:      m.Name,
		Bio:       m.Bio,
		BirthYear: m.BirthYear,
		CreatedAt: m.CreatedAt,
	}
}

type CategoryModel struct {
	ID          int
	Name        string
	Slug        string
	Description *string
}

func (m CategoryModel) ToDomain() domain.Category {
	return domain.NewCategory(m.ID, m.Name, m.Slug, m.Description)
}

type PublisherModel struct {
	ID   int
	Name string
}

func (m PublisherModel) ToDomain() domain.Publisher {
	return domain.NewPublisher(m.ID, m.Name)
}
