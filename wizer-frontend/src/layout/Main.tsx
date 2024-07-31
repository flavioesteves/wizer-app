import { Outlet } from "react-router-dom";

import MainNavigation from "@/components/MainNavigation";
function MainLayout() {
  return (
    <>
      <MainNavigation />
      <h1>Main Layout</h1>
      <main>
        <Outlet />
      </main>
    </>
  )
}

export default MainLayout;
