package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/flavioesteves/wizer-app/authentication/internal/middleware"
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
)

func (s *ServerConfig) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("Login was called with in:%v\n", in)

	token := middleware.GenerateToken()
	// Validate the user service if the user exists?
	session := s.redisClient.Set(ctx, token, "user_id:1", time.Hour*24)
	fmt.Println(session)
	s.redisClient.Save(ctx)
	return &pb.LoginResponse{Token: token}, nil
}
