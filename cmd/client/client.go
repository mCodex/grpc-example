package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mcodex/grpc-example/pb"
	"google.golang.org/grpc"
)

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Mateus",
		Email: "m@m.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	fmt.Println(res)
}

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)
}
