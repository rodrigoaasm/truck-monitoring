package repositories

import "io"

type DataFileRepositoryInterface interface {
	PutDatafile(fileReader io.Reader) error
}
