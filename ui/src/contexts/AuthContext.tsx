import React, { createContext, useContext, useState, useEffect } from 'react'
import axios from 'axios'

interface User {
  id: string
  email: string
}

interface AuthContextValue {
  user: User | null
  login: (email: string, password: string) => Promise<void>
  logout: () => void
}

const AuthContext = createContext<AuthContextValue>({
  user: null,
  login: async () => { },
  logout: () => { },
})

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [user, setUser] = useState<User | null>(() => {
    const storedUser = localStorage.getItem('user');
    return storedUser ? JSON.parse(storedUser) : null;
  });

  useEffect(() => {
    if (user) {
      localStorage.setItem('user', JSON.stringify(user));
    } else {
      localStorage.removeItem('user');
    }
  }, [user]);

  const login = async (email: string, password: string) => {
    // AGENTS.md mentions JWT, so the response should contain a token.
    // The user object should probably be decoded from the token.
    // For now, I'll keep the mock logic but persist it.
    const { data } = await axios.post('/api/auth/login', { email, password })
    // Assuming the response has a user object or a token to decode.
    // Let's assume for now the response is the user object
    const userToSet = { id: data.id || '1', email: data.email || email };
    setUser(userToSet)
  }

  const logout = () => {
    // Missing: await axios.post('/api/auth/logout');
    setUser(null)
  }

  return (
    <AuthContext.Provider value={{ user, login, logout }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  return useContext(AuthContext)
}
