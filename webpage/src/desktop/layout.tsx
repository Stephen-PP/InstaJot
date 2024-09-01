import React from 'react'
import { FaHome, FaUser, FaCog } from 'react-icons/fa'

const Sidebar = () => {
  return (
    <div className="w-1/4 h-screen sticky top-0 bg-gray-200">
      <div className="flex flex-col items-center justify-center h-full">
        <div className="mb-4">
          <FaHome size={24} />
        </div>
        <div className="mb-4">
          <FaUser size={24} />
        </div>
        <div className="mb-4">
          <FaCog size={24} />
        </div>
        {/* Add more sidebar icons as needed */}
      </div>
    </div>
  )
}

const DesktopLayout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="flex h-screen w-screen">
      <div className="bg-green-200 flex flex-col w-full max-w-[22rem]">
        Hello!
      </div>
      <div className="bg-blue-200 flex flex-col w-full">
        {children}
      </div>
    </div>
  )
}

export default DesktopLayout