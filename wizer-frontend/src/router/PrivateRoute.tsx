
import React from "react";
//import React, { useContext } from "react";
import { Navigate } from "react-router-dom";

const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
  // const { isLoggedIn } = useContext()
  const isLoggedIn = false;
  return isLoggedIn ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
