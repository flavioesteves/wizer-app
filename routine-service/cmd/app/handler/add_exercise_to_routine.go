package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) AddExerciseToRoutine(ctx context.Context, in *pb.GetExercisesByRoutineIdRequest) (*pb.AddExercissToRoutineResponse, error) {

	fmt.Printf("AddService was invoked with %v\n", in)

	return &pb.AddExercissToRoutineResponse{
		Success: true,
	}, nil
}
