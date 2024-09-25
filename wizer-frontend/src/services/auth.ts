import axios, { AxiosResponse } from "axios";

import { API_URL_AUTH } from "@/lib/constants";
import User from "@/models/User";
import store from "@/store/store";


type Auth = {
  isValid: boolean,
  token: string
}


async function login(user: User) {
  try {
    const userReq: User = {
      email: user.email,
      password: user.password
    }
    const response: AxiosResponse<Auth> = await axios.post(API_URL_AUTH, userReq)

    if (response.data.isValid) {
      store.setSessionToken(response.data.token)
    }
    return response.data

  } catch (error) {
    console.error(error)
    return null
  }
}


async function isSessionValid(token: string) {
  try {
    const isValidSessionResponse: AxiosResponse<String> = await axios.post(API_URL_AUTH + "/validate", token)
    return isValidSessionResponse.data;

  } catch (error) {
    console.error(error)
    return false
  }
}


export default {
  login,
  isSessionValid
}
