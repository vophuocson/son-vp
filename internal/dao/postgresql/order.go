package dao

import (
	"database/sql"
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type userPostgresql struct {
	db *sql.DB
}

func NewUserPostgresql(db *sql.DB) port.OrderRepository {
	return &userPostgresql{db: db}
}

func (re *userPostgresql) Create(order *order.Order) error {
	commandOrder, err := re.db.Prepare("INSERT INTO orders (state, total_price, special_instruction, discount, customer_id, restaurant_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when preparing the insert statement for the order")
	}
	commandOrderItem, err := re.db.Prepare("INSERT INTO order_items (order_id, menu_item_id, name, quality, price) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when preparing the insert statement for the order item")
	}
	commandOrderDeliveryInfo, err := re.db.Prepare("INSERT INTO delivery_infos (order_id, address, placed_time, delivered_time, status) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when preparing the insert statement for the delivery information")
	}
	commandOrderPaymentInfo, err := re.db.Prepare("INSERT INTO payment_infos (order_id, status, id) VALUES(?, ?, ?)")
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when preparing the insert statement for the payment information")
	}
	tx, err := re.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when beginning the transaction")
	}
	result, err := tx.Stmt(commandOrder).Exec(order.State, order.TotalPrice, order.SpecialInstruction, order.Discount, order.CustomerID, order.RestaurantID)
	if err != nil {
		return errors.Wrap(err, "Create(order *order.Order): error when inserting the order")
	}
	orderID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Create(order *order.Order): error when retrieving the order ID")
	}
	for idx := range order.OrderItems {
		_, err = tx.Stmt(commandOrderItem).Exec(orderID, order.OrderItems[idx].MenuItemID, order.OrderItems[idx].Name, order.OrderItems[idx].Quality, order.OrderItems[idx].Price)
		if err != nil {
			tx.Rollback()
			return errors.Wrap(err, "Create(order *order.Order): error when inserting the order item")
		}
	}
	_, err = tx.Stmt(commandOrderDeliveryInfo).Exec(orderID, order.DeliveryInfo.Address, order.DeliveryInfo.PlacedTime, order.DeliveryInfo.DeliveredTime, order.DeliveryInfo.Status)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Create(order *order.Order): error when inserting the delivery information")
	}
	_, err = tx.Stmt(commandOrderPaymentInfo).Exec(orderID, order.PaymentInfo.Status, order.PaymentInfo.Status)
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Create(order *order.Order): error when inserting the payment information")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.Wrap(err, "Create(order *order.Order): Error when committing the transaction that creates the aggregate order")
	}
	return nil
}
func (re *userPostgresql) GetByID(id uuid.UUID) (*order.Order, error) {
	var order order.Order
	err := re.db.QueryRow("select * from orders where id = ?", id).Scan(&order)
	if err != nil {
		return nil, errors.Wrap(err, "GetByID(id uuid.UUID): Error when retrieving the order")
	}

	err = re.db.QueryRow("select * from order_items where order_id = ?", id).Scan(&order.OrderItems)
	if err != nil {
		return nil, errors.Wrap(err, "GetByID(id uuid.UUID): Error when retrieving order items")
	}

	err = re.db.QueryRow("select * from delivery_infos where order_id = ?", id).Scan(&order.DeliveryInfo)
	if err != nil {
		return nil, errors.Wrap(err, "GetByID(id uuid.UUID): Error when retrieving the order delivery information")
	}

	err = re.db.QueryRow("select * from payment_infos where order_id = ?", id).Scan(&order.PaymentInfo)
	if err != nil {
		return nil, errors.Wrap(err, "GetByID(id uuid.UUID): Error when retrieving the order payment information")
	}
	return nil, nil
}
