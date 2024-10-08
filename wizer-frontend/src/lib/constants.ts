
// Retrieve from environment
const API_URL = "http://localhost:8080/v1" // TODO: this is only for dev
export const CLAIMS_KEY = import.meta.env.VITE_CLAIMS_KEY;

// Constants
export const API_URL_AUTH = API_URL + "/auth";
export const API_URL_USERS = API_URL + "/users";
export const API_URL_EXERCISES = API_URL + "/exercises/";
export const API_URL_PROFILES = API_URL + "/profiles";
export const API_URL_ROUTINES = API_URL + "/routines/";
export const API_URL_USERS_REGISTER = API_URL + "/users/register";
export const SESSION_AUTH_TOKEN = "authToken";

