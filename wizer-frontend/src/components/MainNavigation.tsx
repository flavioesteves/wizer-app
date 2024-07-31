import { useState } from "react";

import Sidebar from "./Sidebar";

const MainNavigation = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState<boolean>(false);

  const toogleSideBar = () => {
    setIsSidebarOpen((prevState) => {
      return !prevState;
    })
  }

  return (
    <>
      <header className="flex justify-between items-center p-4 border-2 border-b-gray-400">
        <h2 onClick={toogleSideBar}>Icon</h2>
        <img
          className="flex-none"
          src="/wizer1.svg"
          alt="wizer-logo"
          style={{ objectFit: 'cover', maxHeight: '6rem', }}
        />
        <ul className="flex space-x-4">
          <li>
            <a href="#" className="hover:text-gray-400 px-3 py-2 rounded-md"
            >
              Home
            </a>
          </li>
          <li>
            <a href="#" className="hover:text-gray-400 px-3 py-2 rounded-md">
              About
            </a>
          </li>
          <li>
            <a href="#" className="hover:text-gray-400 px-3 py-2 rounded-md">
              Contact
            </a>
          </li>
        </ul>
      </header >
      {isSidebarOpen && <Sidebar isOpen={isSidebarOpen} />
      }
    </>
  )

}


export default MainNavigation;
