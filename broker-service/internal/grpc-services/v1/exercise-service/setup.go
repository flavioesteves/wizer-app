package exerciseservice

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var EXERCISE_SERVICE_HOST = "exercise-service:50054"

func Connect() pb.ExerciseServiceClient {
	conn, err := grpc.NewClient(EXERCISE_SERVICE_HOST, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
	}

	fmt.Println("Exercise service connected")

	return pb.NewExerciseServiceClient(conn)
}
