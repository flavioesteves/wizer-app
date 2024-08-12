package handler

import (
	"context"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) CreateProfile(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	fmt.Printf("CreateUser was invoked with %v\n", in)

	return &pb.Profile{Email: in.Email, Password: in.Password, Id: in.Id}, nil
}
