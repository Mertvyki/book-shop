import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { createOrder } from '../../api/orders'
import { getAddresses } from '../../api/addresses'
import { getCart } from '../../api/cart'
import { useAuth } from '../../contexts/AuthContext'
import { useCartCount } from '../../contexts/CartContext'
import { Button } from '../../components/ui/Button'
import type { Address, CartItem } from '../../types'

export const CheckoutPage = () => {
  const { user } = useAuth()
  const navigate = useNavigate()
  const { refresh: refreshCart } = useCartCount()
  const [addresses, setAddresses] = useState<Address[]>([])
  const [selectedAddr, setSelectedAddr] = useState<number | undefined>(undefined)
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [cartItems, setCartItems] = useState<CartItem[]>([])

  useEffect(() => {
    if (!user) return
    Promise.all([getAddresses(user.id), getCart()])
      .then(([addr, cart]) => {
        setAddresses(addr)
        setCartItems(cart)
        const def = addr.find((a) => a.is_default)
        if (def) setSelectedAddr(def.id)
      })
      .finally(() => setLoading(false))
  }, [user])

  const hasPhysical = cartItems.some((i) => i.book_type === 'physical')

  const handleSubmit = async () => {
    setSubmitting(true)
    try {
      const order = await createOrder(selectedAddr)
      refreshCart()
      navigate(`/orders/${order.id}`)
    } catch { alert('Ошибка при оформлении заказа') }
    finally { setSubmitting(false) }
  }

  if (loading) return <div className="y2k-loading">Загрузка...</div>

  return (
    <div className="max-w-[500px] mx-auto">
      <h1 className="y2k-title mb-5">Оформление заказа</h1>
      {hasPhysical ? (
        <div className="y2k-box-thick mb-5">
          <h2 className="y2k-subtitle text-[12px] mb-4">Адрес доставки</h2>
          <p className="text-[11px] text-red-600 mb-3 font-bold">Для физических книг требуется адрес доставки</p>
          {addresses.length === 0 ? (
            <p className="text-[11px] text-gray-400">
              Нет сохранённых адресов. Добавьте в{' '}
              <button className="y2k-link" onClick={() => navigate('/profile/addresses')}>профиле</button>
            </p>
          ) : (
            <div className="space-y-2">
              {addresses.map((addr) => (
                <label key={addr.id} className="flex items-start gap-3 p-3 border border-gray-200 rounded-[6px] cursor-pointer hover:border-[#ff8a00] text-[12px]">
                  <input type="radio" name="address" checked={selectedAddr === addr.id} onChange={() => setSelectedAddr(addr.id)} className="mt-0.5" />
                  <div>
                    <p className="text-gray-700">{addr.street_address}</p>
                    <p className="text-[10px] text-gray-400">{addr.city}, {addr.postal_code}</p>
                  </div>
                </label>
              ))}
            </div>
          )}
        </div>
      ) : (
        <div className="y2k-box-thick mb-5">
          <p className="text-[11px] text-gray-500 font-bold">Только цифровые книги — адрес доставки не требуется</p>
        </div>
      )}
      <Button className="w-full" size="lg" onClick={handleSubmit} loading={submitting} disabled={hasPhysical && !selectedAddr}>
        Подтвердить заказ
      </Button>
    </div>
  )
}
