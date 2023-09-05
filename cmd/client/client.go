package main

import (
	"context"
	"fmt"
	"goproj/pkg/userservice"
	"goproj/utils"
	"log"
	"time"

	"google.golang.org/grpc"
)

const email = "example@example.com"

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	utils.LogError(err)
	defer conn.Close()

	uc := userservice.NewUserServiceClient(conn)

	user := &userservice.User{
		Email:    email,
		FullName: "John Doe",
		IsAdmin:  true,
	}

	us, err := uc.CreateUser(ctx, user)
	if err != nil {
		log.Fatalln(err)
	} else if us.Success {
		fmt.Println("User was successfully created")
	}

	users, err := uc.GetUserList(ctx, &userservice.Empty{})
	utils.LogError(err)
	fmt.Println("user list", users)

	u, err := uc.GetUserByEmail(ctx, &userservice.EmailFilter{Email: email})
	utils.LogError(err)
	fmt.Println("user", u)
}
