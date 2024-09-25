import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";

import Sidebar from "./Sidebar";
import Button from "./ui/button";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars, faCircleChevronLeft } from "@fortawesome/free-solid-svg-icons";


const MainNavigation = () => {
  const navigate = useNavigate();

  const [isSidebarOpen, setIsSidebarOpen] = useState<boolean>(false);


  const toogleSideBar = () => {
    setIsSidebarOpen((prevState) => {
      return !prevState;
    })
  }

  const handleBack = () => {
    navigate(-1);
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
          <Button variant="secondary" onClick={handleBack}>
            <FontAwesomeIcon icon={faCircleChevronLeft} size="lg" className="fa=fw" />
          </Button>
        </ul>
      </header >
      {isSidebarOpen && <Sidebar isOpen={isSidebarOpen} />
      }
    </>
  )

}


export default MainNavigation;
