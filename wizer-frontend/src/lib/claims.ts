import * as jose from "jose";


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

export async function verifyToken(token: string, secretKey: string | undefined): Promise<Claims | null> {

  try {
    const convertedSecret = new TextEncoder().encode(secretKey);
    const { payload } = await jose.jwtVerify(token, convertedSecret);

    return payload as Claims;
  } catch (error) {
    return null;
  }
}

