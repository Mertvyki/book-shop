package orders_service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
	core_errors "github.com/Mertvyki/book-shop/internal/core/errrors"
	orders_postgres_repository "github.com/Mertvyki/book-shop/internal/features/orders/repository/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockOrderRepo struct{ mock.Mock }

func (m *mockOrderRepo) CreateOrder(ctx context.Context, userID int, shippingAddressID *int, items []domain.CartItem, deliveryCost float64) (orders_postgres_repository.CreateOrderResult, error) {
	args := m.Called(ctx, userID, shippingAddressID, items, deliveryCost)
	return args.Get(0).(orders_postgres_repository.CreateOrderResult), args.Error(1)
}

func (m *mockOrderRepo) GetOrder(ctx context.Context, orderID, userID int) (orders_postgres_repository.OrderWithItems, error) {
	args := m.Called(ctx, orderID, userID)
	return args.Get(0).(orders_postgres_repository.OrderWithItems), args.Error(1)
}

func (m *mockOrderRepo) GetOrders(ctx context.Context, userID int, limit, offset int) ([]domain.Order, error) {
	args := m.Called(ctx, userID, limit, offset)
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *mockOrderRepo) CountOrders(ctx context.Context, userID int) (int, error) {
	args := m.Called(ctx, userID)
	return args.Int(0), args.Error(1)
}

func (m *mockOrderRepo) GetAllOrders(ctx context.Context, limit, offset int) ([]domain.Order, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *mockOrderRepo) CountAllOrders(ctx context.Context) (int, error) {
	args := m.Called(ctx)
	return args.Int(0), args.Error(1)
}

func (m *mockOrderRepo) UpdateOrderStatus(ctx context.Context, orderID, version int, status string) error {
	return m.Called(ctx, orderID, version, status).Error(0)
}

func (m *mockOrderRepo) HasUserPurchasedBook(ctx context.Context, userID, bookID int) (bool, error) {
	args := m.Called(ctx, userID, bookID)
	return args.Bool(0), args.Error(1)
}

type mockCartRepo struct{ mock.Mock }

func (m *mockCartRepo) GetCart(ctx context.Context, userID int) ([]domain.CartItem, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]domain.CartItem), args.Error(1)
}

func (m *mockCartRepo) ClearCart(ctx context.Context, userID int) error {
	return m.Called(ctx, userID).Error(0)
}

type mockAddressRepo struct{ mock.Mock }

func (m *mockAddressRepo) GetAddress(ctx context.Context, userID, addrID int) (domain.Address, error) {
	args := m.Called(ctx, userID, addrID)
	return args.Get(0).(domain.Address), args.Error(1)
}

func TestCreateOrder_EmptyCart(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	cartRepo.On("GetCart", mock.Anything, 1).Return([]domain.CartItem{}, nil)

	_, _, _, err := svc.CreateOrder(context.Background(), 1, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestCreateOrder_Success(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	items := []domain.CartItem{
		{ID: 1, Version: 1, UserID: 1, BookID: 10, Quantity: 2},
	}

	now := time.Now()
	expectedOrder := domain.NewOrder(1, 1, 1, "paid", 100.0, nil, nil, now, now)
	expectedItems := []domain.OrderItem{
		domain.NewOrderItem(1, 1, 1, 10, 2, 50.0, "physical", "Test Book", nil),
	}

	cartRepo.On("GetCart", mock.Anything, 1).Return(items, nil)
	orderRepo.On("CreateOrder", mock.Anything, 1, (*int)(nil), items, 0.0).Return(orders_postgres_repository.CreateOrderResult{
		Order: expectedOrder,
		Items: expectedItems,
	}, nil)

	order, orderItems, _, err := svc.CreateOrder(context.Background(), 1, nil)

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder.ID, order.ID)
	assert.Equal(t, expectedOrder.Status, order.Status)
	assert.Len(t, orderItems, 1)
	assert.Equal(t, expectedItems[0].BookID, orderItems[0].BookID)
}

func TestCreateOrder_PhysicalRequiresAddress(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	items := []domain.CartItem{
		{ID: 1, Version: 1, UserID: 1, BookID: 10, Quantity: 2, BookType: "physical"},
	}

	cartRepo.On("GetCart", mock.Anything, 1).Return(items, nil)

	_, _, _, err := svc.CreateOrder(context.Background(), 1, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestCreateOrder_PhysicalWithAddress(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	items := []domain.CartItem{
		{ID: 1, Version: 1, UserID: 1, BookID: 10, Quantity: 2, BookType: "physical"},
	}

	now := time.Now()
	expectedOrder := domain.NewOrder(1, 1, 1, "paid", 350.0, intPtr(5), nil, now, now)
	expectedItems := []domain.OrderItem{
		domain.NewOrderItem(1, 1, 1, 10, 2, 50.0, "physical", "Test Book", nil),
	}

	cartRepo.On("GetCart", mock.Anything, 1).Return(items, nil)
	addrRepo.On("GetAddress", mock.Anything, 1, 5).Return(domain.Address{ID: 5}, nil)
	orderRepo.On("CreateOrder", mock.Anything, 1, intPtr(5), items, 250.0).Return(orders_postgres_repository.CreateOrderResult{
		Order: expectedOrder,
		Items: expectedItems,
	}, nil)

	order, orderItems, _, err := svc.CreateOrder(context.Background(), 1, intPtr(5))

	assert.NoError(t, err)
	assert.Equal(t, expectedOrder.ID, order.ID)
	assert.Equal(t, 350.0, order.TotalAmount)
	assert.Len(t, orderItems, 1)
}

func intPtr(i int) *int {
	return &i
}

func TestPatchOrderStatus_InvalidStatus(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	err := svc.PatchOrderStatus(context.Background(), 1, 1, "invalid-status")

	assert.Error(t, err)
	assert.True(t, errors.Is(err, core_errors.ErrInvalidArgument))
}

func TestPatchOrderStatus_Success(t *testing.T) {
	orderRepo := &mockOrderRepo{}
	cartRepo := &mockCartRepo{}
	addrRepo := &mockAddressRepo{}
	svc := NewOrderService(orderRepo, cartRepo, addrRepo)

	orderRepo.On("UpdateOrderStatus", mock.Anything, 1, 1, "shipped").Return(nil)

	err := svc.PatchOrderStatus(context.Background(), 1, 1, "shipped")

	assert.NoError(t, err)
}
