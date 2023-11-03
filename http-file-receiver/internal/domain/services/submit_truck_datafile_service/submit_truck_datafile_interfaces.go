package submittruckdatafileservice

import (
	"io"

	domainerror "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/error"
)

type ISubmitTruckDatafileService interface {
	Handle(file io.Reader, filename string, size int64) *domainerror.DomainError
}
