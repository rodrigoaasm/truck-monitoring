package api

import (
	"github.com/gorilla/mux"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/cmd/api/controllers"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/configs"
	submittruckdatafileservice "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/services/submit_truck_datafile_service"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/infra/adapters"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/infra/repositories"
)

func CreateApp(apiRouter *mux.Router) {

	appConfig := configs.GetConfig()

	// dependencies
	kafkaPublisherAdapter := adapters.NewKafkaPublisherAdapter(
		appConfig.Kafka.BrokerList,
		appConfig.Kafka.Topic,
	)
	minioRepository := repositories.NewMinIORepository(
		appConfig.Minio.Endpoint,
		appConfig.Minio.AccessKeyID,
		appConfig.Minio.SecretAccessKey,
	)

	//services
	submitTruckDatafileService := submittruckdatafileservice.NewSubmitTruckDatafileService(minioRepository, kafkaPublisherAdapter)

	// controllers
	submitTruckDatafileController := controllers.SubmitTruckDatafileController{
		SubmitTruckDatafileService: submitTruckDatafileService,
	}

	apiRouter.HandleFunc("/datafile/upload", submitTruckDatafileController.Handle).Methods("POST")
}
