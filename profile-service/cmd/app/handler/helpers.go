package handler

import (
	"database/sql"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

type ServerConfig struct {
	db *sql.DB
	pb.ProfileServiceServer
}

func NewServerConfig(db *sql.DB) *ServerConfig {
	return &ServerConfig{db: db}
}
