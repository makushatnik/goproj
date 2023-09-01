package pg

import (
	"context"
	"fmt"
	"goproj/internal/config"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	poolConfig, _ := pgxpool.ParseConfig(cfg.DBDSN)
	poolConfig.MaxConns = 5

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	errChan := make(chan error, poolConfig.MaxConns)
	for i := 0; i < int(poolConfig.MaxConns); i++ {
		wg.Add(1)
		go func(errChan chan<- error, count int) {
			defer wg.Done()

			_, err = pool.Exec(ctx, ";")
			if err != nil {
				errChan <- err
			}

			fmt.Println(count, "Query OK!")
			fmt.Printf("Connections - Max: %d, Idle: %d, Total: %d\n",
				pool.Stat().MaxConns(),
				pool.Stat().IdleConns(),
				pool.Stat().TotalConns())
		}(errChan, i)
	}

	wg.Wait()
	select {
	case err := <-errChan:
		panic(err)
	default:
		close(errChan)
		return pool, nil
	}
}
