import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
//import Exercises from "@/pages/Exercises";
import TestLayout from "@/layout/Test";

const mainRoutes: Array<RouteObject> = [
  // { path: "exercises", element: <Exercises /> }
  { path: "test", element: <TestLayout /> }
]

const router = createBrowserRouter([
  { path: "/", element: <MainLayout />, children: mainRoutes },
])


export default router;
