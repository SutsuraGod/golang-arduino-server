package readings_postgres_repository

import (
	core_postgres_pool "golang-arduino-server/internal/core/repository/postgres"
)

type ReadingsRepository struct {
	pool core_postgres_pool.Pool
}

func NewReadingsRepository(pool core_postgres_pool.Pool) *ReadingsRepository {
	return &ReadingsRepository{
		pool: pool,
	}
}
