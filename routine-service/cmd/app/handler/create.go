package handler

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) CreateRoutine(ctx context.Context, in *pb.CreateRoutineRequest) (*pb.RoutineResponse, error) {
	fmt.Printf("CreateRoutine was invoked with %v\n", in)

	e, err := json.Marshal(in.Routine.Exercises)
	if err != nil {
		fmt.Println("failed to convert into json")
	}

	newRoutine := &pb.Routine{
		ProfileId: in.Routine.ProfileId,
		//	Exercises: in.Routine.Exercises,
		Exercises: e,
	}

	routine, err := db.Insert(s.db, newRoutine)
	if err != nil {
		fmt.Printf("failed to create the routine: %v\n", err)
	}

	return &pb.RoutineResponse{
		Routine: routine,
	}, nil
}
