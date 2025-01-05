package kafka

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port"
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(instance *kafka.Producer) port.OrderProducer {
	return &kafkaProducer{producer: instance}
}

func (p *kafkaProducer) VerifyConsumer(o *order.Order) error {
	var topicName = "order-service.customer.verify.dev.v1"
	replyChannel := "order-service.order.reply.dev.v1"
	value, err := json.Marshal(o.CustomerID)
	if err != nil {
		return err
	}
	header := kafka.Header{
		Key:   "reply-channel",
		Value: []byte(replyChannel),
	}
	mess := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topicName,
		},
		Value:   value,
		Headers: []kafka.Header{header},
	}
	err = p.producer.Produce(&mess, nil)
	return err
}

func (p *kafkaProducer) CreateTicket(o *order.Order) error {
	var topicName = "order-service.kitchen.create.dev.v1"
	replyChannel := "order-service.order.reply.dev.v1"
	value, err := json.Marshal(o.OrderItems)
	if err != nil {
		return err
	}
	header := kafka.Header{
		Key:   "reply-channel",
		Value: []byte(replyChannel),
	}
	mess := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topicName,
		},
		Value:   value,
		Headers: []kafka.Header{header},
	}
	err = p.producer.Produce(&mess, nil)
	return err
}

func (p *kafkaProducer) CompensateTicket(o *order.Order) error {
	var topicName = "order-service.kitchen.compensate.dev.v1"
	replyChannel := "order-service.order.reply.dev.v1"
	value, err := json.Marshal(o.OrderItems)
	if err != nil {
		return err
	}
	header := kafka.Header{
		Key:   "reply-channel",
		Value: []byte(replyChannel),
	}
	mess := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topicName,
		},
		Value:   value,
		Headers: []kafka.Header{header},
	}
	err = p.producer.Produce(&mess, nil)
	return err
}

func (p *kafkaProducer) AuthenticateCard(o *order.Order) error {
	var topicName = "order-service.payment.authenticate.dev.v1"
	replyChannel := "order-service.order.reply.dev.v1"
	value, err := json.Marshal(o.PaymentInfo)
	if err != nil {
		return err
	}
	header := kafka.Header{
		Key:   "reply-channel",
		Value: []byte(replyChannel),
	}
	mess := kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topicName,
		},
		Value:   value,
		Headers: []kafka.Header{header},
	}
	err = p.producer.Produce(&mess, nil)
	return err
}

func (p *kafkaProducer) ApproveTicketCreation(o *order.Order) error {
	return nil
}
func (p *kafkaProducer) ApproveOrderCreation(o *order.Order) error {
	return nil
}
