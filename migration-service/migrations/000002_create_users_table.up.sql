CREATE TABLE IF NOT EXISTS "users" (
  id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,  

  role INTEGER NOT NULL,
  created_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at timestamp WITH TIME ZONE NOT NULL DEFAULT NOW()
);
