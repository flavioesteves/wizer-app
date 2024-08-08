package main

import (
	"fmt"

	pb "github.com/flavioesteves/wizer-app/user/proto"
	"golang.org/x/net/context"
)

func (s *Server) CreateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	fmt.Printf("CreateUser was invoked with %v\n", in)

	return &pb.User{Email: in.Email, Password: in.Password, Id: in.Id}, nil
}
