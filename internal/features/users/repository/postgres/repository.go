package users_postgres_repository

import core_postgres_pull "github.com/Ravenmax/ToDo/internal/core/repository/postgres/pull"

type UserRepository struct {
	pool core_postgres_pull.Pool
}

func NewUserRepository(
	pool core_postgres_pull.Pool,
) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
