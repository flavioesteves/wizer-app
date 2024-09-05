package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	fmt.Printf("CreateUser was invoked with %v\n", in)

	newUser := &pb.User{
		Email:    in.User.Email,
		Password: in.User.Password,
		Role:     in.User.Role,
	}

	user, err := db.Insert(s.db, newUser)
	if err != nil {
		fmt.Printf("Failed to create the user: %v\n", err)
	}

	return &pb.UserResponse{User: user}, nil

}
