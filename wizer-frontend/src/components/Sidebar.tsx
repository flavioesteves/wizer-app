import { ReactNode } from "react";
import { Link } from "react-router-dom";




type SidebarProps = {
  isOpen?: boolean,
  children?: ReactNode
}


const Sidebar: React.FC<SidebarProps> = ({ isOpen = false }) => {

  const navigationLinks = [
    { path: "/test", label: "Test" },
    { path: "/exercises", label: "Exercises" },
    { path: "/routines", label: "Routines" },
    { path: "/profiles", label: "Profiles" },
  ]


  return (
    <nav className={`fixed top-4rem left-0 w-32 h-screen bg-white shadow-md z-50 transition duration-200 ease-in-out ${isOpen ? '' : '-translate-x-full'}`}>
      <ul className="text-center">
        {navigationLinks.map((link) => (
          <li key={link.path} className="hover:bg-gray-400 py-2 rounded-md" >
            <Link className="hover:underline-offset-auto" to={link.path}>{link.label}</Link>
          </li>
        ))}
      </ul>

    </nav>
  )
}

export default Sidebar;
