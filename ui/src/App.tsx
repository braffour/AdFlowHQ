import { Routes, Route } from 'react-router-dom'
import { Workflows } from './pages/Workflows'
import { Dashboard } from './pages/Dashboard'
import { Layout } from './components/Layout'

function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Dashboard />} />
        <Route path="/workflows" element={<Workflows />} />
      </Routes>
    </Layout>
  )
}

export default App 