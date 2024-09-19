import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
import SignIn from "@/pages/SignIn";
import Login from "@/pages/Login";
import PrivateRoute from "@/router/PrivateRoute";

// Lists 
import ProfileList from "@/pages/lists/ProfilesList";
import ExercisesList from "@/pages/lists/ExercisesList";

// Forms
import ExerciseForm from "@/pages/forms/ExerciseForm";

const mainRoutes: Array<RouteObject> = [
  // Lists Path
  { path: "profiles", element: <PrivateRoute><ProfileList /></PrivateRoute> },
  { path: "exercises", element: <PrivateRoute><ExercisesList /></PrivateRoute> },

  // Dynamic route for editing, viewing by id
  { path: "exercises/:id", element: <PrivateRoute><ExerciseForm /></PrivateRoute> },

  // New records
  { path: "exercises/new", element: <PrivateRoute><ExerciseForm /></PrivateRoute> },
]

const router = createBrowserRouter([
  { path: "/", element: <PrivateRoute><MainLayout /></PrivateRoute>, children: mainRoutes },
  { path: "signin", element: <SignIn /> },
  { path: "login", element: <Login /> },
])


export default router;
