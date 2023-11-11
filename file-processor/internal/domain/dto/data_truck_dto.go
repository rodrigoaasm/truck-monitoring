package dto

type TruckPointDataDTO struct {
	Timestamp                     uint64  `json:"timestamp"`
	Cockpit_temperature           float32 `json:"cockpit_temperature"`
	Lon                           float32 `json:"lon"`
	Log                           float64 `json:"log"`
	Motion                        bool    `json:"motion"`
	Speed                         float32 `json:"speed"`
	Travelled_distance_last_hours int32   `json:"travelled_distance_last_hours"`
	Kilometerage                  float64 `json:"kilometerage"`
	Engine_rotation               int32   `json:"engine_rotation"`
	March                         string  `json:"march"`
	Gasoline                      int32   `json:"gasoline"`
	Slight_engine_failure         bool    `json:"slight_engine_failure"`
	Serious_engine_failure        bool    `json:"serious_engine_failure"`
	Oil_temperature               float32 `json:"oil_temperature"`
	Oil_pressure                  float64 `json:"oil_pressure"`
	Water_in_fuel                 bool    `json:"water_in_fuel"`
	Water_temperature             float32 `json:"water_temperature"`
	Water_tank                    int32   `json:"water_tank"`
	Air_pressure                  float32 `json:"air_pressure"`
	Cooler_failed                 bool    `json:"cooler_failed"`
	Battery_voltage               float64 `json:"battery_voltage"`
	Battery_current               float64 `json:"battery_current"`
	Battery_charge                int32   `json:"battery_charge"`
	Alternator_voltage            float32 `json:"alternator_voltage"`
	Alternator_current            float32 `json:"alternator_current"`
}

type TruckPointDTO struct {
	Id      string `json:"id"`
	Plate   string `json:"plate"`
	Chassis string `json:"chassis"`
	Data    []TruckPointDataDTO
}
