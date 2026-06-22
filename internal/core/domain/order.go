package domain

import "time"

type Order struct {
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

type OrderItem struct {
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

func NewOrderItem(id, version, orderID, bookID, quantity int, unitPrice float64, itemType, title string, fileKey *string) OrderItem {
	return OrderItem{
		ID:        id,
		Version:   version,
		OrderID:   orderID,
		BookID:    bookID,
		Quantity:  quantity,
		UnitPrice: unitPrice,
		ItemType:  itemType,
		Title:     title,
		FileKey:   fileKey,
	}
}

func NewOrder(id, version, userID int, status string, totalAmount float64, shippingAddressID *int, paymentMethod *string, createdAt, updatedAt time.Time) Order {
	return Order{
		ID:                id,
		Version:           version,
		UserID:            userID,
		Status:            status,
		TotalAmount:       totalAmount,
		ShippingAddressID: shippingAddressID,
		PaymentMethod:     paymentMethod,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}
}

