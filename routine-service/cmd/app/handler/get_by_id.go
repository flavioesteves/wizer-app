package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) GetRoutine(ctx context.Context, in *pb.GetRoutineRequest) (*pb.RoutineResponse, error) {
	fmt.Printf("Get Routine was invoked with %v\n", in)

	routineId := in.GetId()

	routine, err := db.Get(s.db, routineId)
	if err != nil {
		fmt.Printf("failed to get the routine with id: %s\n", err)
	}

	return &pb.RoutineResponse{
		Routine: routine,
	}, nil
}
