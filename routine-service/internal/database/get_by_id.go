package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func Get(db *sql.DB, id string) (*pb.Routine, error) {
	query := `
    SELECT id, profile_id, exercises, created_by, updated_by, created_at, updated_at
    FROM routines
    WHERE id = $1`

	routine := &pb.Routine{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, query, id).Scan(
		&routine.Id,
		&routine.ProfileId,
		&routine.Exercises,
		&routine.CreatedBy,
		&routine.UpdatedBy,
		&routine.CreatedAt,
		&routine.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return routine, nil

}
