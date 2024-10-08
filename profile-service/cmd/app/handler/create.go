package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/profile/internal/database"
	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) CreateProfile(ctx context.Context, in *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	fmt.Printf("CreateProfile was invoked with %v\n", in)

	newProfile := &pb.Profile{
		Id:                in.Profile.Id,
		UserId:            in.Profile.UserId,
		Gender:            in.Profile.Gender,
		BirthYear:         in.Profile.BirthYear,
		HeightCm:          in.Profile.HeightCm,
		WeightKg:          in.Profile.WeightKg,
		BodyFatPercentage: in.Profile.BodyFatPercentage,
		Goal:              in.Profile.Goal,
	}

	profile, err := db.Insert(s.db, newProfile)
	if err != nil {
		fmt.Printf("Failed to create the profile: %v\n", err)
	}

	return &pb.ProfileResponse{
		Profile: profile,
	}, nil
}
