package authservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var AUTH_SERVICE_HOST = "authentication-service:50055"

func Connect() pb.AuthenticationServiceClient {
	conn, err := grpc.NewClient(AUTH_SERVICE_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
	}

	fmt.Println("Auth service connected")
	return pb.NewAuthenticationServiceClient(conn)
}
