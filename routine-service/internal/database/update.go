package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func Update(db *sql.DB, routine *pb.Routine) (*pb.Routine, error) {
	updatedRoutine := &pb.Routine{}

	query := `
    UPDATE routines
    SET name=$1 profile_id=$2, updated_by=$3, updated_at=now()
    WHERE id=$5
    RETURNING *`

	args := []any{
		routine.Name,
		routine.ProfileId,
		routine.UpdatedBy,
		routine.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&updatedRoutine.Id,
		&updatedRoutine.ProfileId,
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
