import client from './client'
import type { Review } from '../types'

export const getBookReviews = (bookId: number) =>
  client.get<Review[]>(`/books/${bookId}/reviews`).then((r) => r.data)

export const getUserReview = (bookId: number) =>
  client.get<Review>(`/books/${bookId}/reviews/mine`).then((r) => r.data)

export const upsertReview = (bookId: number, data: { rating: number; title?: string | null; body?: string | null }) =>
  client.put<Review>(`/books/${bookId}/reviews`, data).then((r) => r.data)
