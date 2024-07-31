import axios, { AxiosResponse } from "axios";

import Exercise from "@/models/Exercise";

const API_URL = import.meta.env.VITE_SERVER_DEV_URI;


async function getExercises() {
  try {
    const response: AxiosResponse<Exercise> = await axios.get(API_URL + "/exercises")
    return response.data;
  } catch (error) {
    console.log(error)
    return null
  }
}


export default {
  getExercises: getExercises
}
