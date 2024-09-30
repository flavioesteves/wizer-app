package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const DB_TIMEOUT = time.Second * 5

func ConnectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	for i := 0; i < 10; i++ {
		connection, err := sql.Open("pgx", dsn)
		if err != nil {
			log.Printf("Postgres not yed ready (%d attempt): %v\n", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = connection.PingContext(context.Background())
		if err != nil {
			connection.Close()
			log.Printf("Error pinging database (%d attempt): %v\n", i+1, err)
			continue
		}

		log.Println("Connected to Postgres!")
		return connection, nil
	}

	return nil, errors.New("failed to connect to Postgres after 10 attempts")
}
