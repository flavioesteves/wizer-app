package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}

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
