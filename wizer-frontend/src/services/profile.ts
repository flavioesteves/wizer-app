import axios, { AxiosResponse } from "axios";
import { API_URL_PROFILES } from "@/lib/constants";
import Profile from "@/models/Profile";



type GetALlProfilesResponse = {
  profiles: Profile[];
}

type GetProfileResponse = {
  profile: Profile;
}


async function getAll() {
  try {
    const response: AxiosResponse<GetALlProfilesResponse> = await axios.get(API_URL_PROFILES);
    return response.data.profiles;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function getProfileById(id: string) {
  try {
    const response: AxiosResponse<GetProfileResponse> = await axios.get(`${API_URL_PROFILES}/${id}`)
    return response.data.profile;
  } catch (error) {
    console.error(error);
  }
}

export default {
  getAll,
  getProfileById,
}
