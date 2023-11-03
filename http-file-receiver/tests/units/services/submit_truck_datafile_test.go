package submittruckdatafileservice_test

import (
	"errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/entities"
	submittruckdatafileservice "github.com/rodrigoaasm/truck-monitoring/http-file-receiver/internal/domain/services/submit_truck_datafile_service"
	"github.com/rodrigoaasm/truck-monitoring/http-file-receiver/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (
	*mocks.MockDataFileRepositoryInterface,
	*mocks.MockEventPublisherInterface,
	submittruckdatafileservice.SubmitTruckDatafileService,
) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()
	datafileRepository := mocks.NewMockDataFileRepositoryInterface(mockControl)
	eventPublisher := mocks.NewMockEventPublisherInterface(mockControl)

	submitDatafileService := submittruckdatafileservice.SubmitTruckDatafileService{
		DatafileRepository: datafileRepository,
		EventPublisher:     eventPublisher,
	}

	return datafileRepository, eventPublisher, submitDatafileService
}

func TestSubmitTruckDatafile_Handle(t *testing.T) {
	datafileRepository, eventPublisher, submitDatafileService := setup(t)

	file, _ := os.Open("../../assets/json_truck_example.json")
	filedataEvent := entities.DatafileUploadEvent{
		Filename: "datafile_example",
		Size:     966,
	}

	datafileRepository.EXPECT().PutDatafile(file, filedataEvent).Return(nil)
	eventPublisher.EXPECT().SendEvent(filedataEvent)
	err := submitDatafileService.Handle(file, filedataEvent.Filename, filedataEvent.Size)
	require.Nil(t, err, "Should upload sucessful")

	datafileRepository.EXPECT().PutDatafile(file, filedataEvent).Return(errors.New("failed upload"))
	err = submitDatafileService.Handle(file, filedataEvent.Filename, filedataEvent.Size)
	assert.Equal(t, err.Message, "Unable to upload datafile", "Should return an upload error")

	datafileRepository.EXPECT().PutDatafile(file, filedataEvent).Return(nil)
	eventPublisher.EXPECT().SendEvent(filedataEvent).Return(errors.New("publish failed"))
	err = submitDatafileService.Handle(file, filedataEvent.Filename, filedataEvent.Size)
	assert.Equal(t, err.Message, "Unable to make a datafile upload event", "Should return a publish error")
}
