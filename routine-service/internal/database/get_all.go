package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

func GetAll(db *sql.DB) ([]*pb.Routine, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, name ,profile_id, created_by, created_at, updated_by, updated_at
    FROM routines ORDER by created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var routines []*pb.Routine

	for rows.Next() {
		routine := &pb.Routine{}

		err := rows.Scan(
			&routine.Id,
			&routine.Name,
			&routine.ProfileId,
			&routine.CreatedAt,
			&routine.UpdatedAt,
			&routine.CreatedBy,
			&routine.UpdatedBy,
		)
		if err != nil {
			log.Println("error scanning", err)
			return nil, err
		}
		routines = append(routines, routine)
	}
	return routines, nil
}
