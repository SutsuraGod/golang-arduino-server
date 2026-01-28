package model

import "time"

type Readings struct {
	id               int
	gasoline         int
	generatorVoltage int
	networkVoltage   int
	time             time.Time
}

func NewReadingsWithIdAndTime(id int, gasoline int, generatorVoltage int, networkVoltage int, time time.Time) *Readings {
	return &Readings{
		id:               id,
		gasoline:         gasoline,
		generatorVoltage: generatorVoltage,
		networkVoltage:   networkVoltage,
		time:             time,
	}
}

func NewReadings(gasoline int, generatorVoltage int, networkVoltage int) *Readings {
	return &Readings{
		gasoline:         gasoline,
		generatorVoltage: generatorVoltage,
		networkVoltage:   networkVoltage,
	}
}

func (r *Readings) GetId() int {
	return r.id
}

func (r *Readings) GetGasoline() int {
	return r.gasoline
}

func (r *Readings) GetGeneratorVoltage() int {
	return r.generatorVoltage
}

func (r *Readings) GetNetworkVoltage() int {
	return r.networkVoltage
}
