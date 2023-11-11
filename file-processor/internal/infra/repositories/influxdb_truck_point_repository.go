package repositories

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/entities"
)

type TruckPointRepository struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
	context  context.Context
}

func NewTruckPointRepository(url string, token string, org string, bucket string) *TruckPointRepository {
	client := influxdb2.NewClient(url, token)
	return &TruckPointRepository{
		client:   client,
		writeAPI: client.WriteAPIBlocking(org, bucket),
		context:  context.Background(),
	}
}

func (repo *TruckPointRepository) WriteData(truckPoints []entities.TruckPoint) error {
	for _, truckPoint := range truckPoints {
		point := *influxdb2.NewPoint(
			"truck-data",
			map[string]string{},
			map[string]interface{}{
				"cockpit_temperature":           truckPoint.Cockpit_temperature,
				"lon":                           truckPoint.Lon,
				"log":                           truckPoint.Log,
				"motion":                        truckPoint.Motion,
				"speed":                         truckPoint.Speed,
				"travelled_distance_last_hours": truckPoint.Travelled_distance_last_hours,
				"kilometerage":                  truckPoint.Kilometerage,
				"engine_rotation":               truckPoint.Engine_rotation,
				"march":                         truckPoint.March,
				"gasoline":                      truckPoint.Gasoline,
				"slight_engine_failure":         truckPoint.Slight_engine_failure,
				"serious_engine_failure":        truckPoint.Serious_engine_failure,
				"oil_temperature":               truckPoint.Oil_temperature,
				"oil_pressure":                  truckPoint.Oil_pressure,
				"water_in_fuel":                 truckPoint.Water_in_fuel,
				"water_temperature":             truckPoint.Water_temperature,
				"water_tank":                    truckPoint.Water_tank,
				"air_pressure":                  truckPoint.Air_pressure,
				"cooler_failed":                 truckPoint.Cooler_failed,
				"battery_voltage":               truckPoint.Battery_voltage,
				"battery_current":               truckPoint.Battery_current,
				"battery_charge":                truckPoint.Battery_charge,
				"alternator_voltage":            truckPoint.Alternator_voltage,
				"Alternator_current":            truckPoint.Alternator_current,
			},
			time.Unix(int64(truckPoint.Timestamp), 0),
		)
		repo.writeAPI.WritePoint(repo.context, &point)
	}

	return nil
}
