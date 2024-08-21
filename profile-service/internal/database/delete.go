package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Delete(db *sql.DB, id string) error {
	//TODO: Validation of the size of id

	query := `DELETE FROM profiles
    WHERE id=$1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	result, err := db.ExecContext(ctx, query, id)
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
