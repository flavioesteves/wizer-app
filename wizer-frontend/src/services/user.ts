import axios, { AxiosResponse } from "axios";

import { API_URL_USERS, API_URL_USERS_REGISTER } from "@/lib/constants";
import User from "@/models/User";


type GetAllUsersResponse = {
  users: User[];
}

async function getAll() {
  try {
    const response: AxiosResponse<GetAllUsersResponse> = await axios.get(API_URL_USERS)
    return response.data.users;
  } catch (error) {
    console.error(error)
    return null
  }
}

async function register(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<User> = await axios.post(API_URL_USERS_REGISTER, userReq)
    return response.data
  } catch (error) {
    console.error(error)
    return null
  }
}

export default {
  getAll: getAll,
  register: register,
}
