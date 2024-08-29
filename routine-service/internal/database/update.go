package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
	"github.com/lib/pq"
)

func Update(db *sql.DB, routine *pb.Routine) (*pb.Routine, error) {
	updatedRoutine := &pb.Routine{}

	query := `
    UPDATE routines
    SET profile_id=$1, exercises=$2, updated_by=$3, updated_at=now()
    WHERE id=$4
    RETURNING id, profile_id, exercises, updated_by, created_by, updated_at, created_at
  `

	args := []any{
		routine.ProfileId,
		routine.Exercises,
		routine.UpdatedBy,
		routine.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&updatedRoutine.Id,
		&updatedRoutine.ProfileId,
		pq.Array(&updatedRoutine.Exercises),
		&updatedRoutine.UpdatedBy,
		&updatedRoutine.CreatedBy,
		&updatedRoutine.UpdatedAt,
		&updatedRoutine.CreatedAt,
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)

		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return updatedRoutine, nil
}
