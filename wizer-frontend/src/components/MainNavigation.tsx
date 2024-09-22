import { useState } from "react";
import { Link } from "react-router-dom";

import Sidebar from "./Sidebar";
import Button from "./ui/button";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars } from "@fortawesome/free-solid-svg-icons";


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
        <Button variant="secondary" onClick={toogleSideBar} >
          <FontAwesomeIcon icon={faBars} size="lg" className="fa-fw" />
        </Button>
        <img
          className="flex-none"
          src="/wizer1.svg"
          alt="wizer-logo"
          style={{ objectFit: 'cover', maxHeight: '6rem', }}
        />
        <ul className="flex space-x-4">
          <li>
            <Link to={"/"}>Home</Link>
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
