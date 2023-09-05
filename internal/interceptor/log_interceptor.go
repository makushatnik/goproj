package interceptor

import (
	"fmt"

	"google.golang.org/grpc"
)

func LogInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println("Log Interceptor")
	return handler(srv, ss)
}
