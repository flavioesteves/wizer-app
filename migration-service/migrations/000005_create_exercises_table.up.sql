CREATE TABLE IF NOT EXISTS "exercises" (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255),
  muscle_group VARCHAR(255),
  description TEXT,
  steps JSONB,
  video_url VARCHAR(255),
  video_duration_seconds INT,

  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255) NOT NULL,
  created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
)

