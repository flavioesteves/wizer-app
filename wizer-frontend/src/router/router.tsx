import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
//import Exercises from "@/pages/Exercises";
import TestLayout from "@/layout/Test";

const mainRoutes: Array<RouteObject> = [
  // { path: "exercises", element: <Exercises /> }
]

const router = createBrowserRouter([
  { path: "/", element: <MainLayout />, children: mainRoutes },
  { path: "/test", element: <TestLayout /> },
])


export default router;
