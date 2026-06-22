import client from './client'
import type { CartItem } from '../types'

export const getCart = () =>
  client.get<CartItem[]>('/cart').then((r) => r.data)

export const addCartItem = (book_id: number, quantity: number) =>
  client.post<CartItem>('/cart/items', { book_id, quantity }).then((r) => r.data)

export const updateCartItem = (itemId: number, quantity: number) =>
  client.patch<CartItem>(`/cart/items/${itemId}`, { quantity }).then((r) => r.data)

export const removeCartItem = (itemId: number) =>
  client.delete(`/cart/items/${itemId}`)

export const clearCart = () => client.delete('/cart')
