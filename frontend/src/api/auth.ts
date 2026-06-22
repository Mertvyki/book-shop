import client from './client'
import type { LoginResponse } from '../types'

export const login = (email: string, password: string) =>
  client.post<LoginResponse>('/auth/login', { email, password }).then((r) => r.data)

export const refreshToken = (refresh_token: string) =>
  client
    .post<{ access_token: string; refresh_token: string }>('/auth/refresh', {
      refresh_token,
    })
    .then((r) => r.data)
