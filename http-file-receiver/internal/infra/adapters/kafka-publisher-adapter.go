package adapters

import (
	"context"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/interfaces/pubsub"
	"github.com/segmentio/kafka-go"
)

type KafkaPublisherAdapter struct {
	ProducerWriter kafka.Writer
	Context        context.Context
}

func NewKafkaPublisherAdapter(brokerList []string, topic string) (*KafkaPublisherAdapter, error) {
	kafkaClient, err := kafka.Dial("tcp", brokerList[0])
	if err != nil {
		return nil, err
	}
	kafkaTopic := kafka.TopicConfig{Topic: topic, NumPartitions: 10, ReplicationFactor: 1}
	err = kafkaClient.CreateTopics(kafkaTopic)
	if err != nil {
		return nil, err
	}

	producerWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokerList,
		Topic:   topic,
	})

	return &KafkaPublisherAdapter{
		ProducerWriter: *producerWriter,
		Context:        context.Background(),
	}, nil
}

func (publisher KafkaPublisherAdapter) SendEvent(payload pubsub.EventPublisherPayload) error {
	payloadRaw, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	if err = publisher.ProducerWriter.WriteMessages(
		publisher.Context,
		kafka.Message{
			Key:   []byte(uuid.New().String()),
			Value: payloadRaw,
		},
	); err != nil {
		return err
	}

	return nil
}
