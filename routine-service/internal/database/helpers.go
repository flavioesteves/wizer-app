package database

import (
	pb "github.com/flavioesteves/wizer-app/routine/proto"
	"strings"
)

func SliceToString(exercises []*pb.Exercise) string {
	var sb strings.Builder
	for _, exercise := range exercises {
		// Convert exercise to string (adjust this based on pb.exercise structure)
		sb.WriteString(exercise.String())
		sb.WriteString(", ")
	}
	return sb.String()
}
