package database

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func GetExercisesByRoutineId(db *sql.DB, routineID string) ([]*pb.Exercise, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	stmt := `
    SELECT e.id, e.name, e.muscle_group, e.description, e.steps, e.video_url, e.video_duration_seconds
    FROM routine_exercises re
    JOIN exercises e ON re.exercise_id = e.id
    WHERE re.routine = $1
  `

	rows, err := db.QueryContext(ctx, stmt, routineID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve exercises for routine %s: %w", routineID, err)
	}
	defer rows.Close()

	var exercises []*pb.Exercise

	for rows.Next() {
		exercise := &pb.Exercise{}
		err := rows.Scan(
			&exercise.Id,
			&exercise.Name,
			&exercise.MuscleGroup,
			&exercise.Description,
			&exercise.Steps,
			&exercise.VideoUrl,
			&exercise.VideoDurationSeconds,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan exercise: %w", err)
		}

		exercises = append(exercises, exercise)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %w", err)
	}

	return exercises, nil

}
