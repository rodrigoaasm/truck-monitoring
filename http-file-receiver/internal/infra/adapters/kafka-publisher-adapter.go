package adapters

import (
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/interfaces/pubsub"
)

type KafkaPublisherAdapter struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewKafkaPublisherAdapter(brokerList []string, topic string) *KafkaPublisherAdapter {
	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		log.Fatal(69, "Failed to open Kafka producer: %s", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Println("Failed to close Kafka producer cleanly:", err)
		}
	}()

	return &KafkaPublisherAdapter{
		Producer: producer,
		Topic:    topic,
	}
}

func (publisher *KafkaPublisherAdapter) SendEvent(payload pubsub.EventPublisherPayload) error {
	event, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	message := &sarama.ProducerMessage{Topic: publisher.Topic, Value: sarama.StringEncoder(event)}
	if _, _, publishError := publisher.Producer.SendMessage(message); publishError != nil {
		return publishError
	}

	return nil
}
