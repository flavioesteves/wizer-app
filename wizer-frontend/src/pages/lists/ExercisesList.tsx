import { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faPlus } from "@fortawesome/free-solid-svg-icons";

import api from "@/services/api";
import { Exercise, ColumnKeyMapping } from "@/models/Exercise";
import List from "@/components/ui/list";
import Button from "@/components/ui/button";



const renderExerciseItem = (exercise: Exercise, key: string) => {
  switch (key) {
    case "name":
      return (<p className="text-blue-500 font-semibold">{exercise.name}</p>);
    case "muscle_group":
      return (<p className="text-blue-500 font-semibold">{exercise.muscle_group}</p>);
    case "description":
      return (<p className="text-blue-500 font-semibold">{exercise.description}</p>);
    case "steps":
      return (
        <div className="flex space-x-2 text-blue-500 font-semibold">
          {exercise.steps.map((step, index) => (
            <div key={index} className="flex-1">
              <p>{step.description}</p>
            </div>
          ))}
        </div>
      );
    case "video_url":
      return (<p className="text-blue-500 font-semibold">{exercise.video_url}</p>)
    case "video_duration_seconds":
      return (<p className="text-blue-500 font-semibold">{exercise.video_duration_seconds}</p>)
    default:
      return "";
  }
}




const ExercisesList = () => {
  const navigate = useNavigate();
  const excludedKeys: (keyof Exercise)[] = ["id"];
  const col = Object.keys(ColumnKeyMapping);
  const [exercises, setExercises] = useState<Exercise[]>([]);


  useEffect(() => {
    const fetchExercises = async () => {
      try {
        const response: Exercise[] | null = await api.exercise.getAll();
        if (response) {
          setExercises(response)
        }
      } catch (error) {
        console.error("Error fetching exercises:", error);
      }
    };
    fetchExercises();
  }, []);


  const handleEdit = (id: string) => {
    navigate(`/exercises/${id}`)
  }

  const handleDelete = (id: string) => {
    //TODO: Create modal components
    if (window.confirm("Are you sure you want to delete this exercise?")) {
      api.exercise.deleteExercise(id).then(() => {
        setExercises((prev) => prev.filter((exercise) => exercise.id !== id));
      }).catch((error) => console.error("Error deleting exercice:", error));
    }
  }


  const actions = [
    {
      label: "Edit",
      onClick: handleEdit,
      element: <Button variant="primary">Edit</Button>
    },
    {
      label: "Delete",
      onClick: handleDelete,
      element: <Button variant="red">Delete</Button>
    }
  ];


  return (
    <>
      <h3 className="text-2xl font-bold mb-4 text-center">Exercises</h3>
      <div className="flex flex-row justify-center overflow-x-auto">
        <div className="flex justify-center px-4 flex-grow">
          <List
            excludeKeys={excludedKeys}
            items={exercises}
            renderItem={renderExerciseItem}
            actions={actions}
            columnsHeaders={col}
            columnKeyMapping={ColumnKeyMapping}
            className="border rounded-lg p-2">
          </List>
        </div>
        <Link to="/exercises/new">
          <Button variant="green" className="rounded-full">
            <FontAwesomeIcon icon={faPlus} size="lg" className="fa-fw rounded-full" />
            New
          </Button>
        </Link>
      </div >
    </>
  )

}



export default ExercisesList;
