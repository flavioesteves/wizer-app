package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/flavioesteves/wizer-app/authentication/internal/database"
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

	//token := middleware.GenerateToken()
	token, err := GenerateToken(in.GetEmail())
	if err != nil {

		fmt.Printf("GenerateToke: %v\n", err)

		return nil, err
	}

	sessionData := "user_session:" + in.GetEmail()

	err = s.redisClient.Set(ctx, token, sessionData, time.Hour*24).Err()
	if err != nil {
		return nil, err
	}

	s.redisClient.Save(ctx)

	test, err := s.GetSessionData(ctx, token)
	fmt.Printf("TEST: %v\n", test)

	return &pb.LoginResponse{IsValid: true, Token: token}, nil
}
