package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/profile/internal/database"
	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) GetAllProfiles(ctx context.Context, in *pb.GetAllProfilesRequest) (*pb.GetAllProfilesResponse, error) {
	fmt.Printf("Get all profiles retrieved: %v\n", in)

	profiles, err := db.GetAll(s.db)

	if err != nil {
		fmt.Printf("Failed to retrieve all the profiles %v\n", profiles)
	}

	return &pb.GetAllProfilesResponse{
		Profiles: profiles,
	}, nil
}
