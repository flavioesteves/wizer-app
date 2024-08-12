package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	//db.Close()

	err = Migrate(db, "/app/migrations")
	if err != nil {

		fmt.Printf("Migrate: %v\n", err)
		log.Fatal(err)
	}

}

func Migrate(db *sql.DB, migrationDir string) error {
	log.Println("Starting Migration")

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Printf("driver: %v\n", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationDir,
		"postgres",
		driver,
	)

	if err != nil {
		log.Printf("Migrate newDatabase instance %v\n", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Printf("migrate: %v\n", err)
		return err
	}

	log.Println("End Migration")
	return nil
}

func ConnectDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
