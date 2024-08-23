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

	query := ``

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
			&routine.ProfileId,
			&routine.Exercises,
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
