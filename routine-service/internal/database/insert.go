package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

const DB_TIMEOUT = time.Second * 5

func Insert(db *sql.DB, routine *pb.Routine) (*pb.Routine, error) {

	newRoutine := &pb.Routine{}

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	// id, profile_id, exercises, created_by, updated_by, created_at, updated_at

	stmt := `
  INSERT INTO routines
  (profile_id ,exercises, created_by, updated_by, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5, $6)
  RETURNING id, profile_id,exercises ,created_by, updated_by, created_at, updated_at`

	createdAt := time.Now().Format(time.RFC3339Nano)
	updatedAt := time.Now().Format(time.RFC3339Nano)

	exercisesEncoded, err := encodeExercise(routine.Exercises)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to bytes the exercises: %v\n", &routine.Exercises)
	}

	row := db.QueryRowContext(ctx, stmt,
		&routine.ProfileId,
		exercisesEncoded,
		&routine.CreatedBy,
		&routine.UpdatedBy,
		createdAt,
		updatedAt,
	)

	var exercisesBytes []byte

	err = row.Scan(
		&newRoutine.Id,
		&newRoutine.ProfileId,
		&exercisesBytes,
		&newRoutine.CreatedBy,
		&newRoutine.UpdatedBy,
		&newRoutine.CreatedAt,
		&newRoutine.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert routine: %w\n", err)
	}

	fmt.Printf("New Routine was created: %v\n", newRoutine)

	exercises, err := decodeExercise(exercisesBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exercises: %w", err)
	}
	newRoutine.Exercises = exercises

	return newRoutine, err
}
