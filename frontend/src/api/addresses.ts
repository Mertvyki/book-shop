import client from './client'
import type { Address } from '../types'

export const getAddresses = (userId: number) =>
  client.get<Address[]>(`/users/${userId}/addresses`).then((r) => r.data)

export const getAddress = (userId: number, addrId: number) =>
  client.get<Address>(`/users/${userId}/addresses/${addrId}`).then((r) => r.data)

export const createAddress = (
  userId: number,
  data: {
    street_address: string
    city: string
    postal_code: string
    country?: string
    is_default?: boolean
  },
) => client.post<Address>(`/users/${userId}/addresses`, data).then((r) => r.data)

export const patchAddress = (
  userId: number,
  addrId: number,
  payload: Record<string, unknown>,
) =>
  client
    .patch<Address>(`/users/${userId}/addresses/${addrId}`, payload)
    .then((r) => r.data)

export const deleteAddress = (userId: number, addrId: number) =>
  client.delete(`/users/${userId}/addresses/${addrId}`)
