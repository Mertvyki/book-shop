package domain

import "time"

type CartItem struct {
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

func NewCartItem(id, version, userID, bookID, quantity int, addedAt time.Time, title string, price float64, coverImageKey *string, bookType string) CartItem {
	return CartItem{
		ID:            id,
		Version:       version,
		UserID:        userID,
		BookID:        bookID,
		Quantity:      quantity,
		AddedAt:       addedAt,
		Title:         title,
		Price:         price,
		CoverImageKey: coverImageKey,
		BookType:      bookType,
	}
}

func NewCartItemUninitialized(userID, bookID, quantity int) CartItem {
	return CartItem{
		ID:       UninitializedID,
		Version:  UninitializedVersion,
		UserID:   userID,
		BookID:   bookID,
		Quantity: quantity,
		AddedAt:  time.Now().UTC(),
	}
}
