package order

import (
	order "delivery-food/order/internal/core/domain/valueobject"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID                 uuid.UUID           `json:"id"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	State              string              `json:"state"`
	TotalPrice         float64             `json:"total_price"`
	SpecialInstruction string              `json:"special_instruction"`
	Discount           float64             `json:"discount"`
	CustomerID         uuid.UUID           `json:"customer_id"`
	RestaurantID       uuid.UUID           `json:"restaurant_id"`
	PaymentInfo        *order.PaymentInfo  `json:"payment_info"`
	OrderItems         []*order.OrderItem  `json:"order_items"`
	DeliveryInfo       *order.DeliveryInfo `json:"delivery_info"`
}
