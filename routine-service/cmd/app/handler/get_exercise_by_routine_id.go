package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) GetExercisesByRoutineId(ctx context.Context, in *pb.GetExercisesByRoutineIdRequest) (*pb.GetExercisesByRoutineIdResponse, error) {

	exercises, err := db.GetExercisesByRoutineId(s.db, in.GetRoutineId())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the exercises: %w\n", err)
	}

	return &pb.GetExercisesByRoutineIdResponse{
		Exercises: exercises,
	}, nil
}
