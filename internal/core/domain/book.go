package domain

import "time"

type Book struct {
	ID      int
	Version int

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

func NewBook(
	id int,
	version int,
	title string,
	author string,
	description *string,
	isbn *string,
	price float64,
	bookType string,
	stockQuantity *int,
	fileUrl *string,
	coverImageUrl *string,
	createdAt time.Time,
) Book {
	return Book{
		ID:            id,
		Version:       version,
		Title:         title,
		Author:        author,
		Description:   description,
		ISBN:          isbn,
		Price:         price,
		BookType:      bookType,
		StockQuantity: stockQuantity,
		FileURL:       fileUrl,
		CoverImageURL: coverImageUrl,
		CreatedAt:     createdAt,
	}
}

func NewBookUninitialized(
	title string,
	author string,
	description *string,
	isbn *string,
	price float64,
	bookType string,
	stockQuantity *int,
	fileUrl *string,
	coverImageUrl *string,
) Book {
	return NewBook(UninitializedID, UninitializedVersion, title, author, description, isbn, price, bookType, stockQuantity, fileUrl, coverImageUrl, time.Now().UTC())
}
