package datafileprocess_test

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/entities"
	datafileprocess "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/services/datafile-process"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/tests/mocks"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) (
	*mocks.MockDataFileRepositoryInterface,
	*mocks.MockTruckPointRepository,
	*datafileprocess.DatafileProcessService,
) {
	mockControl := gomock.NewController(t)
	defer mockControl.Finish()

	datafileRepository := mocks.NewMockDataFileRepositoryInterface(mockControl)
	influxdbRepository := mocks.NewMockTruckPointRepository(mockControl)

	datafileProcessService := datafileprocess.NewDatafileProcessService(
		datafileRepository,
		influxdbRepository,
	)

	return datafileRepository, influxdbRepository, datafileProcessService
}

func TestDataFileRepositoryInterface__Process(t *testing.T) {
	datafileRepository, influxdbRepository, datafileProcessService := setup(t)

	datafileUploadEvent := datafileprocess.DatafileUploadEvent{
		Filename: "file.json",
		Size:     1874,
	}

	file, _ := os.Open("../../../../tests/assets/json_truck_example.json")
	truckPoints := []entities.TruckPoint{}
	truckPoints = append(truckPoints, entities.TruckPoint{
		Id:                            "1284411abbh-erttfdfgv-198125",
		Plate:                         "ABC1D12",
		Chassis:                       "4T1BD1EB0FU036684",
		Timestamp:                     1698930000,
		Cockpit_temperature:           29.7,
		Lon:                           -22.0132000,
		Log:                           44.0132000,
		Motion:                        true,
		Speed:                         91,
		Travelled_distance_last_hours: 87,
		Kilometerage:                  195000,
		Engine_rotation:               1900,
		March:                         "4r",
		Gasoline:                      181,
		Slight_engine_failure:         false,
		Serious_engine_failure:        false,
		Oil_temperature:               91.2,
		Oil_pressure:                  1.8,
		Water_in_fuel:                 false,
		Water_temperature:             71,
		Water_tank:                    89,
		Air_pressure:                  1.2,
		Cooler_failed:                 false,
		Battery_voltage:               12.31456,
		Battery_current:               149.12544,
		Battery_charge:                89,
		Alternator_voltage:            23.98741,
		Alternator_current:            89.25711,
	},
	)

	datafileRepository.EXPECT().GetDatafile(datafileUploadEvent.Filename).Return(file, nil)
	influxdbRepository.EXPECT().WriteData(truckPoints).Return(nil)
	domainError := datafileProcessService.Process(datafileUploadEvent)
	require.Nil(t, domainError, "Should get file successful")
}
