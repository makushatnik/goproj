package handler

import "goproj/internal/handler/http/user"

type Handlers struct {
	User user.Handlers
}

func New() *Handlers {
	return &Handlers{
		User: user.New(),
	}
}
