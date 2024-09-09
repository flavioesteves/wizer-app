package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/flavioesteves/wizer-app/authentication/proto"
	"github.com/redis/go-redis/v9"
)

func (s *ServerConfig) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	token := in.GetToken()

	sessionData, err := s.GetSessionData(ctx, token)

	if err != nil {
		if err == redis.Nil {
			return &pb.ValidateTokenResponse{IsValid: false, Message: "Invalid Tokes"}, nil
		}
		return nil, err
	}

	userID, err := strconv.Atoi(strings.TrimPrefix(sessionData, "user_session:"))
	if err != nil {
		return nil, err
	}
	fmt.Println(userID)

	return &pb.ValidateTokenResponse{IsValid: true}, nil
}
