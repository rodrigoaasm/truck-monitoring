package configs

import (
	"os"
	"strings"
)

type KafkaConfig struct {
	Topic      string
	BrokerList []string
}

type MinIOConfig struct {
	Endpoint, AccessKeyID, SecretAccessKey string
}

type AppConfig struct {
	Kafka KafkaConfig
	Minio MinIOConfig
}

func GetConfig() AppConfig {
	appConfig := AppConfig{
		Kafka: KafkaConfig{
			Topic:      os.Getenv("KAFKA_TOPIC"),
			BrokerList: strings.Split(os.Getenv("KAFKA_BROKER_LIST"), ","),
		},
		Minio: MinIOConfig{
			Endpoint:        os.Getenv("MINIO_ENDPOINT"),
			AccessKeyID:     os.Getenv("MINIO_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("MINIO_SECRET_KEY"),
		},
	}

	return appConfig
}
