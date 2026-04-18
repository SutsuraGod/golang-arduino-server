package main

import (
	"context"
	"fmt"
	core_logger "golang-arduino-server/internal/core/logger"
	core_pgx_pool "golang-arduino-server/internal/core/repository/postgres/pgx"
	core_http_middleware "golang-arduino-server/internal/core/transport/http/middleware"
	core_http_server "golang-arduino-server/internal/core/transport/http/server"
	readings_postgres_repository "golang-arduino-server/internal/features/readings/repository/postgres"
	readings_service "golang-arduino-server/internal/features/readings/service"
	readings_http_transport "golang-arduino-server/internal/features/readings/transport/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger: %w", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initializing postgres connection pool")
	pool, err := core_pgx_pool.NewPool(
		ctx,
		core_pgx_pool.NewConfigMust(),
	)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool: %w", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initializing feature", zap.String("feature", "readings"))
	readingsRepository := readings_postgres_repository.NewReadingsRepository(pool)
	readingsService := readings_service.NewReadingsService(readingsRepository)
	readingsTransportHTTP := readings_http_transport.NewReadingsHTTPHandler(readingsService)

	logger.Debug("initializing HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)

	apiVersionRouter1 := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter1.RegisterRoutes(readingsTransportHTTP.Routes()...)

	httpServer.RegisterApiRoutes(apiVersionRouter1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("failed to run http server: %w", zap.Error(err))
	}
}
