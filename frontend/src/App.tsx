import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { AuthProvider } from './contexts/AuthContext'
import { CartProvider } from './contexts/CartContext'
import { Layout } from './components/layout/Layout'
import { ProtectedRoute, StaffRoute, AdminRoute } from './components/layout/ProtectedRoute'
import { LoginPage } from './pages/auth/LoginPage'
import { RegisterPage } from './pages/auth/RegisterPage'
import { CatalogPage } from './pages/books/CatalogPage'
import { BookDetailPage } from './pages/books/BookDetailPage'
import { CartPage } from './pages/cart/CartPage'
import { CheckoutPage } from './pages/cart/CheckoutPage'
import { OrderHistoryPage } from './pages/orders/OrderHistoryPage'
import { OrderDetailPage } from './pages/orders/OrderDetailPage'
import { ProfilePage } from './pages/profile/ProfilePage'
import { AddressesPage } from './pages/profile/AddressesPage'
import { AdminBooksPage } from './pages/admin/AdminBooksPage'
import { AdminOrdersPage } from './pages/admin/AdminOrdersPage'
import { AdminEmployeesPage } from './pages/admin/AdminEmployeesPage'
import { AdminAuthorsPage } from './pages/admin/AdminAuthorsPage'
import { AdminCategoriesPage } from './pages/admin/AdminCategoriesPage'
import { AdminPublishersPage } from './pages/admin/AdminPublishersPage'

function App() {
  return (
    <BrowserRouter>
      <AuthProvider>
        <CartProvider>
          <Routes>
          <Route element={<Layout />}>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
            <Route path="/" element={<CatalogPage />} />
            <Route path="/catalog" element={<CatalogPage />} />
            <Route path="/books/:id" element={<BookDetailPage />} />

            <Route path="/cart" element={<ProtectedRoute><CartPage /></ProtectedRoute>} />
            <Route path="/checkout" element={<ProtectedRoute><CheckoutPage /></ProtectedRoute>} />
            <Route path="/orders" element={<ProtectedRoute><OrderHistoryPage /></ProtectedRoute>} />
            <Route path="/orders/:id" element={<ProtectedRoute><OrderDetailPage /></ProtectedRoute>} />
            <Route path="/profile" element={<ProtectedRoute><ProfilePage /></ProtectedRoute>} />
            <Route path="/profile/addresses" element={<ProtectedRoute><AddressesPage /></ProtectedRoute>} />

            <Route path="/admin/books" element={<StaffRoute><AdminBooksPage /></StaffRoute>} />
            <Route path="/admin/orders" element={<StaffRoute><AdminOrdersPage /></StaffRoute>} />
            <Route path="/admin/employees" element={<AdminRoute><AdminEmployeesPage /></AdminRoute>} />
            <Route path="/admin/authors" element={<StaffRoute><AdminAuthorsPage /></StaffRoute>} />
            <Route path="/admin/categories" element={<StaffRoute><AdminCategoriesPage /></StaffRoute>} />
            <Route path="/admin/publishers" element={<StaffRoute><AdminPublishersPage /></StaffRoute>} />
          </Route>
          </Routes>
        </CartProvider>
      </AuthProvider>
    </BrowserRouter>
  )
}

export default App
