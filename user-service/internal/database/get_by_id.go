package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func Get(db *sql.DB, id string) (*pb.User, error) {

	query := `
    SELECT email, password, role, created_atm updated_at
    FROM users
    WHERE id = $1
  `

	user := &pb.User{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, query).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return user, nil
}
