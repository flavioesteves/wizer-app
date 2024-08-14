package handler

import (
	"context"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) CreateProfile(ctx context.Context, in *pb.CreateProfileRequest) (*pb.ProfileResponse, error) {
	fmt.Printf("CreateProfile was invoked with %v\n", in)

	profileResponse := &pb.Profile{
		ProfileId:         in.Profile.ProfileId,
		UserId:            in.Profile.UserId,
		Gender:            in.Profile.Gender,
		BirthYear:         in.Profile.BirthYear,
		HeightCm:          in.Profile.HeightCm,
		WeightKg:          in.Profile.WeightKg,
		BodyFatPercentage: in.Profile.BodyFatPercentage,
		Goal:              in.Profile.Goal,
	}

	return &pb.ProfileResponse{
		Profile: profileResponse,
	}, nil
}
