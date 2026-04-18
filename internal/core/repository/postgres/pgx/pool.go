package core_pgx_pool

import (
	"context"
	"fmt"
	core_postgres_pool "golang-arduino-server/internal/core/repository/postgres"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewPool(
	ctx context.Context,
	config Config,
) (*Pool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("pgx parse config: %w", err)
	}

	pgxpool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool with config: %w", err)
	}

	if err := pgxpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping pgx pool: %w", err)
	}

	return &Pool{
		pgxpool,
		config.OpTimeout,
	}, nil
}

func (p *Pool) QueryRow(
	ctx context.Context,
	sql string,
	args ...any,
) core_postgres_pool.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)

	return pgxRow{row}
}

func (p *Pool) OpTimeout() time.Duration {
	return p.opTimeout
}
