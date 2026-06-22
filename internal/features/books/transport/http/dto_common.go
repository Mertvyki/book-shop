package books_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type AuthorDTOResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Bio       *string   `json:"bio"`
	BirthYear *int      `json:"birth_year"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoryDTOResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	Description *string `json:"description"`
}

type PublisherDTOResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type BookDTOResponse struct {
	ID            int                   `json:"id"`
	Version       int                   `json:"version"`
	Title         string                `json:"title"`
	Description   *string               `json:"description"`
	ISBN          *string               `json:"isbn"`
	Price         float64               `json:"price"`
	BookType      string                `json:"book_type"`
	StockQuantity *int                  `json:"stock_quantity"`
	FileURL       *string               `json:"file_key"`
	CoverImageURL *string               `json:"cover_image_key"`
	Publisher     *PublisherDTOResponse `json:"publisher"`
	Authors       []AuthorDTOResponse   `json:"authors"`
	Categories    []CategoryDTOResponse `json:"categories"`
	AvgRating     float64               `json:"avg_rating"`
	CreatedAt     time.Time             `json:"created_at"`
}

func bookDTOFromDomain(book domain.Book) BookDTOResponse {
	authors := make([]AuthorDTOResponse, len(book.Authors))
	for i, a := range book.Authors {
		authors[i] = AuthorDTOResponse{
			ID:        a.ID,
			Name:      a.Name,
			Bio:       a.Bio,
			BirthYear: a.BirthYear,
			CreatedAt: a.CreatedAt,
		}
	}

	categories := make([]CategoryDTOResponse, len(book.Categories))
	for i, c := range book.Categories {
		categories[i] = CategoryDTOResponse{
			ID:          c.ID,
			Name:        c.Name,
			Slug:        c.Slug,
			Description: c.Description,
		}
	}

	var publisher *PublisherDTOResponse
	if book.Publisher != nil {
		publisher = &PublisherDTOResponse{
			ID:   book.Publisher.ID,
			Name: book.Publisher.Name,
		}
	}

	return BookDTOResponse{
		ID:            book.ID,
		Version:       book.Version,
		Title:         book.Title,
		Description:   book.Description,
		ISBN:          book.ISBN,
		Price:         book.Price,
		BookType:      book.BookType,
		StockQuantity: book.StockQuantity,
		FileURL:       book.FileURL,
		CoverImageURL: book.CoverImageURL,
		Publisher:     publisher,
		Authors:       authors,
		Categories:    categories,
		AvgRating:     book.AvgRating,
		CreatedAt:     book.CreatedAt,
	}
}

func booksDTOFromDomains(books []domain.Book) []BookDTOResponse {
	bookDTO := make([]BookDTOResponse, len(books))
	for i, book := range books {
		bookDTO[i] = bookDTOFromDomain(book)
	}

	return bookDTO
}
