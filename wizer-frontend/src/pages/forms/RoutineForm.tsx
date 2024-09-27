import { Routine } from "@/models/Routine";
import React, { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import api from "@/services/api";
import Exercise from "@/models/Exercise";
import Button from "@/components/ui/button";

const RoutineForm: React.FC = () => {
  const navigate = useNavigate();
  const { id } = useParams<{ id: string }>();
  const isNew = !id;
  const exercisePlaceholder: Exercise = {
    id: id ?? "",
    name: "",
    description: "",
    muscle_group: "",
    steps: [{ description: "", image_url: "" }],
    video_url: "",
    video_duration_seconds: 0,
  }

  const [routine, setRoutine] = useState<Routine>({
    id: id ?? "",
    profile_id: "",
    name: "",
    // This is placeholder the relationship in the backend needed to be implemented
    exercises: [exercisePlaceholder],
  });

  const handleChange = (event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setRoutine((prev) => ({
      ...prev,
      [event.target.name]: event.target.value,
    }))
  }

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault()

    const newRoutine: Routine = {
      name: routine.name,
      profile_id: routine.profile_id,
      exercises: routine.exercises,
    }
    // api
    try {
      let response = await api.routine.create(newRoutine);
      if (response) {
        navigate("/routines")
      }
    } catch (error) {
      console.error(error);
    }
  }

  useEffect(() => {
    const fetchRoutine = async () => {
      if (id) {
        try {
          const response: Routine | null = await api.routine.getRoutineById(id);
          if (response) {
            setRoutine({
              ...response,
              exercises: response.exercises ?? [exercisePlaceholder]
            });
          }
        } catch (error) {
          console.error("Error fetching routine:", error);
        }
      }
    };
    fetchRoutine();
  }, [id]);

  return (
    <div className="max-w-2xl mx-auto mt-10">
      <h1>{isNew ? "Create New Routine" : `Edit Routine ${id}`}</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div className="mb-4">
          <label className="block text-gray-700 font-bold mb-2">Name</label>
          <input
            type="text"
            name="name"
            value={routine.name}
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            required
          />
        </div>

        <div className="mb-4">
          <label className="block text-gray-700 font-bold mb-2">Profile</label>
          <input
            type="text"
            name="profile_id"
            value={routine.profile_id}
            onChange={handleChange}
            className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:border-blue-500"
            required
          />
        </div>
        {isNew ?
          <Button type="submit">Create Routine</Button> :
          <Button type="button">Update Routine</Button>
        }
      </form>
    </div>
  )
}

export default RoutineForm;
