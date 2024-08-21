package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	pb "github.com/flavioesteves/wizer-app/profile/proto"
)

func GetAll(db *sql.DB) ([]*pb.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, user_id, gender, birth_year, height_cm, weight_kg, body_fat_percentafe, goal, created_at, updated_at 
    FROM PROFILES ORDER by created_at`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var profiles []*pb.Profile

	for rows.Next() {
		profile := &pb.Profile{}

		err := rows.Scan(
			&profile.Id,
			&profile.UserId,
			&profile.Gender,
			&profile.BirthYear,
			&profile.HeightCm,
			&profile.WeightKg,
			&profile.BodyFatPercentage,
			&profile.Goal,
			&profile.CreatedAt,
			&profile.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning ", err)
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}
