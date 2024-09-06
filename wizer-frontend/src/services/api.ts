import axios, { AxiosResponse } from "axios";

import User from "@/models/User";
interface GetAllResponse {
  users: User[];
}


//const API_URL = import.meta.env.VITE_SERVER_DEV_URI;
const API_URL = "http://localhost:8080/v1"

// ####### USERS ######
const API_URL_USERS = API_URL + "/users" //TODO: Create a file with constants

async function getUsers() {
  try {
    const response: AxiosResponse<GetAllResponse> = await axios.get(API_URL_USERS)
    return response.data.users;
  } catch (error) {
    console.log(error)
    return null
  }
}

async function registerUser(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<User> = await axios.post(API_URL_USERS, userReq)
    return response.data
  } catch (error) {
    console.log(error)
    return null
  }
}


// ####### AUTH #######
const API_LOGIN_AUTH = API_URL + "/auth"

async function loginUser(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<String> = await axios.post(API_LOGIN_AUTH, userReq)
    return response.data

  } catch (error) {
    console.log(error)
    return null
  }
}


export default {
  getUsers: getUsers,
  registerUser: registerUser,
  loginUser: loginUser
}
