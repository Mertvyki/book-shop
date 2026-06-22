import { useEffect, useState } from 'react'
import { useParams, Link } from 'react-router-dom'
import { getBook, checkPurchased } from '../../api/books'
import { addCartItem } from '../../api/cart'
import { Button } from '../../components/ui/Button'
import { ReviewSection } from '../../components/reviews/ReviewSection'
import { RatingDisplay } from '../../components/reviews/RatingDisplay'
import { useAuth } from '../../contexts/AuthContext'
import { useCartCount } from '../../contexts/CartContext'
import { ArrowLeft, ShoppingCart, Download } from 'lucide-react'
import client from '../../api/client'
import type { Book } from '../../types'

const IMG_BASE = '/files/'

export const BookDetailPage = () => {
  const { id } = useParams<{ id: string }>()
  const { isAuthenticated } = useAuth()
  const { refresh: refreshCart } = useCartCount()
  const [book, setBook] = useState<Book | null>(null)
  const [loading, setLoading] = useState(true)
  const [adding, setAdding] = useState(false)
  const [added, setAdded] = useState(false)
  const [purchased, setPurchased] = useState(false)

  useEffect(() => {
    if (!id) return
    setLoading(true)
    getBook(Number(id)).then(setBook).finally(() => setLoading(false))
  }, [id])

  useEffect(() => {
    if (!id || !isAuthenticated) return
    checkPurchased(Number(id)).then((r) => setPurchased(r.purchased)).catch(() => {})
  }, [id, isAuthenticated])

  const handleAddToCart = async () => {
    if (!book) return
    setAdding(true)
    try {
      await addCartItem(book.id, 1)
      refreshCart()
      setAdded(true)
      setTimeout(() => setAdded(false), 2000)
    } finally { setAdding(false) }
  }

  const handleDownload = async () => {
    if (!book) return
    try {
      const response = await client.get(`/books/${book.id}/download`, { responseType: 'blob' })
      const url = URL.createObjectURL(response.data)
      const a = document.createElement('a')
      a.href = url
      a.download = book.title || 'book'
      document.body.appendChild(a)
      a.click()
      document.body.removeChild(a)
      URL.revokeObjectURL(url)
    } catch { alert('Не удалось скачать книгу') }
  }

  if (loading) return <div className="y2k-loading">Загрузка...</div>
  if (!book) return <div className="y2k-loading">Книга не найдена</div>

  return (
    <div>
      <Link to="/" className="y2k-link inline-flex items-center gap-1 mb-5">
        <ArrowLeft size={14} /> Назад в каталог
      </Link>
      <div className="flex gap-6">
        <div className="w-[320px] shrink-0">
          <div className="aspect-[3/4] bg-gray-50 border-2 border-gray-200 rounded-[8px] overflow-hidden flex items-center justify-center">
            {book.cover_image_key ? (
              <img src={`${IMG_BASE}${book.cover_image_key}`} alt={book.title} className="w-full h-full object-cover" />
            ) : (
              <span className="text-[48px] text-gray-300">?</span>
            )}
          </div>
        </div>
        <div className="flex-1">
          <h1 className="y2k-title text-[22px] mb-1">{book.title}</h1>
          {book.avg_rating > 0 && (
            <div className="flex items-center gap-2 mb-1">
              <RatingDisplay rating={book.avg_rating} />
              <span className="text-[11px] text-gray-500">{book.avg_rating.toFixed(1)}</span>
            </div>
          )}
          <p className="text-[12px] text-gray-500 mb-3">{book.authors.map((a) => a.name).join(', ')}</p>
          {book.publisher && <p className="text-[11px] text-gray-400 mb-3">Издательство: {book.publisher.name}</p>}
          <div className="flex flex-wrap gap-2 mb-4">
            {book.categories.map((c) => (
              <span key={c.id} className="y2k-badge border-[#0066cc] text-[#0066cc]">{c.name}</span>
            ))}
            <span className={`y2k-badge ${book.book_type === 'digital' ? 'border-green-600 text-green-700' : 'border-blue-600 text-blue-700'}`}>
              {book.book_type === 'digital' ? 'Цифровая' : 'Физическая'}
            </span>
          </div>
          {book.description && <p className="text-[12px] text-gray-600 leading-5 mb-5">{book.description}</p>}
          {book.isbn && <p className="text-[11px] text-gray-400 mb-2">ISBN: {book.isbn}</p>}
          <p className="y2k-price text-[22px] mb-5">{book.price.toLocaleString('ru-RU')} ₽</p>
          {book.stock_quantity !== null && book.stock_quantity !== undefined && (
            <p className="text-[11px] text-gray-500 mb-4">{book.stock_quantity > 0 ? `В наличии: ${book.stock_quantity} шт.` : 'Нет в наличии'}</p>
          )}
          {isAuthenticated && (
            <div className="flex gap-3">
              <Button size="lg" onClick={handleAddToCart} loading={adding} disabled={added}>
                <ShoppingCart size={18} /> {added ? 'Добавлено!' : 'В корзину'}
              </Button>
              {book.book_type === 'digital' && purchased && (
                <Button size="lg" variant="secondary" onClick={handleDownload}>
                  <Download size={18} /> Скачать
                </Button>
              )}
            </div>
          )}
        </div>
      </div>
      <ReviewSection bookId={book.id} purchased={purchased} />
    </div>
  )
}
