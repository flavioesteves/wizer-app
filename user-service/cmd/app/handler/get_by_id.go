package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	fmt.Printf("Get User was invoked with %v\n", in)

	userId := in.GetUserId()

	user, err := db.Get(s.db, userId)
	if err != nil {
		fmt.Printf("Failed to get the user with id: %s\n", err)
	}

	return &pb.UserResponse{
		User: user,
	}, nil
}
