package order

import "github.com/google/uuid"

type PaymentInfo struct {
	Status string    `json:"status"`
	ID     uuid.UUID `json:"id"`
}
