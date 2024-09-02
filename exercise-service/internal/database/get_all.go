package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func GetAll(db *sql.DB) ([]*pb.Exercise, error) {
	var exercises []*pb.Exercise

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT * FROM exercises ORDER by created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		exercise := &pb.Exercise{}
		var tempExercises []byte

		err := rows.Scan(
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
			fmt.Println("error scanning", err)
			return nil, err
		}

		var stepSlice []*pb.Step
		err = json.Unmarshal(tempExercises, &stepSlice)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal steps into JSON: %w\n", err)
		}
		exercise.Steps = stepSlice
		exercises = append(exercises, exercise)
	}
	return exercises, nil
}
