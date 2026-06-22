import { useState, useEffect } from 'react'
import { getCart } from '../api/cart'

export const useCartCount = () => {
  const [count, setCount] = useState(0)

  const refresh = () => {
    const token = localStorage.getItem('access_token')
    if (!token) {
      setCount(0)
      return
    }
    getCart()
      .then((items) => setCount(items.reduce((s, i) => s + i.quantity, 0)))
      .catch(() => setCount(0))
  }

  useEffect(() => {
    refresh()
    const interval = setInterval(refresh, 30000)
    return () => clearInterval(interval)
  }, [])

  return { count, refresh }
}
