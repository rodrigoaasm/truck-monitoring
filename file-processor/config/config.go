package config

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

type InfluxDBConfig struct {
	Url, Token, Org, Bucket string
}

type AppConfig struct {
	Kafka  KafkaConfig
	Minio  MinIOConfig
	Influx InfluxDBConfig
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
		Influx: InfluxDBConfig{
			Url:    os.Getenv("INFLUX_URL"),
			Token:  os.Getenv("INFLUX_TOKEN"),
			Org:    os.Getenv("INFLUX_ORG"),
			Bucket: os.Getenv("INFLUX_BUCKET"),
		},
	}

	return appConfig
}
