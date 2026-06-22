package domain

import "time"

type Book struct {
	ID      int
	Version int

	Title         string
	Description   *string
	ISBN          *string
	Price         float64
	BookType      string
	StockQuantity *int
	FileURL       *string
	CoverImageURL *string
	PublisherID   *int
	Publisher     *Publisher
	Authors       []Author
	Categories    []Category
	AvgRating     float64
	CreatedAt     time.Time
}

func NewBook(
	id int,
	version int,
	title string,
	description *string,
	isbn *string,
	price float64,
	bookType string,
	stockQuantity *int,
	fileUrl *string,
	coverImageUrl *string,
	publisherID *int,
	createdAt time.Time,
) Book {
	return Book{
		ID:            id,
		Version:       version,
		Title:         title,
		Description:   description,
		ISBN:          isbn,
		Price:         price,
		BookType:      bookType,
		StockQuantity: stockQuantity,
		FileURL:       fileUrl,
		CoverImageURL: coverImageUrl,
		PublisherID:   publisherID,
		CreatedAt:     createdAt,
	}
}

func NewBookUninitialized(
	title string,
	description *string,
	isbn *string,
	price float64,
	bookType string,
	stockQuantity *int,
	fileUrl *string,
	coverImageUrl *string,
	publisherID *int,
) Book {
	return NewBook(
		UninitializedID, UninitializedVersion,
		title, description, isbn, price, bookType,
		stockQuantity, fileUrl, coverImageUrl, publisherID,
		time.Now().UTC(),
	)
}
