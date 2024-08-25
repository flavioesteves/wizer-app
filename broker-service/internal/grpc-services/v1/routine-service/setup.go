package routineservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ROUTINE_SERVICE_HOST = "routine-service:50053"

func Connect() pb.RoutineServiceClient {
	conn, err := grpc.NewClient(ROUTINE_SERVICE_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("failed to connect: %v\n", err)
	}

	return pb.NewRoutineServiceClient(conn)
}
