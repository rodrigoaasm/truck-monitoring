package repositories

import (
	"io"
)

type DataFileRepositoryInterface interface {
	GetDatafile(filename string) (io.Reader, error)
}
