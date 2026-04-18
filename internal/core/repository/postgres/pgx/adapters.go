package core_pgx_pool

import (
	"errors"
	"fmt"
	core_postgres_pool "golang-arduino-server/internal/core/repository/postgres"

	"github.com/jackc/pgx/v5"
)

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

func mapErrors(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}

	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}
