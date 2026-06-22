package cart_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type CartItemDTOResponse struct {
	ID            int       `json:"id"`
	Version       int       `json:"version"`
	UserID        int       `json:"user_id"`
	BookID        int       `json:"book_id"`
	Quantity      int       `json:"quantity"`
	AddedAt       time.Time `json:"added_at"`
	Title         string    `json:"title"`
	Price         float64   `json:"price"`
	CoverImageKey *string   `json:"cover_image_key"`
	BookType      string    `json:"book_type"`
}

func cartItemDTOFromDomain(item domain.CartItem) CartItemDTOResponse {
	return CartItemDTOResponse{
		ID:            item.ID,
		Version:       item.Version,
		UserID:        item.UserID,
		BookID:        item.BookID,
		Quantity:      item.Quantity,
		AddedAt:       item.AddedAt,
		Title:         item.Title,
		Price:         item.Price,
		CoverImageKey: item.CoverImageKey,
		BookType:      item.BookType,
	}
}

func cartItemsDTOFromDomains(items []domain.CartItem) []CartItemDTOResponse {
	dtos := make([]CartItemDTOResponse, len(items))
	for i, item := range items {
		dtos[i] = cartItemDTOFromDomain(item)
	}

	return dtos
}
