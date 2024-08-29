package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/exercise/proto"
	_ "github.com/lib/pq"
)

const DB_TIMEOUT = time.Second * 5

func Insert(db *sql.DB, exercise *pb.Exercise) (*pb.Exercise, error) {
	newExercise := &pb.Exercise{}

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	stmt := `
  INSERT INTO exercises
  (name, muscle_group, description, steps, video_url, video_duration_seconds, created_by, updated_by, created_at, updated_at)
  VALUES($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
  RETURNING id, name, muscle_group, description, steps, video_url, video_duration_seconds, created_by, updated_by, created_at, updated_at`

	stepsJSON, err := json.Marshal(exercise.Steps)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal steps into JSON: %w\n", err)
	}

	row := db.QueryRowContext(ctx, stmt,
		&exercise.Name,
		&exercise.MuscleGroup,
		&exercise.Description,
		stepsJSON,
		&exercise.VideoUrl,
		&exercise.VideoDurationSeconds,
		&exercise.CreatedBy,
		&exercise.UpdatedBy,
	)

	var tempExercises []byte
	err = row.Scan(
		&newExercise.Id,
		&newExercise.Name,
		&newExercise.MuscleGroup,
		&newExercise.Description,
		&tempExercises,
		&newExercise.VideoUrl,
		&newExercise.VideoDurationSeconds,
		&newExercise.CreatedBy,
		&newExercise.UpdatedBy,
		&newExercise.CreatedAt,
		&newExercise.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert exercise: %w\n", err)
	}

	var stepSlice []*pb.Step
	err = json.Unmarshal(tempExercises, &stepSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal steps into JSON: %w\n", err)
	}

	newExercise.Steps = stepSlice

	fmt.Printf("New Exercise was created: %v\n", newExercise)

	return newExercise, err
}
