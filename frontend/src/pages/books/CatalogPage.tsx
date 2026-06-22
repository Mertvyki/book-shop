import { useEffect, useState } from 'react'
import { Link, useSearchParams } from 'react-router-dom'
import { getBooks, getAuthors, getCategories, getPublishers, type GetBooksParams } from '../../api/books'
import { addCartItem } from '../../api/cart'
import { Pagination } from '../../components/ui/Pagination'
import { ShoppingCart } from 'lucide-react'
import { useCartCount } from '../../contexts/CartContext'
import type { Book, Author, Category, Publisher } from '../../types'

const IMG_BASE = '/files/'

export const CatalogPage = () => {
  const [books, setBooks] = useState<Book[]>([])
  const [total, setTotal] = useState(0)
  const [totalPages, setTotalPages] = useState(0)
  const [page, setPage] = useState(1)
  const [search, setSearch] = useState('')
  const [typeFilter, setTypeFilter] = useState('')
  const [authorFilter, setAuthorFilter] = useState('')
  const [authorSearch, setAuthorSearch] = useState('')
  const [showAuthorDropdown, setShowAuthorDropdown] = useState(false)
  const [categoryFilter, setCategoryFilter] = useState('')
  const [publisherFilter, setPublisherFilter] = useState('')
  const [minPrice, setMinPrice] = useState('')
  const [maxPrice, setMaxPrice] = useState('')
  const [authors, setAuthors] = useState<Author[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [publishers, setPublishers] = useState<Publisher[]>([])
  const [loading, setLoading] = useState(true)
  const [addingId, setAddingId] = useState<number | null>(null)
  const [sortFilter, setSortFilter] = useState('')
  const [searchParams] = useSearchParams()
  const { refresh: refreshCart } = useCartCount()

  useEffect(() => {
    getAuthors().then(setAuthors).catch(() => {})
    getCategories().then(setCategories).catch(() => {})
    getPublishers().then(setPublishers).catch(() => {})
  }, [])

  useEffect(() => {
    const genre = searchParams.get('genre')
    if (genre && categories.length > 0) {
      const match = categories.find((c) => c.name.toUpperCase() === genre)
      if (match) {
        setCategoryFilter(String(match.id))
        setPage(1)
      }
    }
    const sort = searchParams.get('sort')
    if (sort) {
      setSortFilter(sort)
      setPage(1)
    } else {
      setSortFilter('')
    }
    const q = searchParams.get('search')
    if (q !== null && q !== search) {
      setSearch(q)
      setPage(1)
    }
  }, [searchParams, categories])

  useEffect(() => {
    setLoading(true)
    const params: GetBooksParams = { page, limit: 16 }
    if (search) params.search = search
    if (typeFilter) params.type = typeFilter
    if (authorFilter) params.author_id = Number(authorFilter)
    if (categoryFilter) params.category_id = Number(categoryFilter)
    if (publisherFilter) params.publisher_id = Number(publisherFilter)
    if (minPrice) params.min_price = Number(minPrice)
    if (maxPrice) params.max_price = Number(maxPrice)
    if (sortFilter) params.sort = sortFilter

    getBooks(params)
      .then((res) => { setBooks(res.data); setTotal(res.total); setTotalPages(res.total_pages) })
      .finally(() => setLoading(false))
  }, [page, search, typeFilter, authorFilter, categoryFilter, publisherFilter, minPrice, maxPrice, sortFilter])

  const handleAddToCart = async (bookId: number) => {
    setAddingId(bookId)
    try { await addCartItem(bookId, 1); refreshCart() } catch { alert('Ошибка при добавлении') }
    finally { setAddingId(null) }
  }

  const hasFilters = search || typeFilter || authorFilter || categoryFilter || publisherFilter || minPrice || maxPrice || sortFilter

  const resetFilters = () => {
    setSearch('')
    setTypeFilter('')
    setAuthorFilter('')
    setCategoryFilter('')
    setPublisherFilter('')
    setMinPrice('')
    setMaxPrice('')
    setSortFilter('')
    setPage(1)
  }

  return (
    <div className="flex gap-5">
      {/* Sidebar */}
      <aside className="w-[200px] shrink-0">
        <div className="y2k-sidebar-block">
          <div className="y2k-sidebar-header">ПОПУЛЯРНЫЕ АВТОРЫ</div>
          {authors.slice(0, 6).map((a) => (
            <button key={a.id} className="y2k-sidebar-item w-full text-left" onClick={() => { setAuthorFilter(String(a.id)); setPage(1) }}>
              {a.name}
            </button>
          ))}
        </div>

        <div className="y2k-sidebar-block">
          <div className="y2k-sidebar-header">ЖАНРЫ</div>
          {categories.slice(0, 8).map((c) => (
            <button key={c.id} className="y2k-sidebar-item w-full text-left" onClick={() => { setCategoryFilter(String(c.id)); setPage(1) }}>
              {c.name}
            </button>
          ))}
        </div>

        <div className="y2k-sidebar-block">
          <div className="y2k-sidebar-header">ГОРЯЧИЕ НОВИНКИ</div>
          <div className="px-3 py-2 space-y-1">
            <Link to="/catalog" className="block text-[10px] text-gray-500 no-underline hover:text-[#ff8a00] font-bold">ЕЖЕНЕДЕЛЬНЫЕ ПОДБОРКИ</Link>
            <Link to="/catalog?sort=bestsellers" className="block text-[10px] text-gray-500 no-underline hover:text-[#ff8a00] font-bold">ЛИДЕРЫ ПРОДАЖ</Link>
            <Link to="/catalog?sort=newest" className="block text-[10px] text-[#0066cc] no-underline hover:text-[#ff8a00] font-bold">НОВЫЕ ПОСТУПЛЕНИЯ</Link>
          </div>
        </div>
      </aside>

      {/* Main content */}
      <div className="flex-1 min-w-0">
        {/* Search bar */}
        <div className="flex items-center justify-between mb-5">
          <h1 className="y2k-title">КАТАЛОГ КНИГ</h1>
          <div className="flex items-center gap-2">
            <div className="flex items-center">
              <input
                className="y2k-input w-[180px] border-r-0 rounded-r-none"
                placeholder="Поиск книг..."
                value={search}
                onChange={(e) => { setSearch(e.target.value); setPage(1) }}
              />
              <span className="bg-gray-100 border border-gray-300 rounded-r-[4px] px-2 py-[4px] text-gray-400 text-[11px] ml-px">&raquo;</span>
            </div>
            <select className="y2k-select w-[100px]" value={typeFilter} onChange={(e) => { setTypeFilter(e.target.value); setPage(1) }}>
              <option value="">ВСЕ</option>
              <option value="physical">ПЕЧАТНЫЕ</option>
              <option value="digital">ЦИФРОВЫЕ</option>
            </select>
          </div>
        </div>

        {/* Filters row */}
        <div className="flex items-center gap-3 mb-4 text-[11px]">
          <select className="y2k-select flex-1" value={categoryFilter} onChange={(e) => { setCategoryFilter(e.target.value); setPage(1) }}>
            <option value="">ВСЕ ЖАНРЫ</option>
            {categories.map((c) => (<option key={c.id} value={c.id}>{c.name}</option>))}
          </select>
          <select className="y2k-select flex-1" value={publisherFilter} onChange={(e) => { setPublisherFilter(e.target.value); setPage(1) }}>
            <option value="">ВСЕ ИЗДАТЕЛИ</option>
            {publishers.map((p) => (<option key={p.id} value={p.id}>{p.name}</option>))}
          </select>
          <input className="y2k-input w-[70px]" placeholder="От" type="number" value={minPrice} onChange={(e) => { setMinPrice(e.target.value); setPage(1) }} />
          <input className="y2k-input w-[70px]" placeholder="До" type="number" value={maxPrice} onChange={(e) => { setMaxPrice(e.target.value); setPage(1) }} />
          {hasFilters && (
            <button className="y2k-btn y2k-btn-danger text-[11px] px-3 py-1" onClick={resetFilters}>
              СБРОСИТЬ
            </button>
          )}
        </div>

        {/* Book grid */}
        {loading ? (
          <div className="y2k-loading">Загрузка...</div>
        ) : books.length === 0 ? (
          <div className="y2k-loading">Книги не найдены</div>
        ) : (
          <>
            <div className="grid grid-cols-4 gap-4">
              {books.map((book) => (
                <div key={book.id} className="y2k-card-blue">
                  <div className="y2k-card-header">
                    {book.book_type === 'digital' ? 'ЦИФРОВАЯ' : 'ПЕЧАТНАЯ'}
                  </div>
                  <div className="p-3">
                    <div className="w-[120px] h-[180px] mx-auto bg-white border border-gray-300 flex items-center justify-center overflow-hidden rounded-[3px] mb-3 shadow-inner">
                      {book.cover_image_key ? (
                        <img src={`${IMG_BASE}${book.cover_image_key}`} alt={book.title} className="w-full h-full object-cover" />
                      ) : (
                        <span className="text-gray-300 text-[14px] font-bold">НЕТ ОБЛОЖКИ</span>
                      )}
                    </div>
                    <Link to={`/books/${book.id}`} className="no-underline">
                      <h3 className="text-[11px] font-bold text-[#003366] uppercase leading-tight mb-1">{book.title}</h3>
                    </Link>
                    <p className="text-[10px] text-gray-500 mb-2">{book.authors.map((a) => a.name).join(', ')}</p>
                    <p className="y2k-price text-[14px] mb-3">{book.price.toLocaleString('ru-RU')} ₽</p>
                    <div className="flex gap-0.5">
                      <button
                        className="y2k-btn-xp y2k-btn-xp-orange text-[8px] px-1.5 py-1 flex items-center gap-0.5"
                        onClick={() => handleAddToCart(book.id)}
                        disabled={addingId === book.id}
                      >
                        <ShoppingCart size={8} /> {addingId === book.id ? '...' : 'В КОРЗИНУ'}
                      </button>
                      <Link to={`/books/${book.id}`} className="y2k-btn-xp text-[8px] px-1.5 py-1 no-underline">ПОДРОБНЕЕ</Link>
                    </div>
                  </div>
                </div>
              ))}
            </div>
            <Pagination page={page} totalPages={totalPages} total={total} onPageChange={setPage} />
          </>
        )}
      </div>
    </div>
  )
}
