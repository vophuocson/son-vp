package port

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port/dto"

	"github.com/google/uuid"
)

type OrderService interface {
	CreateOrder(*order.Order) error
	FindOrderByID(uuid.UUID) (*order.Order, error)
}

type OrderConsumer interface {
	ConfirmOrderCreation(confirmOb *dto.ConfirmCreateOrder) error
}
