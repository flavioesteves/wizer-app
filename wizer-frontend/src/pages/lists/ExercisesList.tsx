import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Step, Exercise, ExerciseColumns } from "@/models/Exercise";
import api from "@/services/api";
import List from "@/components/ui/list";
import Button from "@/components/ui/button";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";



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
      return (
        <div className="text-blue-500 font-semibold">
          {exercise.steps.map((step, index) => (
            <div key={index}>
              <p>{step.description}</p>
            </div>
          ))}
        </div>
      );
    case "Video":
      return (<p className="text-blue-500 font-semibold">{exercise.video_url}</p>)
    case "Video Duration":
      return (<p className="text-blue-500 font-semibold">{exercise.video_duration_seconds}</p>)
    default:
      return "";
  }
}


const ExercisesList = () => {
  const col = ExerciseColumns;
  const [exercises, setExercises] = useState<Exercise[]>([]);
  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response: Exercise[] | null = await api.exercise.getAll();
        console.log(response)
        if (response) {
          setExercises(response)
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
      <div className="flex flex-row justify-center overflow-x-auto">
        <div className="flex justify-center px-4 flex-grow">
          <List
            items={exercises}
            renderItem={renderExerciseItem}
            columns={col}
            className="border rounded-lg p-2"
          />
        </div>
        <Link to="/exercises/new">
          <Button variant="green" className="rounded-full">
            <FontAwesomeIcon icon={faPlus} size="lg" className="fa-fw rounded-full" />
          </Button>
        </Link>
      </div >
    </>
  )

}



export default ExercisesList;
