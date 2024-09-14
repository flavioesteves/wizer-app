
import React from "react";
import { Navigate } from "react-router-dom";
import store from "@/store/store";

const PrivateRoute = ({ children }: { children: React.ReactNode }) => {
  const isLoggedIn = store.getSessionToken()
  return isLoggedIn ? children : <Navigate to="/login" />;
};

export default PrivateRoute;
