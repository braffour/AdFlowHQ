import { Routes, Route } from 'react-router-dom'
import { Workflows } from './pages/Workflows'
import { Dashboard } from './pages/Dashboard'
import { Login } from './pages/Login'
import { Layout } from './components/Layout'
import { AuthProvider } from './contexts/AuthContext'

function App() {
  return (
    <AuthProvider>
      <Layout>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Dashboard />} />
          <Route path="/workflows" element={<Workflows />} />
        </Routes>
      </Layout>
    </AuthProvider>
  )
}

export default App
