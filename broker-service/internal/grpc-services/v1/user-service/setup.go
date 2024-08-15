package userservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var USER_SERVICE_HOST = "user-service:50052"

func Connect() pb.UserServiceClient {
	conn, err := grpc.NewClient(USER_SERVICE_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}

	return pb.NewUserServiceClient(conn)
}
