package readings_postgres_repository

import (
	"context"
	"fmt"
	"golang-arduino-server/internal/core/domain"
)

func (r *ReadingsRepository) SaveReadings(
	ctx context.Context,
	readings domain.Readings,
) (domain.Readings, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	INSERT INTO generator_readings (gasoline, generator_voltage, network_voltage)
	VALUES ($1, $2, $3)
	RETURNING id, gasoline, generator_voltage, network_voltage, time;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		readings.Gasoline,
		readings.GeneratorVoltage,
		readings.NetworkVoltage,
	)

	var model ReadingsModel
	err := row.Scan(
		&model.ID,
		&model.Gasoline,
		&model.GeneratorVoltage,
		&model.NetworkVoltage,
		&model.Time,
	)
	if err != nil {
		return domain.Readings{}, fmt.Errorf(
			"scan readings: %w",
			err,
		)
	}

	readingsDomain := domain.NewReadings(
		model.ID,
		model.Gasoline,
		model.GeneratorVoltage,
		model.NetworkVoltage,
		model.Time,
	)

	return readingsDomain, nil
}
