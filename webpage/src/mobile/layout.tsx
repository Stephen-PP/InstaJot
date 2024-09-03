import React from 'react'

const MobileLayout = ({ children }: { children: React.ReactNode }) => {
  console.log("mobile rendering")
  return <div className="mobile-layout">{children}</div>
}

export default MobileLayout