package server

import (
	"context"
	"database/sql"
	"fmt"
	myhttp "goproj/http"
	"goproj/internal/config"
	"goproj/internal/db/pg"
	"goproj/internal/repository"
	"net/http"
	"os"

	"github.com/pressly/goose"
)

type WebHandler struct {
	message string
}

func (wh WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, wh.message)
}

func Server() {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a config: %v\n", err)
		os.Exit(1)
	}

	ctx, _ := context.WithCancel(context.Background())
	// pool, err := pg.NewPool(ctx, cfg)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to create a pool: %v\n", err)
	// 	os.Exit(1)
	// }
	// defer pool.Close()
	pgRepo, err := pg.NewRepository(ctx, cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a repository: %v\n", err)
		os.Exit(1)
	}
	defer pgRepo.GetPoolConnection().Close()

	migrationsUp(pgRepo.GetPoolConnection().Config().ConnString())

	repoContainer := repository.New(pgRepo)
	repoContainer.User.List(ctx)

	fmt.Println("Postgres connected!")

	// http.HandleFunc("/", handler.Index)
	// http.Handle("/hello", WebHandler{"Webapp on Go"})
	// http.Handle("/redirect", http.RedirectHandler("/", http.StatusMovedPermanently))
	// http.HandleFunc("/order", handler.ParseGetOrder)
	// http.HandleFunc("/send/message", handler.ParseSendMessage)
	// http.HandleFunc("/user/create", handler.ParsePostUser)
	// http.HandleFunc("/secure", handler.BasicAuth)
	// http.HandleFunc("/json/response", handler.JsonResponse)
	// r := mux.NewRouter()
	// r.HandleFunc("/", handler.Index)

	_ = http.ListenAndServe(":8080", myhttp.NewRouter())
}

func migrationsUp(connString string) {
	mdb, _ := sql.Open("postgres", connString)
	err := mdb.Ping()
	if err != nil {
		panic(err)
	}

	err = goose.Up(mdb, "/migrations")
	if err != nil {
		panic(err)
	}
}
