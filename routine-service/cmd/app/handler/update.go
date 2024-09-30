package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) UpdateRoutine(ctx context.Context, in *pb.UpdateRoutineRequest) (*pb.RoutineResponse, error) {

	fmt.Printf("UpdatedRequest was invoked with %v\n", in)

	updateRoutine := &pb.Routine{
		Id: in.Routine.Id,
	}

	routine, err := db.Update(s.db, updateRoutine)
	if err != nil {
		fmt.Printf("failed to update the routine: %v", err)
	}

	return &pb.RoutineResponse{
		Routine: routine,
	}, nil
}
