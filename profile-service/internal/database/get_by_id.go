package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func Get(db *sql.DB, id string) (*pb.Profile, error) {
	query := `
    SELECT id, user_id,gender ,birth_year, height_cm, weight_kg, body_fat_percentage, goal, created_at, updated_at
    FROM profiles
    WHERE id = $1`

	profile := &pb.Profile{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := db.QueryRowContext(ctx, query, id).Scan(
		&profile.Id,
		&profile.UserId,
		&profile.BirthYear,
		&profile.HeightCm,
		&profile.WeightKg,
		&profile.BodyFatPercentage,
		&profile.Goal,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return profile, nil
}
