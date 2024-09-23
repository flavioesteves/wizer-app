
import React, { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import * as jwt from "@/lib/auth";
const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
  const [isLoggedIn, setIsLoggedIn] = useState<boolean | null>(null);

  useEffect(() => {
    const checkTokenValidity = async () => {
      const valid = await jwt.isTokenValid();
      setIsLoggedIn(valid);
    };

    checkTokenValidity();
  }, []);

  //TODO: enter some animation to load
  if (isLoggedIn === null) {

    return <div>Loading ...</div>;
  }

  return isLoggedIn ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
