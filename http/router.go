package http

import (
	"goproj/http/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/order/{id:[0-9]+}", handler.GetOrderById).Methods(http.MethodGet)
	r.HandleFunc("/user/create", handler.ParsePostUser).Methods(http.MethodPost)
	// orderRouter.Methods(http.MethodPost)

	r.Use(mux.CORSMethodMiddleware(r))
	return r
}
