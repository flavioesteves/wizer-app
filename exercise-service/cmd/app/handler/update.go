package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func (s *ServerConfig) UpdateExercise(ctx context.Context, in *pb.UpdateExerciseRequest) (*pb.ExerciseResponse, error) {

	updateExercise := &pb.Exercise{
		Id:                   in.Exercise.Id,
		Name:                 in.Exercise.Name,
		Description:          in.Exercise.Description,
		MuscleGroup:          in.Exercise.MuscleGroup,
		Steps:                in.Exercise.Steps,
		VideoUrl:             in.Exercise.VideoUrl,
		VideoDurationSeconds: in.Exercise.VideoDurationSeconds,
		UpdatedBy:            in.Exercise.UpdatedBy,
	}

	exercise, err := db.Update(s.db, updateExercise)
	if err != nil {
		fmt.Printf("failed to update the exercise: %v", err)
	}

	return &pb.ExerciseResponse{
		Exercise: exercise,
	}, nil

}
