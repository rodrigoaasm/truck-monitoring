package repositories

import (
	"io"

	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/entities"
)

type DataFileRepositoryInterface interface {
	PutDatafile(fileReader io.Reader, datafileEvent entities.DatafileUploadEvent) error
}
