import React from 'react'

const DesktopPage = () => {
  return (
    <div className="h-screen flex flex-col">
      {/* Top bar */}
      <div className="bg-orange-700 w-full h-12 flex flex-row justify-between items-center relative">
        <div></div>
        <div className="text-white text-xl font-bold absolute left-1/2 transform -translate-x-1/2">Notes</div>
        <div className="text-white text-xl font-bold">N</div>
      </div>
      
      {/* Note content */}
      <div className="flex-grow p-4">
        Hello!
        {/* Note content goes here */}
      </div>
    </div>
  )
}

export default DesktopPage