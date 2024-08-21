package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

func Update(db *sql.DB, user *pb.User) (*pb.User, error) {

	updatedUser := &pb.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}

	query := `
    UPDATE users
    SET email = $1, password = $2, role = $3, updated_at= $4
    WHERE id = $5
    RETURNING id, email,password, role, created_at, updated_at`

	updatedAt := time.Now().Format(time.RFC3339Nano)

	args := []any{
		user.Email,
		hashedPassword,
		user.Role,
		updatedAt,
		user.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password, &updatedUser.Role, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)

	//BUG

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			//TODO: ERRORS Handle
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}

	return updatedUser, nil
}
