package repository

import (
	"goproj/internal/db/pg"
	"goproj/internal/repository/user"
)

type Container struct {
	User user.UserRepository
}

func New(poolConn pg.Repository) *Container {
	return &Container{
		User: user.New(poolConn),
	}
}
