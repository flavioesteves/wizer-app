package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) RemoveExerciseFromRoutine(ctx context.Context, in *pb.RemoveExerciseFromRoutineRequest) (*pb.RemoveExerciseFromRoutineResponse, error) {

	err := db.RemoveExerciseFromRoutine(s.db, in.GetRoutineId(), in.GetExerciseIds())
	if err != nil {
		return &pb.RemoveExerciseFromRoutineResponse{
				Success: false,
			},
			fmt.Errorf("failed to remove the routine_exercise: %w\n", err)
	}

	return &pb.RemoveExerciseFromRoutineResponse{
		Success: true,
	}, nil
}
