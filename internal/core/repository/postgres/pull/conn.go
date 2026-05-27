package core_postgres_pull

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
	OpTimeout() time.Duration
}

type ConnectionPull struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewConnectionPool(
	ctx context.Context,
	config Config,
) (*ConnectionPull, error) {
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
		return nil, fmt.Errorf("parse pgxconfig: %w", err)
	}
	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgxconfig: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping pool: %w", err)
	}
	return &ConnectionPull{
		Pool:      pool,
		opTimeout: config.TimeOut,
	}, nil
}

func (p *ConnectionPull) OpTimeout() time.Duration {
	return p.opTimeout
}
