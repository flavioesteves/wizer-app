package handler

import (
	"database/sql"
	"strings"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

type ServerConfig struct {
	db *sql.DB
	pb.RoutineServiceServer
}

func NewServerConfig(db *sql.DB) *ServerConfig {
	return &ServerConfig{db: db}
}

func SliceToString(exercises []*pb.Exercise) string {
	var sb strings.Builder
	for _, exercise := range exercises {
		// Convert exercise to string (adjust this based on pb.exercise structure)
		sb.WriteString(exercise.String())
		sb.WriteString(", ")
	}
	return sb.String()
}
