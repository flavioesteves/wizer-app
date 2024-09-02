package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func Get(db *sql.DB, id string) (*pb.Exercise, error) {

	exercise := &pb.Exercise{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
    SELECT *
    FROM exercises
    WHERE id = $1`

	var tempExercises []byte
	err := db.QueryRowContext(ctx, query, id).Scan(
		&exercise.Id,
		&exercise.Name,
		&exercise.MuscleGroup,
		&exercise.Description,
		&tempExercises,
		&exercise.VideoUrl,
		&exercise.VideoDurationSeconds,
		&exercise.CreatedBy,
		&exercise.UpdatedBy,
		&exercise.CreatedAt,
		&exercise.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get exercise wiht id :%s\n with error: $w\n", id, err)
	}

	var stepSlice []*pb.Step
	err = json.Unmarshal(tempExercises, &stepSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal steps into JSON: %w\n", err)
	}

	exercise.Steps = stepSlice

	return exercise, err
}
