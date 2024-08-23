package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func Update(db *sql.DB, routine *pb.Routine) (*pb.Routine, error) {
	updatedRoutine := &pb.Routine{}

	query := `
    UPDATE routines
    SET profile_id =$1, exercises=$2, updated_by=$3, created_by=$4, updated_at=$5, created_at=$6
    WHERE id = $7
    RETURNING id, profile_id, exercises, updated_by, created_by, updated_at, created_at
  `

	updateAt := time.Now().Format(time.RFC3339Nano)

	args := []any{
		routine.ProfileId,
		routine.Exercises,
		routine.UpdatedBy,
		routine.CreatedBy,
		updateAt,
		routine.CreatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&updatedRoutine.Id,
		&updatedRoutine.ProfileId,
		&updatedRoutine.Exercises,
		&updatedRoutine.UpdatedBy,
		&updatedRoutine.CreatedBy,
		&updatedRoutine.UpdatedAt,
		&updatedRoutine.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return updatedRoutine, nil
}
