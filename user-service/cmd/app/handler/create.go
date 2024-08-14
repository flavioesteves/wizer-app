package handler

import (
	"context"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	fmt.Printf("CreateUser was invoked with %v\n", in)

	userResponse := &pb.User{
		UserId:       in.User.UserId,
		Email:        in.User.Email,
		PasswordHasg: in.User.PasswordHasg,
	}

	return &pb.UserResponse{
		User: userResponse,
	}, nil

}
