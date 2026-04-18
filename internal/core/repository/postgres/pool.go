package core_postgres_pool

import (
	"context"
	"time"
)

type Pool interface {
	QueryRow(ctx context.Context, sql string, args ...any) Row

	OpTimeout() time.Duration
}

type Row interface {
	Scan(dest ...any) error
}
