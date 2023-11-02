package submittruckdatafileservice

import (
	"io"

	domainerror "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/error"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/interfaces/pubsub"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/interfaces/repositories"
)

type SubmitTruckDatafileService struct {
	DatafileRepository repositories.DataFileRepositoryInterface
	EventPublisher     pubsub.EventPublisherInterface
}

func (controller SubmitTruckDatafileService) Handle(file io.Reader, datafileEvent DatafileUploadEvent) *domainerror.DomainError {
	err := controller.DatafileRepository.PutDatafile(file)
	if err != nil {
		return domainerror.New("Unable to upload datafile")
	}

	err = controller.EventPublisher.SendEvent(datafileEvent)
	if err != nil {
		return domainerror.New("Unable to make a datafile upload event")
	}

	return nil
}
