package app

import (
	"delivery-food/order/internal/core/adapter"
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type orderHandler struct {
	s port.OrderService
}

func NewOrderHandler(s port.OrderService) adapter.OrderHandler {
	return &orderHandler{s: s}
}

func (h *orderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order order.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		return
	}
	err = h.s.CreateOrder(&order)
	if err != nil {
		return
	}
}

func (h *orderHandler) FindOrderID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	order, err := h.s.FindOrderByID(uuid.MustParse(id))
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(order)
}
