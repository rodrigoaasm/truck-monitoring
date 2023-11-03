package submittruckdatafileservice

import (
	"io"

	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/entities"
	domainerror "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/error"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/interfaces/pubsub"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/interfaces/repositories"
)

type SubmitTruckDatafileService struct {
	DatafileRepository repositories.DataFileRepositoryInterface
	EventPublisher     pubsub.EventPublisherInterface
}

func NewSubmitTruckDatafileService(
	datafileRepository repositories.DataFileRepositoryInterface,
	eventPublisher pubsub.EventPublisherInterface,
) *SubmitTruckDatafileService {

	return &SubmitTruckDatafileService{
		DatafileRepository: datafileRepository,
		EventPublisher:     eventPublisher,
	}
}

func (service *SubmitTruckDatafileService) Handle(file io.Reader, filename string, size int64) *domainerror.DomainError {
	datafileEvent := entities.DatafileUploadEvent{
		Filename: filename,
		Size:     size,
	}

	err := service.DatafileRepository.PutDatafile(file, datafileEvent)
	if err != nil {
		return domainerror.New("Unable to upload datafile")
	}

	err = service.EventPublisher.SendEvent(datafileEvent)
	if err != nil {
		return domainerror.New("Unable to make a datafile upload event")
	}

	return nil
}
