import client from './client'
import type { User } from '../types'

export const getUsers = (limit?: number, offset?: number) =>
  client
    .get<{ data: User[]; total: number }>('/users', { params: { limit, offset } })
    .then((r) => r.data)

export const getUser = (id: number) =>
  client.get<User>(`/users/${id}`).then((r) => r.data)

export const createUser = (data: {
  email: string
  password: string
  full_name: string
  phone_number?: string
}) => client.post<User>('/users', data).then((r) => r.data)

export const patchUser = (id: number, payload: Record<string, unknown>) =>
  client.patch<User>(`/users/${id}`, payload).then((r) => r.data)

export const deleteUser = (id: number) => client.delete(`/users/${id}`)
