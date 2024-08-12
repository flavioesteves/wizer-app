package userservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var PROFILE_SERVICE_HOST = "profile-service:50051"

func Connect() pb.ProfileServiceClient {
	conn, err := grpc.NewClient(PROFILE_SERVICE_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}

	return pb.NewProfileServiceClient(conn)
}
