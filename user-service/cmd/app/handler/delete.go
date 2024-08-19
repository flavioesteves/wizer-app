package handler

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"

	db "github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func (s *ServerConfig) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	fmt.Printf("Delete user was invoked with %v\n", in)

	userId := in.UserId

	err := db.Delete(s.db, userId)
	if err != nil {
		fmt.Printf("Failed to delete the user: %v\n", err)
	}
	return nil, nil
}
