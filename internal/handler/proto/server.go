package proto

import (
	"goproj/internal/handler/proto/user"
	"goproj/pkg/userservice"
)

type Server struct {
	user.Handlers
	userservice.UnsafeUserServiceServer
}

func NewServer() *Server {
	return &Server{Handlers: user.New()}
}
