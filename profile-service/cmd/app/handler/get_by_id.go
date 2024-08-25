package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/profile/internal/database"
	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) GetProfile(ctx context.Context, in *pb.GetProfileRequest) (*pb.ProfileResponse, error) {
	fmt.Printf("GetProfile was invoked with %v\n", in)

	profileId := in.Id

	profile, err := db.Get(s.db, profileId)
	if err != nil {
		fmt.Printf("Failed to get the profile with id: %s\n", err)
	}

	return &pb.ProfileResponse{
		Profile: profile,
	}, nil
}
