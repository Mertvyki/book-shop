package books_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type BookDTOResponse struct {
	ID            int       `json:"id"`
	Version       int       `json:"version"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   *string   `json:"description"`
	ISBN          *string   `json:"isbn"`
	Price         float64   `json:"price"`
	BookType      string    `json:"book_type"`
	StockQuantity *int      `json:"stock_quantity"`
	FileURL       *string   `json:"file_key"`
	CoverImageURL *string   `json:"cover_image_key"`
	CreatedAt     time.Time `json:"created_at"`
}

func bookDTOFromDomain(
	book domain.Book,
) BookDTOResponse {
	return BookDTOResponse{
		ID:            book.ID,
		Version:       book.Version,
		Title:         book.Title,
		Author:        book.Author,
		Description:   book.Description,
		ISBN:          book.ISBN,
		Price:         book.Price,
		BookType:      book.BookType,
		StockQuantity: book.StockQuantity,
		FileURL:       book.FileURL,
		CoverImageURL: book.CoverImageURL,
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
