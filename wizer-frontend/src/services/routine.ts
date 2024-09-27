import axios, { AxiosResponse } from "axios";
import { API_URL_ROUTINES } from "@/lib/constants";
import { Routine } from "@/models/Routine";

type GetAllRoutinesResponse = {
  routines: Routine[];
}

type RoutineResponse = {
  routine: Routine;
}

async function getAll() {
  try {
    const response: AxiosResponse<GetAllRoutinesResponse> = await axios.get(API_URL_ROUTINES);
    return response.data.routines;
  } catch (error) {
    console.error(error)
    return null;
  }
}

async function create(newRoutine: Routine) {
  try {
    const response: AxiosResponse<RoutineResponse> = await axios.post(API_URL_ROUTINES, newRoutine);
    return response.data.routine;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function getRoutineById(id: string) {
  try {
    const response: AxiosResponse<RoutineResponse> = await axios.get(`${API_URL_ROUTINES}/${id}`);
    return response.data.routine;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function updatedRoutine(updatedRoutine: Routine) {
  try {
    const response: AxiosResponse<RoutineResponse> = await axios.put(API_URL_ROUTINES, updatedRoutine);
    return response.data.routine;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function deleteRoutine(id: string) {
  try {
    const response: AxiosResponse<null> = await axios.delete(`${API_URL_ROUTINES}/${id}`)
    return response.data;
  } catch (error) {
    console.error(error);
    return null;
  }
}

export default {
  create,
  getAll,
  getRoutineById,
  updatedRoutine,
}
