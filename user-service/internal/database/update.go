package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func Update(db *sql.DB, user *pb.User) (*pb.User, error) {

	updatedUser := &pb.User{}

	query := `
    UPDATE users
    SET email = $1, password = $2, role = $3, updated_at= $4
    RETURNING id, email, password, role, created_at, updated_at`

	updatedAt := time.Now().Format(time.RFC3339Nano)

	args := []any{
		user.Email,
		user.Password,
		user.Role,
		updatedAt,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, query, args...).Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password, &updatedUser.Role, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			//TODO: ERRORS Handle
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}

	return user, nil
}
