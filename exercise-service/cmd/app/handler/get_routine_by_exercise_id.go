package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func (s *ServerConfig) GetRoutinesByExerciseId(ctx context.Context, in *pb.GetRoutinesByExerciseIdRequest) (*pb.GetRoutinesByExerciseIdResponse, error) {

	routines, err := db.GetRoutinesByExerciseId(s.db, in.GetExerciseId())
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve the routines: %w\n", err)
	}

	return &pb.GetRoutinesByExerciseIdResponse{
		Routines: routines,
	}, nil
}
