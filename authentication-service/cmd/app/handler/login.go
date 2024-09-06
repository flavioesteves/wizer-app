package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/flavioesteves/wizer-app/authentication/internal/database"
	"github.com/flavioesteves/wizer-app/authentication/internal/middleware"
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
)

func (s *ServerConfig) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Printf("Login was called with in:%v\n", in)

	validateUser, err := database.GetByEmail(s.db, in.GetEmail())
	if err != nil {
		return nil, err
	}

	if validateUser.Email != in.GetEmail() {
		return nil, err
	}

	err = DecryptPassword(validateUser.Password, in.GetPassword())

	if err != nil {
		fmt.Printf("Login err DecryptPassword %v\n", err)
		return nil, err
	}

	token := middleware.GenerateToken()
	session := s.redisClient.Set(ctx, token, "user_session:"+in.GetEmail(), time.Hour*24)
	fmt.Println(session)
	s.redisClient.Save(ctx)
	return &pb.LoginResponse{IsValid: true, Token: token}, nil
}
