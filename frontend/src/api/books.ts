import client from './client'
import type { Book, PaginatedResponse, Author, Category, Publisher } from '../types'

export interface GetBooksParams {
  type?: string
  author_id?: number
  category_id?: number
  publisher_id?: number
  search?: string
  min_price?: number
  max_price?: number
  sort?: string
  page?: number
  limit?: number
}

export const getBooks = (params: GetBooksParams) =>
  client.get<PaginatedResponse<Book>>('/books', { params }).then((r) => r.data)

export const getBook = (id: number) =>
  client.get<Book>(`/books/${id}`).then((r) => r.data)

export const createBook = (data: FormData) =>
  client.post<Book>('/books', data, {
    headers: { 'Content-Type': 'multipart/form-data' },
  }).then((r) => r.data)

export const patchBook = (id: number, payload: unknown, cover?: File, bookFile?: File) => {
  const fd = new FormData()
  fd.append('request', JSON.stringify(payload))
  if (cover) fd.append('cover_image', cover)
  if (bookFile) fd.append('book_file', bookFile)
  return client
    .patch<Book>(`/books/${id}`, fd, {
      headers: { 'Content-Type': 'multipart/form-data' },
    })
    .then((r) => r.data)
}

export const deleteBook = (id: number) =>
  client.delete(`/books/${id}`).then((r) => r.data)

export const checkPurchased = (id: number) =>
  client.get<{ purchased: boolean }>(`/books/${id}/purchased`).then((r) => r.data)

export const createAuthor = (data: { name: string; bio?: string | null; birth_year?: number | null }) =>
  client.post<Author>('/authors', data).then((r) => r.data)

export const patchAuthor = (id: number, data: { name?: string; bio?: string | null; birth_year?: number | null }) =>
  client.patch<Author>(`/authors/${id}`, data).then((r) => r.data)

export const deleteAuthor = (id: number) =>
  client.delete(`/authors/${id}`)

export const createCategory = (data: { name: string; slug?: string; description?: string | null }) =>
  client.post<Category>('/categories', data).then((r) => r.data)

export const patchCategory = (id: number, data: { name?: string; slug?: string; description?: string | null }) =>
  client.patch<Category>(`/categories/${id}`, data).then((r) => r.data)

export const deleteCategory = (id: number) =>
  client.delete(`/categories/${id}`)

export const getAuthors = () =>
  client.get<Author[]>('/authors').then((r) => r.data)

export const getCategories = () =>
  client.get<Category[]>('/categories').then((r) => r.data)

export const getPublishers = () =>
  client.get<Publisher[]>('/publishers').then((r) => r.data)

export const createPublisher = (name: string) =>
  client.post<Publisher>('/publishers', { name }).then((r) => r.data)

export const patchPublisher = (id: number, name: string) =>
  client.patch<Publisher>(`/publishers/${id}`, { name }).then((r) => r.data)

export const deletePublisher = (id: number) =>
  client.delete(`/publishers/${id}`)
