package cart_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type CartService struct {
	cartRepository CartRepository
	bookRepository BookRepository
}

type CartRepository interface {
	GetCart(ctx context.Context, userID int) ([]domain.CartItem, error)
	AddItem(ctx context.Context, item domain.CartItem) (domain.CartItem, error)
	UpdateItem(ctx context.Context, itemID, userID, quantity int) (domain.CartItem, error)
	RemoveItem(ctx context.Context, itemID, userID int) error
	ClearCart(ctx context.Context, userID int) error
}

type BookRepository interface {
	GetBook(ctx context.Context, id int) (domain.Book, error)
}

func NewCartService(cartRepository CartRepository, bookRepository BookRepository) *CartService {
	return &CartService{
		cartRepository: cartRepository,
		bookRepository: bookRepository,
	}
}
