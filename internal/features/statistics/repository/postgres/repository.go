package statistics_postgres_repository

import core_postgres_pool "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"

type StatisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(
	pool core_postgres_pool.Pool,
) *StatisticsRepository {
	return &StatisticsRepository{
		pool: pool,
	}
}
