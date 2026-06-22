export interface User {
  id: number
  email: string
  full_name: string
  phone_number: string | null
  role: string
  created_at: string
}

export interface Author {
  id: number
  name: string
  bio: string | null
  birth_year: number | null
  created_at: string
}

export interface Category {
  id: number
  name: string
  slug: string
  description: string | null
}

export interface Publisher {
  id: number
  name: string
}

export interface Book {
  id: number
  version: number
  title: string
  description: string | null
  isbn: string | null
  price: number
  book_type: string
  stock_quantity: number | null
  file_key: string | null
  cover_image_key: string | null
  publisher: Publisher | null
  authors: Author[]
  categories: Category[]
  avg_rating: number
  created_at: string
}

export interface Review {
  id: number
  version: number
  book_id: number
  user_id: number
  rating: number
  title: string | null
  body: string | null
  created_at: string
  updated_at: string
  user_name: string
}

export interface CartItem {
  id: number
  version: number
  user_id: number
  book_id: number
  quantity: number
  added_at: string
  title: string
  price: number
  cover_image_key: string | null
  book_type: string
}

export interface Order {
  id: number
  version: number
  user_id: number
  status: string
  total_amount: number
  shipping_address_id: number | null
  payment_method: string | null
  created_at: string
  updated_at: string
}

export interface OrderItem {
  id: number
  version: number
  order_id: number
  book_id: number
  quantity: number
  unit_price: number
  item_type: string
  title: string
  file_key: string | null
}

export interface Address {
  id: number
  version: number
  user_id: number
  street_address: string
  city: string
  postal_code: string
  country: string
  is_default: boolean
  created_at: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
  total_pages: number
}

export interface LoginResponse {
  access_token: string
  refresh_token: string
  user: User
}

export interface CreateBookPayload {
  title: string
  price: number
  book_type: string
  description?: string
  isbn?: string
  stock_quantity?: number
  publisher_id?: number
  author_ids?: string
  category_ids?: string
  cover_image: File
  book_file?: File
}

export interface PatchBookPayload {
  title?: { Value?: string; Set: boolean }
  description?: { Value?: string; Set: boolean }
  isbn?: { Value?: string; Set: boolean }
  price?: { Value?: number; Set: boolean }
  book_type?: { Value?: string; Set: boolean }
  stock_quantity?: { Value?: number; Set: boolean }
  publisher_id?: { Value?: number; Set: boolean }
  author_ids?: number[]
  category_ids?: number[]
}
