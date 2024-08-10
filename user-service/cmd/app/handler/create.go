package handler

import (
	"context"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	fmt.Printf("CreateUser was invoked with %v\n", in)

	return &pb.User{Email: in.Email, Password: in.Password, Id: in.Id}, nil
}
