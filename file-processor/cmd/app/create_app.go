package app

import (
	"github.com/rodrigoaasm/truck-monitoring/file-processor/cmd/app/pubsub"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/config"
	datafileprocess "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/services/datafile-process"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/infra/repositories"
)

func CreateApp() *pubsub.KafkaConsumer {
	appConfig := config.GetConfig()

	// dependencies
	minioRepository := repositories.NewMinIORepository(
		appConfig.Minio.Endpoint,
		appConfig.Minio.AccessKeyID,
		appConfig.Minio.SecretAccessKey,
	)
	influxdbTruckPointRepository := repositories.NewTruckPointRepository(
		appConfig.Influx.Url, appConfig.Influx.Token, appConfig.Influx.Org, appConfig.Influx.Bucket,
	)

	// services
	datafileProcessService := datafileprocess.NewDatafileProcessService(minioRepository, influxdbTruckPointRepository)

	kafkaConsumer := pubsub.NewKafkaConsumer(appConfig.Influx.Url, appConfig.Kafka.Topic, datafileProcessService)

	return kafkaConsumer
}
