package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func (s *ServerConfig) CreateExercise(ctx context.Context, in *pb.CreateExerciseRequest) (*pb.ExerciseResponse, error) {
	fmt.Printf("CreateExercise was invoked with %v\n", in)

	newExercise := &pb.Exercise{
		Name:                 in.Exercise.Name,
		MuscleGroup:          in.Exercise.MuscleGroup,
		Description:          in.Exercise.Description,
		Steps:                in.Exercise.Steps,
		VideoUrl:             in.Exercise.VideoUrl,
		VideoDurationSeconds: in.Exercise.VideoDurationSeconds,
	}

	exercise, err := db.Insert(s.db, newExercise)
	if err != nil {
		fmt.Printf("failed to create the exercise: %v\n", err)
	}

	return &pb.ExerciseResponse{
		Exercise: exercise,
	}, nil

}
