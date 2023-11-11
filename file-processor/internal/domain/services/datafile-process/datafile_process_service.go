package datafileprocess

import (
	"bufio"
	"encoding/json"
	"fmt"

	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/dto"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/entities"
	domainerror "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/error"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/interfaces/repositories"
)

type DatafileProcessService struct {
	DatafileRepository   repositories.DataFileRepositoryInterface
	TruckPointRepository repositories.TruckPointRepository
}

func NewDatafileProcessService(
	datafileRepository repositories.DataFileRepositoryInterface,
	truckPointRepository repositories.TruckPointRepository,
) *DatafileProcessService {
	return &DatafileProcessService{
		DatafileRepository:   datafileRepository,
		TruckPointRepository: truckPointRepository,
	}
}

func (service *DatafileProcessService) Process(datafileUploadEvent DatafileUploadEvent) *domainerror.DomainError {
	file, err := service.DatafileRepository.GetDatafile(datafileUploadEvent.Filename)
	if err != nil {
		return domainerror.New("File Not found")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var dataText string
	for fileScanner.Scan() {
		dataText += fileScanner.Text()
	}

	var truckPointDTO dto.TruckPointDTO
	if err = json.Unmarshal([]byte(dataText), &truckPointDTO); err != nil {
		textErr := fmt.Sprintf("Unable to unmarshal data stream. %v", err)
		return domainerror.New(textErr)
	}

	var truckPoints []entities.TruckPoint
	for _, truckPointData := range truckPointDTO.Data {
		truckPoints = append(truckPoints, entities.TruckPoint{
			Id:                            truckPointDTO.Id,
			Plate:                         truckPointDTO.Plate,
			Chassis:                       truckPointDTO.Chassis,
			Timestamp:                     truckPointData.Timestamp,
			Cockpit_temperature:           truckPointData.Cockpit_temperature,
			Lon:                           truckPointData.Lon,
			Log:                           truckPointData.Log,
			Kilometerage:                  truckPointData.Kilometerage,
			Motion:                        truckPointData.Motion,
			Speed:                         truckPointData.Speed,
			Engine_rotation:               truckPointData.Engine_rotation,
			Travelled_distance_last_hours: truckPointData.Travelled_distance_last_hours,
			March:                         truckPointData.March,
			Gasoline:                      truckPointData.Gasoline,
			Slight_engine_failure:         truckPointData.Slight_engine_failure,
			Serious_engine_failure:        truckPointData.Serious_engine_failure,
			Oil_temperature:               truckPointData.Oil_temperature,
			Oil_pressure:                  truckPointData.Oil_pressure,
			Water_in_fuel:                 truckPointData.Water_in_fuel,
			Water_temperature:             truckPointData.Water_temperature,
			Water_tank:                    truckPointData.Water_tank,
			Cooler_failed:                 truckPointData.Cooler_failed,
			Battery_voltage:               truckPointData.Battery_voltage,
			Battery_current:               truckPointData.Battery_current,
			Battery_charge:                truckPointData.Battery_charge,
			Air_pressure:                  truckPointData.Air_pressure,
			Alternator_voltage:            truckPointData.Alternator_voltage,
			Alternator_current:            truckPointData.Alternator_current,
		})
	}

	if err = service.TruckPointRepository.WriteData(truckPoints); err != nil {
		return domainerror.New("Unable to write points")
	}

	return nil
}
