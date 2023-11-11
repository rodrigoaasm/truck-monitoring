package entities

type TruckPoint struct {
	Id                            string
	Plate                         string
	Chassis                       string
	Timestamp                     uint64
	Cockpit_temperature           float32
	Lon                           float32
	Log                           float64
	Motion                        bool
	Speed                         float32
	Travelled_distance_last_hours int32
	Kilometerage                  float64
	Engine_rotation               int32
	March                         string
	Gasoline                      int32
	Slight_engine_failure         bool
	Serious_engine_failure        bool
	Oil_temperature               float32
	Oil_pressure                  float64
	Water_in_fuel                 bool
	Water_temperature             float32
	Water_tank                    int32
	Air_pressure                  float32
	Cooler_failed                 bool
	Battery_voltage               float64
	Battery_current               float64
	Battery_charge                int32
	Alternator_voltage            float32
	Alternator_current            float32
}
