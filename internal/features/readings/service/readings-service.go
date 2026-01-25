package service

import "golang-arduino-server/internal/features/readings/dto"

type ReadingsService interface {
	SaveReadings(dto *dto.ReadingsDTO) error
}
