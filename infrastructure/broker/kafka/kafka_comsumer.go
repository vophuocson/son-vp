package kafka

import (
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/port/dto"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/avrov2"
)

type kafkaConsumer struct {
	client *kafka.Consumer
}

func NewKafkaConsumer(client *kafka.Consumer) port.OrderConsumer {
	return &kafkaConsumer{client: client}
}

func (c *kafkaConsumer) ConfirmOrderCreation(confirmOb *dto.ConfirmCreateOrder) error {
	channel := "order-service.order.confirm-creation.dev.v1"
	c.client.Subscribe(channel, nil)
	var url = ""
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(url))
	if err != nil {
		return err
	}

	deser, err := avrov2.NewDeserializer(client, serde.ValueSerde, avrov2.NewDeserializerConfig())
	if err != nil {
		return err
	}

	for {
		isSuccess := true
		for _, serviceSuccess := range confirmOb.ChannelNamesReply {
			if !serviceSuccess {
				isSuccess = false
				break
			}
		}
		if isSuccess {
			return nil
		}
		msg, err := c.client.ReadMessage(10 * time.Second)
		if err != nil {
			c.client.Unsubscribe()
			return err
		}
		if *(msg.TopicPartition.Topic) == channel {
			var replyData dto.ReplyOrderCreation
			err := deser.DeserializeInto(*msg.TopicPartition.Topic, msg.Value, &replyData)
			if err != nil {
				return err
			} else if replyData.OrderID == confirmOb.OrderID {
				switch {
				case !confirmOb.ChannelNamesReply[replyData.ServiceNameReply]:
					confirmOb.ChannelNamesReply[replyData.ServiceNameReply] = true
				default:

				}
			}
		}
	}
}
