package repository

import (
	"context"
	"golang-arduino-server/internal/features/readings/model"
	"sync"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ReadingsRepositoryImpl struct {
	sqlContext context.Context
	logger     *zap.Logger
	mtx        sync.Mutex
	conn       *pgx.Conn
}

func NewReadingsRepositoryImpl(sqlContext context.Context, logger *zap.Logger, conn *pgx.Conn) *ReadingsRepositoryImpl {
	return &ReadingsRepositoryImpl{
		sqlContext: sqlContext,
		logger:     logger,
		conn:       conn,
	}
}

func (r *ReadingsRepositoryImpl) Save(mdl *model.Readings) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	queryString := `
	INSERT INTO generator_readings(gasoline, generator_voltage, network_voltage)
	VALUES($1, $2, $3);		
	`
	_, err := r.conn.Exec(r.sqlContext, queryString, mdl.GetGasoline(), mdl.GetGeneratorVoltage(), mdl.GetNetworkVoltage())
	if err == nil {
		r.logger.Info("Успешное сохранение данных")
	} else {
		r.logger.Error("Ошибка при записи данных в бд: " + err.Error())
	}

	return err
}
