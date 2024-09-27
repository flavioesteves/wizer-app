import Exercise from "@/models/Exercise";

export type Routine = {
  id: string;
  name: string;
  profile_id: string;
  exercises: Exercise[];
}

export const ColumnKeyMapping: Record<string, string> = {
  "Name": "name",
  "Profile": "profile_id",
  "Exercises": "exercises",
}



