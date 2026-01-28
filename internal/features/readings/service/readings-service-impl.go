package service

import (
	"golang-arduino-server/internal/features/readings/dto"
	"golang-arduino-server/internal/features/readings/model"
	"golang-arduino-server/internal/features/readings/repository"

	"go.uber.org/zap"
)

type ReadingsServiceImpl struct {
	rep    repository.ReadingsRepository
	logger *zap.Logger
}

func NewReadingsServiceImpl(rep repository.ReadingsRepository, logger *zap.Logger) *ReadingsServiceImpl {
	return &ReadingsServiceImpl{
		rep:    rep,
		logger: logger,
	}
}

func (s *ReadingsServiceImpl) SaveReadings(dto *dto.ReadingsDTO) error {
	if err := dto.ValidateForCreate(); err != nil {
		s.logger.Error("Проблема с присланной dto: " + err.Error())
		return err
	}

	mdl := model.NewReadings(dto.Gasoline, dto.GeneratorVoltage, dto.NetworkVoltage)
	s.logger.Info("Успешный перевод dto в model")

	return s.rep.Save(mdl)
}
