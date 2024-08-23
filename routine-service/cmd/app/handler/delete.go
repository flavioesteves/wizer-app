package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerConfig) DeleteRoutine(ctx context.Context, in *pb.DeleteRoutineRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete routine was invoked with %v\n", in)

	routineId := in.Id

	err := db.Delete(s.db, routineId)
	if err != nil {
		fmt.Printf("failed to delete the routine: %v\n", err)
	}
	return nil, nil
}
