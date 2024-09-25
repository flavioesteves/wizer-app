import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
import SignIn from "@/pages/SignIn";
import Login from "@/pages/Login";
import PrivateRoute from "@/router/PrivateRoute";

// Lists 
import ProfileList from "@/pages/lists/ProfilesList";
import ExercisesList from "@/pages/lists/ExercisesList";
import RoutinesList from "@/pages/lists/RoutinesList";

// Forms
import ExerciseForm from "@/pages/forms/ExerciseForm";
import RoutineForm from "@/pages/forms/RoutineForm";


const mainRoutes: Array<RouteObject> = [
  // Lists Path
  { path: "profiles", element: <PrivateRoute><ProfileList /></PrivateRoute> },

  // Exercise
  { path: "exercises", element: <PrivateRoute><ExercisesList /></PrivateRoute> },
  { path: "exercises/:id", element: <PrivateRoute><ExerciseForm /></PrivateRoute> },
  { path: "exercises/new", element: <PrivateRoute><ExerciseForm /></PrivateRoute> },
  // Routines
  { path: "routines", element: <PrivateRoute><RoutinesList /></PrivateRoute> },
  { path: "routines/:id", element: <PrivateRoute><RoutineForm /></PrivateRoute> },
  { path: "routines/new", element: <PrivateRoute><RoutineForm /></PrivateRoute> },
]

const router = createBrowserRouter([
  { path: "/", element: <PrivateRoute><MainLayout /></PrivateRoute>, children: mainRoutes },
  { path: "signin", element: <SignIn /> },
  { path: "login", element: <Login /> },
])


export default router;
