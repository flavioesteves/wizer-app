package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

func Delete(db *sql.DB, id string) error {
	query := `DELETE FROM exercises
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
		return fmt.Errorf("Record not found")
	}
	return nil
}
