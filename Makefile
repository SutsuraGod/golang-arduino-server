include .env
export

server-run:
	go run main.go

migration-up:
	migrate -path migrations -database ${CONN_STRING} up

migration-down:
	migrate -path migrations -database ${CONN_STRING} down