export type Profile = {
  id: string;
  user_id: string;
  gender: string;
  birth_year: number;
  height_cm: number;
  weight_kg: number;
  body_fat_percentage: string;
  goal: string;
}

export const ColumnKeyMapping: Record<string, string> = {
  "User": "user_id",
  "Gender": "gender",
  "Birth Year": "birth_year",
  "Height": "height_cm",
  "Weight": "weight_kg",
  "Body Fat %": "body_fat_percentage",
  "Goal": "goal",
}

export default Profile;
