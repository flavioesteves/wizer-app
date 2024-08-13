import { Outlet } from "react-router-dom";


function TestLayout() {
  return (
    <>
      <h1>Test Page</h1>
      <main>
        <Outlet />
      </main>
    </>
  )
}

export default TestLayout;
