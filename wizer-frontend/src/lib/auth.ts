import * as jwt from "@/lib/claims";
import { SESSION_AUTH_TOKEN, CLAIMS_KEY } from "@/lib/constants";

export async function isTokenValid(): Promise<boolean> {
  const token = sessionStorage.getItem(SESSION_AUTH_TOKEN);

  if (token) {
    const claims = await jwt.verifyToken(token, CLAIMS_KEY);

    if (claims && claims.exp) {
      const currentTime = Math.floor(Date.now() / 1000); // seconds

      if (claims.exp > currentTime) {
        return true;
      } else {
        sessionStorage.removeItem(SESSION_AUTH_TOKEN);
        console.warn("Session token has expired.");
        return false
      }
    }
  }
  return false;
}
