CREATE TABLE IF NOT EXISTS "routines" (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  profile_id uuid NOT NULL REFERENCES profiles(id),
  created_by uuid NOT NULL REFERENCES users(id),
  updated_by uuid NOT NULL REFERENCES users(id),
  created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
)
