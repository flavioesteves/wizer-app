CREATE TABLE IF NOT EXISTS "profiles"(
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id VARCHAR(255) NOT NULL,
  gender VARCHAR(255) NOT NULL,
  birth_year INTEGER NOT NULL,
  height_cm INTEGER NOT NULL,
  weight_kg INTEGER NOT NULL,
  body_fat_percentage VARCHAR(255) NOT NULL,
  goal VARCHAR(255) NOT NULL,

  created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
)
