package repositories

import "github.com/rodrigoaasm/truck-monitoring/file-processor/internal/domain/entities"

type TruckPointRepository interface {
	WriteData([]entities.TruckPoint) error
}
