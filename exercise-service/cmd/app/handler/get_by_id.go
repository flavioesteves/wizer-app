package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func (s *ServerConfig) GetExercise(ctx context.Context, in *pb.GetExerciseRequest) (*pb.ExerciseResponse, error) {

	exerciseId := in.GetId()

	exercise, err := db.Get(s.db, exerciseId)
	if err != nil {
		fmt.Printf("failed ot get the exercise with id: %s\n", err)
	}

	return &pb.ExerciseResponse{
		Exercise: exercise,
	}, nil
}
