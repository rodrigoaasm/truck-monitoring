package controllers

import (
	"fmt"
	"net/http"

	service "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/services/submit_truck_datafile_service"
)

type SubmitTruckDatafileController struct {
	SubmitTruckDatafileService service.ISubmitTruckDatafileService
}

func (controller SubmitTruckDatafileController) Handle(resWriter http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(0 << 20)

	file, handler, err := req.FormFile("truck-datafile")
	if err != nil {
		fmt.Println("Unable to find datafile")
		resWriter.WriteHeader(400)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	domainerror := controller.SubmitTruckDatafileService.Handle(file, handler.Filename, handler.Size)
	if domainerror != nil {
		fmt.Println(domainerror.Message)
		resWriter.WriteHeader(500)
		return
	}

	fmt.Fprintf(resWriter, "Successfully Uploaded File\n")
}
