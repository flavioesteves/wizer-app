export type Step = {
  description: string;
  image_url: string;
}

export type Exercise = {
  id: string;
  name: string;
  muscle_group: string;
  steps: Step[];
  video_url: string;
  video_duration_seconds: number;
}
export const ExerciseColumns = ["Id", "Name", "Muscle Group", "Steps", "Video", "Video Duration"];


export default Exercise;
