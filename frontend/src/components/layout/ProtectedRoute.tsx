import { Navigate } from 'react-router-dom'
import { useAuth } from '../../contexts/AuthContext'

export const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated } = useAuth()
  if (!isAuthenticated) return <Navigate to="/login" replace />
  return <>{children}</>
}

export const StaffRoute = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated, isStaff } = useAuth()
  if (!isAuthenticated) return <Navigate to="/login" replace />
  if (!isStaff) return <Navigate to="/" replace />
  return <>{children}</>
}

export const AdminRoute = ({ children }: { children: React.ReactNode }) => {
  const { isAuthenticated, isAdmin } = useAuth()
  if (!isAuthenticated) return <Navigate to="/login" replace />
  if (!isAdmin) return <Navigate to="/" replace />
  return <>{children}</>
}
