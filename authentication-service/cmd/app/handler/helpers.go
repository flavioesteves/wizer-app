package handler

import (
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
)

type ServerConfig struct {
	pb.AuthenticationServiceServer
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
}
