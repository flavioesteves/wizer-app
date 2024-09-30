package database

import (
	"context"
	"database/sql"
	"fmt"
)

func InsertExerciseToRoutine(db *sql.DB, routineID string, exerciseIDs []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	stmt := `
  INSERT INTO routine_exercises (routine_id, exercise_id)
  VALUES ($1, $2)`

	for _, exerciseID := range exerciseIDs {
		_, err := tx.ExecContext(ctx, stmt, routineID, exerciseID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert exercise into routine: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	fmt.Printf("Sucessfully added exercises to routine %s\n", routineID)
	return nil
}
