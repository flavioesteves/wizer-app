package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) AddExerciseToRoutine(ctx context.Context, in *pb.AddExerciseToRoutineRequest) (*pb.AddExerciseToRoutineResponse, error) {

	err := db.InsertExerciseToRoutine(s.db, in.GetRoutineId(), in.GetExerciseIds())
	if err != nil {
		return &pb.AddExerciseToRoutineResponse{
				Success: false,
			},
			fmt.Errorf("failed to create the routine_exercise: %w\n", err)
	}

	fmt.Printf("AddService was invoked with %v\n", in)

	return &pb.AddExerciseToRoutineResponse{
		Success: true,
	}, nil
}
