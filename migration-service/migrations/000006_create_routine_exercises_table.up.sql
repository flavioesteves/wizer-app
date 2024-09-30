CREATE TABLE IF NOT EXISTS "routine_exercises" (
  routine_id uuid NOT NULL REFERENCES routines(id) ON DELETE CASCADE,
  exercise_id uuid NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
  PRIMARY KEY (routine_id, exercise_id)
);
