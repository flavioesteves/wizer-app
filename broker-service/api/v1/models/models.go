package models

type RequestProfile struct {
	ProfileID         string `json:"profile_id"`
	UserID            string `json:"user_id"`
	Gender            string `json:"gender"`
	BirthYear         int32  `json:"birth_year"`
	HeightCm          int32  `json:"height_cm"`
	WeightKg          int32  `json:"weight_kg"`
	BodyFatPercentage string `json:"body_fat_percentage"`
	Goal              string `json:"goal"`
}
