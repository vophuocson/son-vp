package kafka

import (
	"errors"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	instanceProducer *kafka.Producer
	onceProducer     sync.Once
	instanceConsumer *kafka.Consumer
	onceConsumer     sync.Once
)

func NewKafkaProducerInstance(brokerList string) (*kafka.Producer, error) {
	onceProducer.Do(func() {
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": brokerList,
		})
		if err != nil {
			return
		}
		instanceProducer = producer
	})
	if instanceProducer == nil {
		return nil, errors.New("kafka producer instanceProducer is not initialized")
	}
	return instanceProducer, nil
}

func NewKafkaConsumerInstance(brokerList string) (*kafka.Consumer, error) {
	onceConsumer.Do(func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			// "bootstrap.servers": brokerList,
		})
		if err != nil {
			return
		}
		instanceConsumer = consumer
	})
	if instanceConsumer == nil {
		return nil, errors.New("kafka consumer instanceConsumer is not initialized")
	}
	return instanceConsumer, nil
}
