import { useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { ShoppingCart, LogOut, User, Shield, Users, ClipboardList } from 'lucide-react'
import { useAuth } from '../../contexts/AuthContext'
import { useCartCount } from '../../contexts/CartContext'

export const Header = () => {
  const { logout, isAuthenticated, isStaff, isAdmin, user } = useAuth()
  const navigate = useNavigate()
  const { count } = useCartCount()
  const [searchQuery, setSearchQuery] = useState('')

  const handleLogout = () => { logout(); navigate('/login') }
  const handleSearch = () => {
    const q = searchQuery.trim()
    if (q) navigate(`/catalog?search=${encodeURIComponent(q)}`)
  }

  return (
    <header className="bg-white border-b border-gray-250">
      <div className="max-w-[1000px] mx-auto px-4">
        {/* Main header bar */}
        <div className="flex items-center justify-between gap-x-8 py-3">
          {/* Logo + slogan */}
          <div>
            <Link to="/" className="text-[28px] font-bold tracking-[4px] text-[#003366] no-underline leading-none" style={{ fontFamily: 'Arial, sans-serif' }}>
              BOOK<span className="text-[#ff8a00]">ZONE</span>
            </Link>
            <p className="text-[9px] text-gray-400 tracking-[2px] uppercase mt-0.5 font-bold">
              книги • медиа • истории
            </p>
          </div>

          {/* Search */}
          <div className="flex items-center gap-2">
            <div className="flex items-center">
              <input
                className="y2k-input w-[200px] border-r-0 rounded-r-none text-[11px]"
                placeholder="Поиск книг..."
                value={searchQuery}
                onChange={(e) => setSearchQuery(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && handleSearch()}
              />
              <button className="y2k-btn text-[10px] px-3 py-[4px] rounded-l-none" onClick={handleSearch}>
                ИСКАТЬ
              </button>
            </div>
          </div>

          {/* User panel */}
          <div className="flex items-center gap-1">
            <Link to="/" className="y2k-btn-xp"><User size={12} /> КАТАЛОГ</Link>
            <Link to="/cart" className="y2k-btn-xp relative">
              <ShoppingCart size={12} /> КОРЗИНА{count > 0 && <span className="text-[#ff8a00] ml-0.5">({count})</span>}
            </Link>

            {isAuthenticated && (
              <>
                <Link to="/orders" className="y2k-btn-xp"><ClipboardList size={12} /> ЗАКАЗЫ</Link>
                <Link to="/profile" className="y2k-btn-xp"><User size={12} /> ПРОФИЛЬ</Link>
                {isStaff && (
                  <Link to="/admin/books" className="y2k-btn-xp"><Shield size={12} /> АДМИН</Link>
                )}
                {isAdmin && (
                  <Link to="/admin/employees" className="y2k-btn-xp"><Users size={12} /> СОТР.</Link>
                )}
                <button onClick={handleLogout} className="y2k-btn-xp flex items-center gap-1">
                  <LogOut size={12} /> ВЫХОД
                </button>
              </>
            )}

            {!isAuthenticated && (
              <>
                <Link to="/login" className="y2k-btn-xp">ВХОД</Link>
                <Link to="/register" className="y2k-btn-xp y2k-btn-xp-orange">РЕГИСТРАЦИЯ</Link>
              </>
            )}
          </div>
        </div>

        {/* Decorative Y2K line */}
        <div className="y2k-deco-line" />

        {/* Navigation menu */}
        <div className="bg-[#f8f9fa] border-x border-gray-200 border-b">
          <div className="flex divide-x divide-gray-200 text-[11px] font-bold uppercase tracking-wide">
            <Link to="/catalog?genre=ХУДОЖЕСТВЕННАЯ ЛИТЕРАТУРА" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ХУДОЖЕСТВЕННАЯ</Link>
            <Link to="/catalog?genre=ФАНТАСТИКА" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ФАНТАСТИКА</Link>
            <Link to="/catalog?genre=ТЕХНОЛОГИИ" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ТЕХНОЛОГИИ</Link>
            <Link to="/catalog?genre=ИСТОРИЯ" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ИСТОРИЯ</Link>
            <Link to="/catalog?genre=НАУЧНАЯ ЛИТЕРАТУРА" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">НАУЧНАЯ</Link>
            <Link to="/catalog?genre=ДЕТСКАЯ ЛИТЕРАТУРА" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ДЕТСКАЯ</Link>
            <Link to="/catalog" className="px-3 py-2 text-gray-400 no-underline hover:text-[#ff8a00]">ВСЕ КНИГИ</Link>
            <span className="px-3 py-2 text-gray-400 line-through decoration-1">АКЦИИ</span>
          </div>
        </div>

        {/* User mini-toolbar */}
        <div className="y2k-toolbar flex items-center justify-between mt-1 mb-0">
          <span className="font-bold text-[#556]">
            С ПРИХОДОМ, <span className="text-[#0066cc]">{(user?.full_name || 'ГОСТЬ').toUpperCase()}</span>
          </span>
          <div className="flex items-center gap-3 text-[10px]">
            <Link to="/catalog" className="text-gray-500 no-underline hover:text-[#ff8a00] font-bold">ЕЖЕНЕДЕЛЬНЫЕ ПОДБОРКИ</Link>
            <span className="text-gray-400">|</span>
            <Link to="/catalog?sort=bestsellers" className="text-gray-500 no-underline hover:text-[#ff8a00] font-bold">ЛИДЕРЫ ПРОДАЖ</Link>
            <span className="text-gray-400">|</span>
            <Link to="/catalog?sort=newest" className="text-[#0066cc] no-underline hover:text-[#ff8a00] font-bold">НОВЫЕ ПОСТУПЛЕНИЯ</Link>
          </div>
        </div>
      </div>
    </header>
  )
}
