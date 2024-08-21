package handler

import (
	"context"
	"fmt"

	db "github.com/flavioesteves/wizer-app/profile/internal/database"
	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) UpdateProfile(ctx context.Context, in *pb.UpdateProfileRequest) (*pb.ProfileResponse, error) {
	fmt.Printf("UpdateProfile was invoked with %v\n", in)

	updatedProfile := &pb.Profile{
		Id:                in.Profile.Id,
		UserId:            in.Profile.UserId,
		Gender:            in.Profile.Gender,
		BirthYear:         in.Profile.BirthYear,
		HeightCm:          in.Profile.HeightCm,
		WeightKg:          in.Profile.WeightKg,
		BodyFatPercentage: in.Profile.BodyFatPercentage,
		Goal:              in.Profile.Goal,
	}

	profile, err := db.Update(s.db, updatedProfile)
	if err != nil {
		fmt.Printf("Failed to update the profile: %v\n", err)
	}

	return &pb.ProfileResponse{
		Profile: profile,
	}, nil
}
