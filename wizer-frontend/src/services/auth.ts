import axios, { AxiosResponse } from "axios";

import { API_URL_AUTH } from "@/lib/constants";
import User from "@/models/User";


async function login(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<string> = await axios.post(API_URL_AUTH, userReq)
    return response.data

  } catch (error) {
    console.log(error)
    return null
  }
}


async function isSessionValid(token: string) {

  try {

    const isValidSessionResponse: AxiosResponse<String> = await axios.post(API_URL_AUTH + "/validate", token)

    console.log(isValidSessionResponse.data)
    return isValidSessionResponse.data;



  } catch (error) {
    console.log(error)
    return false
  }
}


export default {
  login,
  isSessionValid
}
