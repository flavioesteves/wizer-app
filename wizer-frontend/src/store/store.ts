import { SESSION_AUTH_TOKEN } from "@/lib/constants"

//zustand

//WARN:   session storage is temporary for dev
function logout() {
  sessionStorage.removeItem(SESSION_AUTH_TOKEN);
}

function getSessionToken(): string | null {
  return sessionStorage.getItem(SESSION_AUTH_TOKEN);
}

function setSessionToken(token: string) {
  sessionStorage.setItem(SESSION_AUTH_TOKEN, token)
}

function deleteSessionToken() {
  sessionStorage.removeItem(SESSION_AUTH_TOKEN)
}

export default {
  logout, getSessionToken, setSessionToken, deleteSessionToken
}
