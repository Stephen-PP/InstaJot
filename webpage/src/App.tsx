import React, { useState, useEffect, useLayoutEffect } from 'react'
import MobilePage from './mobile/page'
import DesktopPage from './desktop/page'
import DesktopLayout from './desktop/layout'
import MobileLayout from './mobile/layout'

import '../assets/index.css'

const App = () => {
  return (
    <DesktopLayout>
      <DesktopPage />
    </DesktopLayout>
  )
}

export default App