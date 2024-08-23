import axios, { AxiosResponse } from "axios";

import User from "@/models/User";
interface ApiResponse {
  users: User[];
}



//const API_URL = import.meta.env.VITE_SERVER_DEV_URI;
const API_URL = "http://0.0.0.0:8080/v1"

async function getUsers() {
  try {
    const response: AxiosResponse<ApiResponse> = await axios.get(API_URL + "/users")
    return response.data.users;
  } catch (error) {
    console.log(error)
    return null
  }
}


export default {
  getUsers: getUsers
}
