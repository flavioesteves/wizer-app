CREATE TABLE IF NOT EXISTS "routines" (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  profile_id VARCHAR(255) NOT NULL,
  exercises  JSONB NOT NULL,
  created_by VARCHAR(255) NOT NULL,
  updated_by VARCHAR(255) NOT NULL,
  created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
)