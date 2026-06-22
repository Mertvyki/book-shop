import { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Trash2, Minus, Plus, ShoppingBag } from 'lucide-react'
import { getCart, updateCartItem, removeCartItem, clearCart } from '../../api/cart'
import { Button } from '../../components/ui/Button'
import { useCartCount } from '../../contexts/CartContext'
import type { CartItem } from '../../types'

const IMG_BASE = '/files/'

export const CartPage = () => {
  const navigate = useNavigate()
  const [items, setItems] = useState<CartItem[]>([])
  const [loading, setLoading] = useState(true)
  const { refresh: refreshCart } = useCartCount()

  const fetchCart = () => {
    setLoading(true)
    getCart().then(setItems).finally(() => setLoading(false))
  }

  useEffect(() => { fetchCart() }, [])

  const handleQuantity = async (item: CartItem, delta: number) => {
    const newQty = item.quantity + delta
    if (newQty < 1) return
    await updateCartItem(item.id, newQty)
    fetchCart()
    refreshCart()
  }

  const handleRemove = async (id: number) => {
    await removeCartItem(id)
    fetchCart()
    refreshCart()
  }

  const handleClear = async () => {
    await clearCart()
    setItems([])
    refreshCart()
  }

  const total = items.reduce((s, i) => s + i.price * i.quantity, 0)
  const hasPhysical = items.some((i) => i.book_type === 'physical')
  const deliveryCost = hasPhysical ? 250 : 0

  if (loading) return <div className="y2k-loading">Загрузка...</div>

  return (
    <div>
      <div className="flex items-center justify-between mb-5">
        <h1 className="y2k-title">Корзина</h1>
        {items.length > 0 && <Button variant="danger" size="sm" onClick={handleClear}>Очистить</Button>}
      </div>

      {items.length === 0 ? (
        <div className="text-center py-16">
          <ShoppingBag className="mx-auto text-gray-300 mb-4" size={48} />
          <p className="text-[12px] text-gray-400 uppercase tracking-wide mb-4">Корзина пуста</p>
          <Link to="/"><Button>В каталог</Button></Link>
        </div>
      ) : (
        <div className="space-y-3">
          {items.map((item) => (
            <div key={item.id} className="y2k-box flex items-center gap-4">
              <div className="w-14 h-[68px] bg-gray-50 border border-gray-200 rounded-[4px] flex items-center justify-center overflow-hidden shrink-0">
                {item.cover_image_key ? <img src={`${IMG_BASE}${item.cover_image_key}`} alt="" className="w-full h-full object-cover" /> : <span className="text-[18px] text-gray-300">?</span>}
              </div>
              <div className="flex-1 min-w-0">
                <Link to={`/books/${item.book_id}`} className="text-[12px] font-bold text-gray-700 hover:text-[#ff8a00] no-underline">{item.title}</Link>
                <p className="y2k-price text-[13px] mt-1">{item.price.toLocaleString('ru-RU')} ₽</p>
              </div>
              <div className="flex items-center gap-2">
                <button className="y2k-btn y2k-btn-secondary text-[11px] px-2 py-1" onClick={() => handleQuantity(item, -1)}><Minus size={12} /></button>
                <span className="w-8 text-center font-bold text-[13px]">{item.quantity}</span>
                <button className="y2k-btn y2k-btn-secondary text-[11px] px-2 py-1" onClick={() => handleQuantity(item, 1)}><Plus size={12} /></button>
              </div>
              <p className="font-bold text-[14px] text-gray-700 w-[100px] text-right">{(item.price * item.quantity).toLocaleString('ru-RU')} ₽</p>
              <button className="text-gray-300 hover:text-red-500" onClick={() => handleRemove(item.id)}><Trash2 size={16} /></button>
            </div>
          ))}
          <div className="y2k-box-thick flex items-center justify-between">
            <div>
              <p className="text-[11px] text-gray-400 uppercase tracking-wide">Итого</p>
              <p className="y2k-price text-[22px]">{total.toLocaleString('ru-RU')} ₽</p>
              {deliveryCost > 0 && <p className="text-[11px] text-gray-400 mt-1">Доставка: {deliveryCost.toLocaleString('ru-RU')} ₽</p>}
            </div>
            <Button size="lg" onClick={() => navigate('/checkout')}>Оформить заказ</Button>
          </div>
        </div>
      )}
    </div>
  )
}
