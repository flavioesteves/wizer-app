package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func (s *ServerConfig) GetAllExercises(ctx context.Context, in *pb.GetAllExercisesRequest) (*pb.GetAllExercisesResponse, error) {

	exercises, err := db.GetAll(s.db)
	if err != nil {
		fmt.Printf("failed to retrieve all exercises %v\n", exercises)
	}

	return &pb.GetAllExercisesResponse{
		Exercises: exercises,
	}, nil
}
