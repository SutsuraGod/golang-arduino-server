package readings_service

import (
	"context"
	"fmt"
	"golang-arduino-server/internal/core/domain"
)

func (s *ReadingsService) SaveReadings(
	ctx context.Context,
	readings domain.Readings,
) (domain.Readings, error) {
	if err := readings.Validate(); err != nil {
		return domain.Readings{}, fmt.Errorf(
			"failed to validate readings domain: %w",
			err,
		)
	}

	readings, err := s.readingsRepository.SaveReadings(ctx, readings)
	if err != nil {
		return domain.Readings{}, fmt.Errorf(
			"failed to save readings by repository: %w",
			err,
		)
	}

	return readings, nil
}
