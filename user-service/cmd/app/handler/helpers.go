package handler

import (
	"database/sql"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

type ServerConfig struct {
	db *sql.DB
	pb.UserServiceServer
}

func NewServerConfig(db *sql.DB) *ServerConfig {
	return &ServerConfig{db: db}
}

