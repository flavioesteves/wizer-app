package database

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

func GetRoutinesByExerciseId(db *sql.DB, exerciseID string) ([]*pb.Routine, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	stmt := `
    SELECT r.id, r.name, r.profile_id, r.created_by, r.updated_by, r.created_at, r.updated_at
		FROM routine_exercises re
		JOIN routines r ON re.routine_id = r.id
		WHERE re.exercise_id = $1`

	rows, err := db.QueryContext(ctx, stmt, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve routienes for exercises %s: %w:\n", exerciseID, err)
	}
	defer rows.Close()

	var routines []*pb.Routine

	for rows.Next() {
		routine := &pb.Routine{}
		err := rows.Scan(
			&routine.Id,
			&routine.Name,
			&routine.ProfileId,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan routines %w", err)
		}
		routines = append(routines, routine)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during rows iteration: %w", err)
	}

	return routines, nil
}
