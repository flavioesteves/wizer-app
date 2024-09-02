import { createBrowserRouter, RouteObject } from "react-router-dom";

import MainLayout from "@/layout/Main";
import TestLayout from "@/layout/Test";
import SignIn from "@/pages/SignIn";
import PrivateRoute from "@/router/PrivateRoute";


const mainRoutes: Array<RouteObject> = [
  { path: "test", element: <PrivateRoute><TestLayout /></PrivateRoute> },
]

const router = createBrowserRouter([
  { path: "/", element: <PrivateRoute><MainLayout /></PrivateRoute>, children: mainRoutes },
  { path: "signin", element: <SignIn /> }
])


export default router;
