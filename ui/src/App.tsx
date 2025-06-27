import { Routes, Route } from 'react-router-dom'
import { Workflows } from './pages/Workflows'
import { Dashboard } from './pages/Dashboard'
import { Settings } from './pages/Settings'
import { Layout } from './components/Layout'

function App() {
  return (
    <Routes>
      <Route path="/" element={<Layout />}> 
        <Route index element={<Dashboard />} />
        <Route path="workflows" element={<Workflows />} />
        <Route path="settings" element={<Settings />} />
      </Route>
    </Routes>
  )
}

export default App 