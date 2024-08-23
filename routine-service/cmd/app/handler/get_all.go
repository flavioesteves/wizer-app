package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func (s *ServerConfig) GetAllRoutines(ctx context.Context, in *pb.GetAllRoutinesRequest) (*pb.GetAllRoutinesResponse, error) {

	fmt.Printf("Get all routines retrieved: %v\n", in)

	routines, err := db.GetAll(s.db)
	if err != nil {
		fmt.Printf("failed to retrieve all the routines %v\n", routines)
	}

	return &pb.GetAllRoutinesResponse{
		Routines: routines,
	}, nil
}
