import React from 'react'
import { FaHome, FaUser, FaCog } from 'react-icons/fa'
import DesktopSidebar from './components/Sidebar'

const DesktopLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="flex h-screen w-screen">
      <div className="bg-green-200 flex flex-col w-[50vw] max-w-[22rem]">
        <DesktopSidebar/>
      </div>
      <div className="bg-blue-200 flex flex-col w-full p-2">
        {children}
      </div>
    </div>
  )
}

export default DesktopLayout