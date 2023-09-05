package main

import (
	"context"
	"fmt"
	"goproj/pkg/orderservice"
	"goproj/utils"
	"io"
	"time"

	"google.golang.org/grpc"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	utils.LogError(err)
	defer conn.Close()

	oc := orderservice.NewOrderServiceClient(conn)
	orderStream, err := oc.GetOrders(ctx, &orderservice.Empty{})
	utils.LogError(err)

	for {
		order, err := orderStream.Recv()
		if err == io.EOF {
			break
		}
		utils.LogError(err)
		fmt.Println(order)
		time.Sleep(1 * time.Second)
	}
}
