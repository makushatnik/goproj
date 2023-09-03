package pg

import (
	"context"
	"goproj/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetPoolConnection() *pgxpool.Pool
}

type repository struct {
	pool *pgxpool.Pool
}

func NewRepository(ctx context.Context, cfg *config.Config) (Repository, error) {
	pool, err := NewPool(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &repository{pool: pool}, nil
}

func (r *repository) GetPoolConnection() *pgxpool.Pool {
	return r.pool
}
