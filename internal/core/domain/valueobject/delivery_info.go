package order

import "time"

type DeliveryInfo struct {
	Address       string    `json:"address"`
	PlacedTime    time.Time `json:"placed_time"`
	DeliveredTime time.Time `json:"delivery_time,omitempty"`
	Status        string    `json:"status"`
}
