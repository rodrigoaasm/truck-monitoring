package pubsub

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	datafileprocess "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/services/datafile-process"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	kafkaReader            *kafka.Reader
	datafileProcessService datafileprocess.DatafileProcessServiceInterface
}

func NewKafkaConsumer(url string, topic string, datafileProcessService datafileprocess.DatafileProcessServiceInterface) *KafkaConsumer {
	return &KafkaConsumer{
		kafkaReader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:  []string{url},
			GroupID:  "file-processor",
			Topic:    topic,
			MaxBytes: 10e6, // 10MB
		}),
	}
}

func (consumer *KafkaConsumer) Run() error {
	for {
		m, err := consumer.kafkaReader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		datafileUploadEvent := datafileprocess.DatafileUploadEvent{}
		if err = json.Unmarshal(m.Value, &datafileUploadEvent); err != nil {
			textErr := fmt.Sprintf("Unable to unmarshal data stream. %v", err)
			return errors.New(textErr)
		}
		consumer.datafileProcessService.Process(datafileUploadEvent)
	}
}
