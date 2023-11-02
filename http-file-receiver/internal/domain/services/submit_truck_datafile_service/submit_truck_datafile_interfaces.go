package submittruckdatafileservice

import (
	"io"

	domainerror "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/error"
)

type DatafileUploadEvent struct {
	Filename string
}

type ISubmitTruckDatafileService interface {
	Handle(io.Reader, DatafileUploadEvent) *domainerror.DomainError
}
