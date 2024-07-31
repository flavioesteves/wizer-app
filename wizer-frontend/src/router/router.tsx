import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
//import Exercises from "@/pages/Exercises";


const mainRoutes: Array<RouteObject> = [
  // { path: "exercises", element: <Exercises /> }
]

const router = createBrowserRouter([
  { path: "/", element: <MainLayout />, children: mainRoutes }
])


export default router;
