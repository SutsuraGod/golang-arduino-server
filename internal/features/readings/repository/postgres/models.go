package readings_postgres_repository

import "time"

type ReadingsModel struct {
	ID               int
	Gasoline         int
	GeneratorVoltage int
	NetworkVoltage   int
	Time             time.Time
}
