import { SESSION_AUTH_TOKEN } from "@/lib/constants";
import * as jose from "jose";
//zustand
//const CLAIMS_KEY = process.env.CLAIMS_KEY;
//const CLAIMS_KEY = "something"; //TEST : DEV

//WARN:   session storage is temporary for dev
function logout() {
  sessionStorage.removeItem(SESSION_AUTH_TOKEN);
}

function getSessionToken(): string | null {
  return sessionStorage.getItem(SESSION_AUTH_TOKEN);
}

function setSessionToken(token: string) {

  sessionStorage.setItem(SESSION_AUTH_TOKEN, token);
  (async () => {
    const claims = await verifyToken(token, "something");
    if (claims) {

      console.log("Decoded claims:", claims)
      console.log("Username: ", claims.Username)
    }
  })();
}

function deleteSessionToken() {
  sessionStorage.removeItem(SESSION_AUTH_TOKEN)
}

type StandardClaims = {
  iss?: string;  // Issuer
  sub?: string;  // Subject
  aud?: string;  // Audience
  exp?: number;  // Expiration time (as Unix timestamp)
  nbf?: number;  // Not before (as Unix timestamp)
  iat?: number;  // Issued at (as Unix timestamp)
  jti?: string;  // JWT ID (unique identifier)
};


type Claims = StandardClaims & {
  Username: string;
}


async function verifyToken(token: string, secretKey: string): Promise<Claims | null> {
  try {
    const convertedSecret = new TextEncoder().encode(secretKey);
    const { payload } = await jose.jwtVerify(token, convertedSecret);

    console.log(payload)
    return payload as Claims;
  } catch (error) {
    return null;
  }
}

export default {
  logout, getSessionToken, setSessionToken, deleteSessionToken
}
