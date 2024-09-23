export type Step = {
  description: string;
  image_url: string;
}

export type Exercise = {
  id?: string;
  name: string;
  description: string;
  muscle_group: string;
  steps: Step[];
  video_url: string;
  video_duration_seconds: number;
}


export const ColumnKeyMapping: Record<string, string> = {
  //"Id": "id",
  "Name": "name",
  "Muscle Group": "muscle_group",
  "Description": "description",
  "Steps": "steps",
  "Video": "video_url",
  "Video Duration": "video_duration_seconds"
}

export default Exercise;
