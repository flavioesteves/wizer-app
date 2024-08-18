package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/flavioesteves/wizer-app/user/proto"
)

const DB_TIMEOUT = time.Second * 3

func Insert(db *sql.DB, user *pb.User) (*pb.User, error) {
	newUser := &pb.User{}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stmt := `INSERT INTO users (email, password, role, created_at, updated_at)
    VALUES ($1, $2, $3, $4, $5) RETURNING id, email, password, role, created_at, updated_at`

	createdAt := time.Now().Format(time.RFC3339Nano)
	updatedAt := time.Now().Format(time.RFC3339Nano)

	row := db.QueryRowContext(ctx, stmt,
		&user.Email,
		&hashedPassword,
		&user.Role,
		createdAt,
		updatedAt,
	)

	err = row.Scan(&newUser.Id, &newUser.Email, &newUser.Password, &newUser.Role, &newUser.CreatedAt, &newUser.UpdatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w\n", err)
	}

	fmt.Printf("New user created: %+v\n", newUser) // Add logging for debugging

	return newUser, err
}


