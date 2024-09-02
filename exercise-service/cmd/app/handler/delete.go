package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ServerConfig) DeleteExercise(ctx context.Context, in *pb.DeleteExerciseRequest) (*emptypb.Empty, error) {
	exerciseId := in.GetId()

	err := db.Delete(s.db, exerciseId)
	if err != nil {
		fmt.Printf("failed to delete the exercise: %v", err)
	}
	return nil, nil
}
