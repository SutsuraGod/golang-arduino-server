package repository

import (
	"golang-arduino-server/internal/features/readings/model"
)

type ReadingsRepository interface {
	Save(mdl *model.Readings) error
}
