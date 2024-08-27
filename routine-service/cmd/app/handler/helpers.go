package handler

import (
	"database/sql"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

type ServerConfig struct {
	db *sql.DB
	pb.RoutineServiceServer
}

func NewServerConfig(db *sql.DB) *ServerConfig {
	return &ServerConfig{db: db}
}
