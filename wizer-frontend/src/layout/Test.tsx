import { Outlet } from "react-router-dom";
import Test from "@/pages/Test"

function TestLayout() {
  return (
    <>
      <h1>Test Page Random</h1>
      <main>
        <Outlet />
      </main>
      <Test />
    </>
  )
}

export default TestLayout;
