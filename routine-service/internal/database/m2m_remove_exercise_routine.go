package database

import (
	"context"
	"database/sql"
	"fmt"
)

func RemoveExerciseFromRoutine(db *sql.DB, routineID string, exerciseIDs []string) error {

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	stmt := `
    DELETE FROM routine_exercises
    WHERE routine_id = $1 AND exercise_id = $2`

	for _, exerciseID := range exerciseIDs {
		result, err := tx.ExecContext(ctx, stmt, routineID, exerciseID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to remove exercise from routine: %w\n", err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get affected rows: %w\n", routineID, exerciseID)
		}

		if rowsAffected == 0 {
			tx.Rollback()
			return fmt.Errorf("no record found for routine_id %s exercise_id %s\n", routineID, exerciseID)
		}
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	fmt.Printf("Successfully removed exercises from routine %s\n", routineID)
	return nil
}
