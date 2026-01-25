package dto

import "errors"

type ReadingsDTO struct {
	Gasoline         int
	GeneratorVoltage int
	NetworkVoltage   int
}

func (d *ReadingsDTO) ValidateForCreate() error {
	if d.Gasoline < 0 || d.Gasoline > 100 {
		return errors.New("Некорректный уровень топлива")
	}

	if d.GeneratorVoltage < 0 || d.GeneratorVoltage > 260 {
		return errors.New("Некорректное напряжение генератора")
	}

	if d.NetworkVoltage < 0 || d.NetworkVoltage > 260 {
		return errors.New("Некорректное напряжение сети")
	}

	return nil
}
