package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func connectDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	for i := 0; i < 10; i++ {

		connection, err := sql.Open("pgx", dsn)
		if err != nil {
			log.Printf("Postgres not yet ready (%d attempt): %v\n", i+1, err)
			continue
		}

		err = connection.PingContext(context.Background())
		if err != nil {
			log.Printf("Error pinging database (%d attempt): %v\n", i+1, err)
			continue
		}

		log.Println("Connected to Postgres!")

		return connection, nil
	}
	return nil, errors.New("failed to connect to Postgres after 10 attempts")
}
