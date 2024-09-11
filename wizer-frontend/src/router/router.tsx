import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
import TestLayout from "@/layout/Test";
import SignIn from "@/pages/SignIn";
import Login from "@/pages/Login";
import PrivateRoute from "@/router/PrivateRoute";
import ProfileList from "@/pages/lists/ProfilesList";

const mainRoutes: Array<RouteObject> = [
  { path: "test", element: <PrivateRoute><TestLayout /></PrivateRoute> },
  { path: "profiles", element: <ProfileList /> }
]

const router = createBrowserRouter([
  { path: "/", element: <MainLayout />, children: mainRoutes },
  { path: "signin", element: <SignIn /> },
  { path: "login", element: <Login /> },
])


export default router;
