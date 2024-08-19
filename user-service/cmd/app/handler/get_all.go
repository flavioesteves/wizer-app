package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

// rpc GetAllUsers (GetAllUsersRequest) returns (GetAllUsersResponse);
func (s *ServerConfig) GetAllUsers(ctx context.Context, in *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {

	users, err := db.GetAll(s.db)

	if err != nil {
		fmt.Printf("Failed to retrieve all the the users %v\n", users)
	}

	return &pb.GetAllUsersResponse{
		Users: users,
	}, nil

}
