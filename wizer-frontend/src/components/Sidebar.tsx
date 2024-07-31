import { ReactNode } from "react";



interface SidebarProps {
  isOpen?: boolean,
  children?: ReactNode
}


const Sidebar: React.FC<SidebarProps> = ({ isOpen = false }) => {
  return (
    <nav className={`fixed top-4rem left-0 w-64 h-screen bg-white shadow-md z-50 transition duration-200 ease-in-out ${isOpen ? '' : '-translate-x-full'}`}>
      <ul>
        <li>Exercises</li>
        <li>Routines</li>
        <li>Tests</li>
      </ul>

    </nav>
  )
}

export default Sidebar;
