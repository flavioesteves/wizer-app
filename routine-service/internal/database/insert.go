package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
	"github.com/lib/pq"
)

const DB_TIMEOUT = time.Second * 5

func Insert(db *sql.DB, routine *pb.Routine) (*pb.Routine, error) {

	newRoutine := &pb.Routine{}

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	stmt := `
  INSERT INTO routines
  (profile_id ,exercises, created_by, updated_by, created_at, updated_at)
  VALUES ($1, $2, $3, $4, now(), now())
  RETURNING id, profile_id, exercises ,created_by, updated_by, created_at, updated_at`

	row := db.QueryRowContext(ctx, stmt,
		&routine.ProfileId,
		&routine.Exercises,
		&routine.CreatedBy,
		&routine.UpdatedBy,
	)

	err := row.Scan(
		&newRoutine.Id,
		&newRoutine.ProfileId,
		pq.Array(&newRoutine.Exercises),
		&newRoutine.CreatedBy,
		&newRoutine.UpdatedBy,
		&newRoutine.CreatedAt,
		&newRoutine.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert routine: %w\n", err)
	}

	fmt.Printf("New Routine was created: %v\n", newRoutine)

	return newRoutine, err
}
