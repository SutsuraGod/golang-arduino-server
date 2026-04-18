package domain

import (
	"fmt"
	core_errors "golang-arduino-server/internal/core/errors"
	"time"
)

type Readings struct {
	ID               int
	Gasoline         int
	GeneratorVoltage int
	NetworkVoltage   int
	Time             time.Time
}

func NewReadings(
	id int,
	gasoline int,
	generatorVoltage int,
	networkVoltage int,
	time time.Time,
) Readings {
	return Readings{
		ID:               id,
		Gasoline:         gasoline,
		GeneratorVoltage: generatorVoltage,
		NetworkVoltage:   networkVoltage,
		Time:             time,
	}
}

func NewUnitializedReadings(
	gasoline int,
	generatorVoltage int,
	networkVoltage int,
) Readings {
	return NewReadings(
		UnitializedID,
		gasoline,
		generatorVoltage,
		networkVoltage,
		time.Now(),
	)
}

func (r *Readings) Validate() error {
	if r.Gasoline < 0 || r.Gasoline > 100 {
		return fmt.Errorf(
			"ivalid `Gasoline` value: %d: %w",
			r.Gasoline,
			core_errors.ErrInvalidArgument,
		)
	}

	if r.GeneratorVoltage < 0 || r.GeneratorVoltage > 250 {
		return fmt.Errorf(
			"invalid `GeneratorVoltage` value: %d: %w",
			r.GeneratorVoltage,
			core_errors.ErrInvalidArgument,
		)
	}

	if r.NetworkVoltage < 0 || r.NetworkVoltage > 250 {
		return fmt.Errorf(
			"invalid `NetworkVoltage` value: %d: %w",
			r.NetworkVoltage,
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}
