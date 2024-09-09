import axios, { AxiosResponse } from "axios";

import User from "@/models/User";

const API_URL_AUTH = "http://localhost:8080/v1/auth"


async function signIn(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<String> = await axios.post(API_URL_AUTH, userReq)
    return response.data

  } catch (error) {
    console.log(error)
    return null
  }
}


async function validateSession(token: String) {

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
  signIn: signIn,
  validateSession: validateSession
}
