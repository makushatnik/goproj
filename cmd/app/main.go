package main

import (
	"fmt"
	"goproj/internal/config"
	"goproj/internal/generated/restapi"
	"goproj/internal/generated/restapi/operations"
	"goproj/internal/handler/proto/order"
	"goproj/internal/handler/proto/user"
	"goproj/internal/interceptor"
	"goproj/pkg/orderservice"
	"goproj/pkg/userservice"
	"goproj/utils"
	"log"
	"net"
	"os"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	// grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
)

func GRPCServer(cfg *config.Config) {
	fmt.Println(fmt.Sprintf("Starting server... tcp://%s:%d", cfg.Host, cfg.GRPCPort))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	utils.LogError(err)

	// s := grpc.NewServer()
	// server := proto.NewServer()
	// userservice.RegisterUserServiceServer(s, server)
	// s := grpc.NewServer(grpc_middleware.WithStreamServerChain(interceptor.LogInterceptor))
	s := grpc.NewServer(grpc.ChainStreamInterceptor(interceptor.LogInterceptor))
	userservice.RegisterUserServiceServer(s, user.New())
	orderservice.RegisterOrderServiceServer(s, order.New())
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}

func HTTPServer(cfg *config.Config) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	utils.LogError(err)

	api := operations.NewWebApplicationAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "webapp"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		utils.LogError(err)

	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.Host = cfg.Host
	server.Port = cfg.HTTPPort
	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create a config: %v\n", err)
		os.Exit(1)
	}

	GRPCServer(cfg)
	// HTTPServer(cfg)
}
