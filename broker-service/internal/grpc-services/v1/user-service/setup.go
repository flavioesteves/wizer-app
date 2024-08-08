package userservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(addr string) pb.UserServiceClient {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}

	return pb.NewUserServiceClient(conn)
}
