package readings_service

import (
	"context"
	"golang-arduino-server/internal/core/domain"
)

type ReadingsService struct {
	readingsRepository ReadingsRepository
}

type ReadingsRepository interface {
	SaveReadings(
		ctx context.Context,
		readings domain.Readings,
	) (domain.Readings, error)
}

func NewReadingsService(repository ReadingsRepository) *ReadingsService {
	return &ReadingsService{
		readingsRepository: repository,
	}
}
