package orders_transport_http

import (
	"time"

	"github.com/Mertvyki/book-shop/internal/core/domain"
)

type OrderDTOResponse struct {
	ID                int                  `json:"id"`
	Version           int                  `json:"version"`
	UserID            int                  `json:"user_id"`
	Status            string               `json:"status"`
	TotalAmount       float64              `json:"total_amount"`
	DeliveryCost      float64              `json:"delivery_cost"`
	ShippingAddressID *int                 `json:"shipping_address_id"`
	PaymentMethod     *string              `json:"payment_method"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	Items             []OrderItemDTOResponse `json:"items,omitempty"`
}

type OrderItemDTOResponse struct {
	ID        int     `json:"id"`
	Version   int     `json:"version"`
	OrderID   int     `json:"order_id"`
	BookID    int     `json:"book_id"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	ItemType  string  `json:"item_type"`
	Title     string  `json:"title"`
	FileKey   *string `json:"file_key"`
}

func orderDTOFromDomain(order domain.Order) OrderDTOResponse {
	return OrderDTOResponse{
		ID:                order.ID,
		Version:           order.Version,
		UserID:            order.UserID,
		Status:            order.Status,
		TotalAmount:       order.TotalAmount,
		ShippingAddressID: order.ShippingAddressID,
		PaymentMethod:     order.PaymentMethod,
		CreatedAt:         order.CreatedAt,
		UpdatedAt:         order.UpdatedAt,
	}
}

func orderWithItemsDTO(order domain.Order, items []domain.OrderItem, deliveryCost float64) OrderDTOResponse {
	dto := orderDTOFromDomain(order)
	dto.DeliveryCost = deliveryCost
	dto.Items = orderItemsDTOFromDomains(items)
	return dto
}

func ordersDTOFromDomains(orders []domain.Order) []OrderDTOResponse {
	dtos := make([]OrderDTOResponse, len(orders))
	for i, order := range orders {
		dtos[i] = orderDTOFromDomain(order)
	}

	return dtos
}

func orderItemDTOFromDomain(item domain.OrderItem) OrderItemDTOResponse {
	return OrderItemDTOResponse{
		ID:        item.ID,
		Version:   item.Version,
		OrderID:   item.OrderID,
		BookID:    item.BookID,
		Quantity:  item.Quantity,
		UnitPrice: item.UnitPrice,
		ItemType:  item.ItemType,
		Title:     item.Title,
		FileKey:   item.FileKey,
	}
}

func orderItemsDTOFromDomains(items []domain.OrderItem) []OrderItemDTOResponse {
	dtos := make([]OrderItemDTOResponse, len(items))
	for i, item := range items {
		dtos[i] = orderItemDTOFromDomain(item)
	}

	return dtos
}
