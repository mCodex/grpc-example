package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mcodex/grpc-example/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	// AddUsers(client)
	AddUserStreamBoth(client)
}

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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Mateus",
		Email: "m@m.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC Request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the message: %v", err)

		}

		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "User1",
			Email: "mail1@mail.com",
		},
		{
			Id:    "2",
			Name:  "User2",
			Email: "mail2@mail.com",
		},
		{
			Id:    "3",
			Name:  "User3",
			Email: "mail3@mail.com",
		},
		{
			Id:    "4",
			Name:  "User4",
			Email: "mail4@mail.com",
		},
		{
			Id:    "5",
			Name:  "User5",
			Email: "mail5@mail.com",
		},
		{
			Id:    "6",
			Name:  "User6",
			Email: "mail6@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)

}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())

	reqs := []*pb.User{
		{
			Id:    "1",
			Name:  "User1",
			Email: "mail1@mail.com",
		},
		{
			Id:    "2",
			Name:  "User2",
			Email: "mail2@mail.com",
		},
		{
			Id:    "3",
			Name:  "User3",
			Email: "mail3@mail.com",
		},
		{
			Id:    "4",
			Name:  "User4",
			Email: "mail4@mail.com",
		},
		{
			Id:    "5",
			Name:  "User5",
			Email: "mail5@mail.com",
		},
		{
			Id:    "6",
			Name:  "User6",
			Email: "mail6@mail.com",
		},
	}

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error receiving data: %v", err)
				break
			}

			fmt.Printf("Recebendo user %v com status: %v\n", res.GetUser().GetName(), res.GetStatus())
		}

		close(wait)
	}()

	<-wait
}
