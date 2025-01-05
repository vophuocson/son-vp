package order

import "github.com/google/uuid"

type OrderItem struct {
	MenuItemID uuid.UUID `json:"menu_id"`
	Name       string    `json:"name"`
	Quality    int       `json:"quality"`
	Price      float64   `json:"price"`
}
