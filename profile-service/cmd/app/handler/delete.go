package handler

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	db "github.com/flavioesteves/wizer-app/profile/internal/database"
	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func (s *ServerConfig) DeleteProfile(ctx context.Context, in *pb.DeleteProfileRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete profile was invoked with %v\n", in)

	profileId := in.ProfileId

	err := db.Delete(s.db, profileId)
	if err != nil {
		fmt.Printf("Failed to delete the profile: %v\n", err)
	}
	return nil, nil
}
