import client from './client'
import type { Order, OrderItem, PaginatedResponse } from '../types'

export interface OrderWithItems extends Order {
  items: OrderItem[]
  delivery_cost: number
}

export const createOrder = (shipping_address_id?: number) =>
  client
    .post<OrderWithItems>('/orders', { shipping_address_id })
    .then((r) => r.data)

export const getOrders = (limit = 20, offset = 0) =>
  client
    .get<PaginatedResponse<Order>>('/orders', { params: { limit, offset } })
    .then((r) => r.data)

export const getOrder = (id: number) =>
  client.get<OrderWithItems>(`/orders/${id}`).then((r) => r.data)

export const getAllOrders = (limit = 20, offset = 0) =>
  client
    .get<PaginatedResponse<Order>>('/orders/all', { params: { limit, offset } })
    .then((r) => r.data)

export const patchOrderStatus = (id: number, status: string, version: number) =>
  client.patch(`/orders/${id}/status`, { status, version })
