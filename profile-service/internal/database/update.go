package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func Update(db *sql.DB, profile *pb.Profile) (*pb.Profile, error) {
	updatedProfile := &pb.Profile{}

	query := `
    UPDATE profilers
    SET user_id = $1, gender =$2, birth_year=$3, height_cm=$4, weight_kg=$5, body_fat_percentage=$6, goal =$7, updated_at=$8
    WHERE id =$
    RETURNING id, user_id, gender, birth_year, height_cm, weight_kg, body_fat_percentage, goal, create_at, updated_at
  `

	updateAt := time.Now().Format(time.RFC3339Nano)

	args := []any{
		profile.UserId,
		profile.Gender,
		profile.BirthYear,
		profile.HeightCm,
		profile.WeightKg,
		profile.BodyFatPercentage,
		profile.Goal,
		updateAt,
		profile.Id,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.QueryRowContext(ctx, query, args...)
	err := row.Scan(
		&updatedProfile.Id,
		&updatedProfile.UserId,
		&updatedProfile.Gender,
		&updatedProfile.BirthYear,
		&updatedProfile.HeightCm,
		&updatedProfile.WeightKg,
		&updatedProfile.BodyFatPercentage,
		&updatedProfile.Goal,
		&updatedProfile.CreatedAt,
		&updatedProfile.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("Edit conflict")
		default:
			return nil, err
		}
	}
	return updatedProfile, nil
}
