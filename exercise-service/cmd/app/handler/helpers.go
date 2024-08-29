package handler

import (
	"database/sql"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

type ServerConfig struct {
	db *sql.DB
	pb.ExerciseServiceServer
}

func NewServerConfig(db *sql.DB) *ServerConfig {
	return &ServerConfig{db: db}
}
