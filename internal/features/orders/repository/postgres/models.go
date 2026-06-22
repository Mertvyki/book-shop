package orders_postgres_repository

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type OrderModel struct {
	ID                int
	Version           int
	UserID            int
	Status            string
	TotalAmount       float64
	ShippingAddressID *int
	PaymentMethod     *string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

func (m OrderModel) ToDomain() domain.Order {
	return domain.Order{
		ID:                m.ID,
		Version:           m.Version,
		UserID:            m.UserID,
		Status:            m.Status,
		TotalAmount:       m.TotalAmount,
		ShippingAddressID: m.ShippingAddressID,
		PaymentMethod:     m.PaymentMethod,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}

type OrderItemModel struct {
	ID        int
	Version   int
	OrderID   int
	BookID    int
	Quantity  int
	UnitPrice float64
	ItemType  string
	Title     string
	FileKey   *string
}

func (m OrderItemModel) ToDomain() domain.OrderItem {
	return domain.OrderItem{
		ID:        m.ID,
		Version:   m.Version,
		OrderID:   m.OrderID,
		BookID:    m.BookID,
		Quantity:  m.Quantity,
		UnitPrice: m.UnitPrice,
		ItemType:  m.ItemType,
		Title:     m.Title,
		FileKey:   m.FileKey,
	}
}
