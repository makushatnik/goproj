package server

import (
	"context"
	"fmt"
	"goproj/internal/config"
	"goproj/internal/db/pg"
	"net/http"
)

func Server() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithCancel(context.Background())
	// conn, err := pg.New(cfg)
	// if err != nil {
	// 	panic(err)
	// }
	pool, err := pg.NewPool(ctx, cfg)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// defer func() {
	// 	if err := conn.Close(context.Background()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	fmt.Println("Postgres connected!")
	_ = http.ListenAndServe(":8080", nil)
}
