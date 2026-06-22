import { createContext, useContext, useState, useEffect, useCallback, type ReactNode } from 'react'
import { getCart } from '../api/cart'

interface CartContextType {
  count: number
  refresh: () => void
}

const CartContext = createContext<CartContextType>({ count: 0, refresh: () => {} })

export const CartProvider = ({ children }: { children: ReactNode }) => {
  const [count, setCount] = useState(0)

  const refresh = useCallback(() => {
    const token = localStorage.getItem('access_token')
    if (!token) {
      setCount(0)
      return
    }
    getCart()
      .then((items) => setCount(items.reduce((s, i) => s + i.quantity, 0)))
      .catch(() => setCount(0))
  }, [])

  useEffect(() => {
    refresh()
    const interval = setInterval(refresh, 30000)
    return () => clearInterval(interval)
  }, [refresh])

  return (
    <CartContext.Provider value={{ count, refresh }}>
      {children}
    </CartContext.Provider>
  )
}

export const useCartCount = () => useContext(CartContext)
