package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func Update(db *sql.DB, exercise *pb.Exercise) (*pb.Exercise, error) {
	updatedExercise := &pb.Exercise{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
    UPDATE exercises
    SET name=$1, muscle_group=$2, description=$3 ,steps=$4, video_url=$5, video_duration_seconds=$6, updated_by=$7, updated_at=now()
    WHERE id=$8
    RETURNING *`

	stepsJSON, err := json.Marshal(exercise.Steps)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal steps into JSON: %w\n", err)
	}

	args := []any{
		exercise.Name,
		exercise.MuscleGroup,
		exercise.Description,
		stepsJSON,
		exercise.VideoUrl,
		exercise.VideoDurationSeconds,
		exercise.UpdatedBy,
		exercise.Id,
	}

	row := db.QueryRowContext(ctx, query, args...)
	var tempExercises []byte

	err = row.Scan(
		&updatedExercise.Id,
		&updatedExercise.Name,
		&updatedExercise.MuscleGroup,
		&updatedExercise.Description,
		&tempExercises,
		&updatedExercise.VideoUrl,
		&updatedExercise.VideoDurationSeconds,
		&updatedExercise.CreatedBy,
		&updatedExercise.UpdatedBy,
		&updatedExercise.CreatedAt,
		&updatedExercise.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed ot update exercise: %w\n", err)
	}
	var stepSlice []*pb.Step

	err = json.Unmarshal(tempExercises, &stepSlice)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal steps into JSON: %w\n", err)
	}

	updatedExercise.Steps = stepSlice

	fmt.Printf("Updated Exercise : %v", updatedExercise)
	return updatedExercise, err

}
