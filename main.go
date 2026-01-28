package main

import (
	"context"
	"golang-arduino-server/internal/core/logger"
	"golang-arduino-server/internal/core/sql"
	"golang-arduino-server/internal/features/readings/repository"
	"golang-arduino-server/internal/features/readings/service"
	"golang-arduino-server/internal/features/readings/transport"
	"golang-arduino-server/internal/server"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn
var rep repository.ReadingsRepository
var ser service.ReadingsService
var hand *transport.HTTPHandlers
var serv *server.HTTPServer

func main() {
	sqlContext, sqlContextCancel := context.WithCancel(context.Background())
	defer sqlContextCancel()

	logger, fileClose, err := logger.NewLogger("INFO")
	if err != nil {
		panic(err)
	}
	defer fileClose()

	conn, err = sql.Connect(sqlContext)
	if err != nil {
		logger.Error("Ошибка при подключении к бд: " + err.Error())
		panic(err)
	}

	rep = repository.NewReadingsRepositoryImpl(sqlContext, logger, conn)
	ser = service.NewReadingsServiceImpl(rep, logger)
	hand = transport.NewHTTPHandlers(ser, logger)
	serv = server.NewHTTPServer(hand)

	if err = serv.StartServer(); err != nil {
		logger.Error("Ошибка запуска сервера: " + err.Error())
	}

}
