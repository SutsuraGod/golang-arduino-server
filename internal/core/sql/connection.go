package sql

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connect(ctx context.Context) (*pgx.Conn, error) {
	conn_string := os.Getenv("CONN_STRING")
	return pgx.Connect(ctx, conn_string)
}
