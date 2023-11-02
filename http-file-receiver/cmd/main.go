package main

import (
	"net/http"

	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/cmd/api"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	api.CreateApp(router)
	http.ListenAndServe(":7001", router)
}
