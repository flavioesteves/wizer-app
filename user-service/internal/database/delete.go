package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Delete(db *sql.DB, userId string) error {

	//TODO: Validation of the size og id

	query := `
    DELETE FROM users
    WHERE id =$1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		//TODO: Implement a correct error log
		return fmt.Errorf("Record not found")
	}

	return nil
}
