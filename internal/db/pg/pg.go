package pg

import (
	"context"
	"goproj/internal/config"

	"github.com/jackc/pgx/v5"
)

func New(cfg *config.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DBDSN)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
