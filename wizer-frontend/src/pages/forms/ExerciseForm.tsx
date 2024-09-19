import React, { useState } from "react";
import { useParams } from "react-router-dom";
import { Exercise } from "@/models/Exercise";
import Button from "@/components/ui/button";
import api from "@/services/api";

const ExerciseForm: React.FC = () => {
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
      muscle_group: exercise.muscle_group,
      description: exercise.description,
      steps: exercise.steps,
      video_url: exercise.video_url,
      video_duration_seconds: exercise.video_duration_seconds,
    }

    // api call

    let res = await api.exercise.create(newExercise);
    if (res) {
      console.log("Response: ", res)
    }
    console.log("Exercise submitted:", exercise);
  }

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
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block texwwwt-gray-700 font-bold mb-2">Descritpion</label>
            <input
              type="text"
              name="description"
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
                  <Button
                    type="button"
                    variant="red"
                    onClick={() => removeStep(index)}
                  >
                    Remove
                  </Button>
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
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-bold mb-2">Video Duration</label>
            <input
              type="number"
              name="vide_duration"
              onChange={handleChange}
              className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
              required
            />
          </div>

          <Button type="submit">
            {isNew ? "Create Exercise" : "Update Exercise"}
          </Button>
        </form>
      </div>


    </>
  )
}


export default ExerciseForm;
