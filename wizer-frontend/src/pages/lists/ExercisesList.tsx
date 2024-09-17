import { useEffect, useState } from "react";
// import { ExerciseColumns } from "@/models/Exercise";
import { Step, Exercise, ExerciseColumns } from "@/models/Exercise";
import api from "@/services/api";
import List from "@/components/ui/list";



const renderExerciseItem = (exercise: {
  id: string;
  name: string;
  muscle_group: string;
  steps: Step[];
  video_url: string;
  video_duration_seconds: number;
}, column: string) => {
  switch (column) {
    case "Id":
      return (<p className="text-gray-700 font-semibold">{exercise.id}</p>);
    case "Name":
      return (<p className="text-blue-500 font-semibold">{exercise.name}</p>);
    case "Muscle Group":
      return (<p className="text-blue-500 font-semibold">{exercise.muscle_group}</p>);
    case "Steps":
      return (<p className="text-blue-500 font-semibold">"PlaceHoder for steps "</p>);
    case "Video":
      return (<p className="text-blue-500 font-semibold">{exercise.video_url}</p>)
    case "Video Duration":
      return (<p className="text-blue-500 font-semibold">{exercise.video_duration_seconds}</p>)
    default:
      return "";
  }
}


const ExercisesList = () => {
  const [exercises, setExercises] = useState<(Exercise | null)[]>([]);
  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response: Exercise[] | null = await api.exercise.getAll();
        if (response !== null) {
          setExercises((prev: any) => ({
            ...prev,
            exercises: response
          }))
        }
      } catch (error) {
        console.error("Error fetching exercises:", error);
      }
    };
    fetchExercises();
  }, []);


  return (
    <>
      <h3 className="text-2xl font-bold mb-4 text-center">Exercises</h3>
      <div className="flex flex-row justify-center">
        <div className="flex justify-center px-4 flex-grow">
          <List
            items={exercises}
            renderItem={renderExerciseItem}
            columns={ExerciseColumns}
            className="border rounded-lg p-2"
          />
        </div>
      </div>
    </>
  )

}



export default ExercisesList;
