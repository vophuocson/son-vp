package adapter

import "net/http"

type OrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	FindOrderID(w http.ResponseWriter, r *http.Request)
}
