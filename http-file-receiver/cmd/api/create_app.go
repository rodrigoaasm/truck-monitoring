package api

import (
	"github.com/gorilla/mux"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/cmd/api/controllers"
)

func CreateApp(apiRouter *mux.Router) {

	submitTruckDatafileController := controllers.SubmitTruckDatafileController{}

	apiRouter.HandleFunc("/datafile/upload", submitTruckDatafileController.Handle).Methods("POST")
}
