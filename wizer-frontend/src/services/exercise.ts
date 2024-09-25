import axios, { AxiosResponse } from "axios";
import { API_URL_EXERCISES } from "@/lib/constants";

import Exercise from "@/models/Exercise";

type GetAllExercisesResponse = {
  exercises: Exercise[];
}

async function getAll() {
  try {
    const response: AxiosResponse<GetAllExercisesResponse> = await axios.get(API_URL_EXERCISES)
    return response.data.exercises;

  } catch (error) {
    console.error(error)
    return null;
  }
}

async function create(newExercise: Exercise) {
  try {
    const response: AxiosResponse<Exercise> = await axios.post(API_URL_EXERCISES, newExercise)
    return response.data
  } catch (error) {
    console.error(error)
    return null;
  }
}

async function getExerciseById(id: string) {
  try {
    const response: AxiosResponse<Exercise> = await axios.get(API_URL_EXERCISES + "/" + id);
    return response.data;
  } catch (error) {
    console.error(error)
  }
}

async function updateExercise(updatedExercise: Exercise) {
  try {
    const response: AxiosResponse<Exercise> = await axios.put(API_URL_EXERCISES, updatedExercise);
    return response.data;
  } catch (error) {
    console.error(error);
  }
}

async function deleteExercise(id: string) {
  try {
    const response: AxiosResponse<null> = await axios.delete(API_URL_EXERCISES + "/" + id);
    return response.data;
  } catch (error) {
    console.error(error)
  }
}

export default {
  getAll,
  create,
  updateExercise,
  getExerciseById,
  deleteExercise,
}
