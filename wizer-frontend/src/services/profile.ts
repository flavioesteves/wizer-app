import axios, { AxiosResponse } from "axios";
import { API_URL_PROFILES } from "@/lib/constants";
import Profile from "@/models/Profile";


type GetAllProfilesResponse = {
  profiles: Profile[];
}

type ProfileResponse = {
  profile: Profile;
}


async function getAll() {
  try {
    const response: AxiosResponse<GetAllProfilesResponse> = await axios.get(API_URL_PROFILES);
    return response.data.profiles;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function create(newProfile: Profile) {
  try {
    const response: AxiosResponse<ProfileResponse> = await axios.post(API_URL_PROFILES, newProfile);
    return response.data.profile;
  } catch (error) {
    console.error(error);
  }
}

async function getProfileById(id: string) {
  try {
    const response: AxiosResponse<ProfileResponse> = await axios.get(`${API_URL_PROFILES}/${id}`);
    return response.data.profile;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function updateProfile(updatedProfile: Profile) {
  try {
    const response: AxiosResponse<ProfileResponse> = await axios.put(API_URL_PROFILES, updatedProfile);
    return response.data.profile;
  } catch (error) {
    console.error(error);
    return null;
  }
}

async function deleteProfile(id: string) {
  try {
    const response: AxiosResponse<null> = await axios.delete(`${API_URL_PROFILES}/${id}`);
    return response.data;
  } catch (error) {
    console.error(error);
    return null;
  }
}

export default {
  create,
  deleteProfile,
  getAll,
  getProfileById,
  updateProfile,
}
