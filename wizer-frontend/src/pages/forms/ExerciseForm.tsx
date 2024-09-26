import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { Exercise } from "@/models/Exercise";
import Button from "@/components/ui/button";
import api from "@/services/api";

const ExerciseForm: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const isNew = !id;


  const [exercise, setExercise] = useState<Exercise>({
    id: id ?? "",
    name: "",
    description: "",
    muscle_group: "",
    steps: [{ description: "", image_url: "" }],
    video_url: "",
    video_duration_seconds: 0,
  });

  const handleChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setExercise((prev) => ({
      ...prev,
      [event.target.name]: event.target.value,
    }))
  }

  const handleStepChange = (index: number, event: React.ChangeEvent<HTMLInputElement>) => {
    const newSteps = [...exercise.steps];
    newSteps[index] = {
      ...newSteps[index],
      [event.target.name]: event.target.value,
    }
    setExercise((prev) => ({
      ...prev,
      ...exercise, steps: newSteps
    }))
  };

  const addStep = () => {
    setExercise((prev) => ({
      ...prev,
      steps: [...prev.steps, { description: "", image_url: "" }]
    }));
  }

  const removeStep = (index: number) => {
    const newSteps = exercise.steps.filter((_, i) => i !== index);
    setExercise((prev) => ({
      ...prev,
      steps: newSteps
    }));
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()

    const newExercise: Exercise = {
      name: exercise.name,
      description: exercise.description,
      muscle_group: exercise.muscle_group,
      steps: exercise.steps,
      video_url: exercise.video_url,
      video_duration_seconds: parseInt(exercise.video_duration_seconds.toString(), 10),
    }
    // api call
    try {
      let res = await api.exercise.create(newExercise);
      if (res) {
        navigate("/exercises")
      }
    } catch (error) {
      console.error(error)
    }
  }

  const handleUpdate = async () => {

    const updatedExercise: Exercise = {
      id: id,
      name: exercise.name,
      description: exercise.description,
      muscle_group: exercise.muscle_group,
      steps: exercise.steps,
      video_url: exercise.video_url,
      video_duration_seconds: parseInt(exercise.video_duration_seconds.toString(), 10),
    }

    try {
      let res = await api.exercise.updateExercise(updatedExercise);
      if (res) {
        navigate("/exercises")
      }
    } catch (error) {
      console.error(error)
    }
  }



  useEffect(() => {
    const fetchExercise = async () => {
      if (id) {

        try {
          const response: Exercise | null = await api.exercise.getExerciseById(id);
          if (response) {
            setExercise({
              ...response,
              steps: response.steps ?? [{ description: "", image_url: "" }]
            });
          }
        } catch (error) { console.error("Error fetching exercise:", error) }

      }
    };
    fetchExercise();
  }, [id]);

  return (
    <>
      <div className="max-w-2xl mx-auto mt-10">
        <h1>{isNew ? "Create New Exercise" : `Edit Exercise ${id}`}</h1>
        <form onSubmit={handleSubmit} className="space-y-4">

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Name</label>
            <input
              type="text"
              name="name"
              value={exercise.name}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Description</label>
            <input
              type="text"
              name="description"
              value={exercise.description}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Muscle Group</label>
            <input
              type="text"
              name="muscle_group"
              value={exercise.muscle_group}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-gray-700">Steps</label>
            {exercise.steps.map((step, index) => (
              <div key={index} className="mb-4">
                <div className="flex space-x-2">
                  <input
                    type="text"
                    name="description"
                    placeholder="Step description"
                    value={step.description}
                    onChange={(e) => handleStepChange(index, e)}
                    className="p-2 block w-full border border-gray-300 rounded-md"
                  />
                  <input
                    type="text"
                    name="image_url"
                    placeholder="Step image URL"
                    value={step.image_url}
                    onChange={(e) => handleStepChange(index, e)}
                    className="p-2 block w-full border border-gray-300 rounded-md"
                  />
                  {exercise.steps.length > 1 && (<Button
                    type="button"
                    variant="red"
                    onClick={() => removeStep(index)}
                  >
                    Remove
                  </Button>)}
                </div>
              </div>
            ))}
            <Button
              type="button"
              variant="green"
              onClick={addStep}
            >
              Add Step
            </Button>
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Video URL</label>
            <input
              type="text"
              name="video_url"
              value={exercise.video_url}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Video Duration</label>
            <input
              type="number"
              name="video_duration_seconds"
              value={exercise.video_duration_seconds}
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>
          {isNew ?
            <Button type="submit">Create Exercise</Button> :
            <Button type="button" onClick={handleUpdate}>Update Exercise</Button>
          }
        </form>
      </div>


    </>
  )
}


export default ExerciseForm;
