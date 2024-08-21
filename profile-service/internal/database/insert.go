package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

const DB_TIMEOUT = time.Second * 3

func Insert(db *sql.DB, profile *pb.Profile) (*pb.Profile, error) {
	newProfile := &pb.Profile{}

	ctx, cancel := context.WithTimeout(context.Background(), DB_TIMEOUT)
	defer cancel()

	stmt := `
  INSERT INTO profiles
  (user_id, gender, birth_year, height_cm, weight_kg, body_fat_percentage, goal, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
  RETURNING id, user_id, gender, birth_year, height_cm, weight_kg, body_fat_percentage, goal, created_at, updated_at`

	createdAt := time.Now().Format(time.RFC3339Nano)
	updatedAt := time.Now().Format(time.RFC3339Nano)

	row := db.QueryRowContext(ctx, stmt,
		&profile.UserId,
		&profile.Gender,
		&profile.BirthYear,
		&profile.HeightCm,
		&profile.WeightKg,
		&profile.BodyFatPercentage,
		&profile.Goal,
		createdAt,
		updatedAt,
	)

	err := row.Scan(
		&newProfile.Id,
		&newProfile.UserId,
		&newProfile.Gender,
		&newProfile.BirthYear,
		&newProfile.HeightCm,
		&newProfile.WeightKg,
		&newProfile.BodyFatPercentage,
		&newProfile.Goal,
		&newProfile.CreatedAt,
		&newProfile.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert profile: %w\n", err)
	}

	fmt.Printf("New profile was created: %v\n", newProfile)

	return newProfile, err
}
