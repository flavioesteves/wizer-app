package models

type RequestProfile struct {
	Id                string `json:"id"`
	UserID            string `json:"user_id"`
	Gender            string `json:"gender"`
	BirthYear         int32  `json:"birth_year"`
	HeightCm          int32  `json:"height_cm"`
	WeightKg          int32  `json:"weight_kg"`
	BodyFatPercentage string `json:"body_fat_percentage"`
	Goal              string `json:"goal"`
}

type RequestUser struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
