include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	@docker compose up -d postgres && docker compose up -d grafana

env-down:
	@docker compose down postgres && docker compose down grafana

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migration-up:
	@migrate -path migrations -database ${CONN_STRING} up

migration-down:
	@migrate -path migrations -database ${CONN_STRING} down

server-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/arduino-server/main.go