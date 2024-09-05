package handler

import (
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
	"github.com/redis/go-redis/v9"
)

type ServerConfig struct {
	redisClient *redis.Client
	pb.AuthenticationServiceServer
}

func NewServerConfig(redisClient *redis.Client) *ServerConfig {
	return &ServerConfig{redisClient: redisClient}
}
