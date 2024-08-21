package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UserResponse, error) {

	fmt.Printf("UpdatedUser was invoked with %v\n", in)

	updateUser := &pb.User{
		Id:       in.User.Id,
		Email:    in.User.Email,
		Password: in.User.Password,
		Role:     in.User.Role,
	}

	user, err := db.Update(s.db, updateUser)
	if err != nil {
		fmt.Printf("Failed to update the user: %v\n", err)
	}

	return &pb.UserResponse{
		User: user,
	}, nil
}
