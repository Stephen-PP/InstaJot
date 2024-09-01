import React, { useState, useEffect } from 'react'
import MobilePage from './mobile/page'
import DesktopPage from './desktop/page'
import DesktopLayout from './desktop/layout'
import MobileLayout from './mobile/layout'

import '../assets/index.css'

const App = () => {
  const [isMobile, setIsMobile] = useState(window.innerWidth < 768)

  useEffect(() => {
    const handleResize = () => {
      setIsMobile(window.innerWidth < 768)
    }

    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  }, [])

  return isMobile ? (
    <MobileLayout>
      <MobilePage />
    </MobileLayout>
  ) : (
    <DesktopLayout>
      <DesktopPage />
    </DesktopLayout>
  )
}

export default App