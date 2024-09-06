package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type Credentials struct {
	Email    string
	Password string
}

func GetByEmail(db *sql.DB, email string) (*Credentials, error) {

	credentials := &Credentials{}

	query := `
    SELECT email, password
    FROM users
    WHERE email = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, query, email).Scan(
		&credentials.Email,
		&credentials.Password,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Not found account with that email")
		default:
			return nil, err
		}

	}

	return credentials, nil

}
