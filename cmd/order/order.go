package main

import (
	"database/sql"
	"delivery-food/order/infrastructure/broker/kafka"
	"delivery-food/order/infrastructure/saga/temporal"
	"delivery-food/order/internal/app"
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/service"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var orderRepo port.OrderRepository
	// create the producer outbound adapter
	kafKaProducerClient, err := kafka.NewKafkaProducerInstance("abc")
	if err != nil {
		return
	}
	// create the the producer outbound port
	p := kafka.NewKafkaProducer(kafKaProducerClient)

	// create the consumer inbound adapter
	kafkaConsumerClient, err := kafka.NewKafkaConsumerInstance("")
	if err != nil {
		return
	}
	// create the consumer inbound port
	c := kafka.NewKafkaConsumer(kafkaConsumerClient)

	// create the saga orchestrator outbound adapter
	tClient, err := temporal.NewTemporalClientInstance()
	if err != nil {
		return
	}
	// create the saga orchestrator outbound port
	orchestrator := temporal.NewTemporalClient(tClient)

	// create the business logic for the order service
	orderService := service.NewOrderService(orderRepo, p, c, orchestrator)
	// create the inbound adapter
	orderHandler := app.NewOrderHandler(orderService)
	http.HandleFunc("/orders", orderHandler.CreateOrder)
	http.HandleFunc("/order/{order_id}", orderHandler.FindOrderID)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func OpenConnectionPostgresql() *sql.DB {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	// Open the connection
	db, err := sql.Open("pgx", "")
	if err != nil {
		log.Fatalf("impossible to create the connection: %s", err)
	}
	// close connection(s) when the surrounding function exits
	defer db.Close()
	// define some settings (refer to the driver documentation)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)
	return db
}
