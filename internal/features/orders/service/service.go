package orders_service

import (
	"context"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	orders_postgres_repository "github.com/Mertvyki/book-shop/internal/features/orders/repository/postgres"
)

type OrderService struct {
	orderRepository   OrderRepository
	cartRepository    CartRepository
	addressRepository AddressRepository
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, userID int, shippingAddressID *int, items []domain.CartItem, deliveryCost float64) (orders_postgres_repository.CreateOrderResult, error)
	GetOrder(ctx context.Context, orderID, userID int) (orders_postgres_repository.OrderWithItems, error)
	GetOrders(ctx context.Context, userID int, limit, offset int) ([]domain.Order, error)
	CountOrders(ctx context.Context, userID int) (int, error)
	GetAllOrders(ctx context.Context, limit, offset int) ([]domain.Order, error)
	CountAllOrders(ctx context.Context) (int, error)
	UpdateOrderStatus(ctx context.Context, orderID, version int, status string) error
	HasUserPurchasedBook(ctx context.Context, userID, bookID int) (bool, error)
}

type CartRepository interface {
	GetCart(ctx context.Context, userID int) ([]domain.CartItem, error)
	ClearCart(ctx context.Context, userID int) error
}

type AddressRepository interface {
	GetAddress(ctx context.Context, userID, addrID int) (domain.Address, error)
}

func NewOrderService(
	orderRepository OrderRepository,
	cartRepository CartRepository,
	addressRepository AddressRepository,
) *OrderService {
	return &OrderService{
		orderRepository:   orderRepository,
		cartRepository:    cartRepository,
		addressRepository: addressRepository,
	}
}
