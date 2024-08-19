package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func GetAll(db *sql.DB) ([]*pb.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, email, password, role, created_at, updated_at
    FROM USERS ORDER by created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*pb.User

	for rows.Next() {
		user := &pb.User{}

		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning ", err)
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil

}
