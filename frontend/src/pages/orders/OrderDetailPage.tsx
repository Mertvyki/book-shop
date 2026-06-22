import { useEffect, useState } from 'react'
import { useParams, Link } from 'react-router-dom'
import { getOrder } from '../../api/orders'
import { ArrowLeft, Download } from 'lucide-react'
import type { OrderWithItems } from '../../api/orders'

const IMG_BASE = '/files/'

const statusLabels: Record<string, string> = {
  pending: 'Ожидает',
  paid: 'Оплачен',
  shipped: 'Отправлен',
  delivered: 'Доставлен',
  cancelled: 'Отменён',
}

export const OrderDetailPage = () => {
  const { id } = useParams<{ id: string }>()
  const [order, setOrder] = useState<OrderWithItems | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    if (!id) return
    setLoading(true)
    getOrder(Number(id))
      .then(setOrder)
      .finally(() => setLoading(false))
  }, [id])

  if (loading) return <div className="y2k-loading">Загрузка...</div>
  if (!order) return <div className="y2k-loading">Заказ не найден</div>

  return (
    <div className="max-w-[600px]">
      <Link to="/orders" className="y2k-link inline-flex items-center gap-1 mb-5">
        <ArrowLeft size={14} /> Мои заказы
      </Link>

      <div className="y2k-box-thick mb-5">
        <div className="flex items-center justify-between mb-3">
          <h1 className="y2k-title text-[18px]">Заказ #{order.id}</h1>
          <span className={`y2k-badge ${
            order.status === 'delivered' ? 'border-green-600 text-green-700' :
            order.status === 'cancelled' ? 'border-red-500 text-red-700' :
            'border-yellow-500 text-yellow-700'
          }`}>
            {statusLabels[order.status] || order.status}
          </span>
        </div>
        <p className="text-[11px] text-gray-400 mb-2">
          Создан: {new Date(order.created_at).toLocaleString('ru-RU')}
        </p>
        <p className="y2k-price text-[20px] mb-1">
          {order.total_amount.toLocaleString('ru-RU')} ₽
        </p>
        {order.delivery_cost > 0 && (
          <p className="text-[11px] text-gray-400">Доставка: {order.delivery_cost.toLocaleString('ru-RU')} ₽</p>
        )}
      </div>

      <h2 className="y2k-subtitle text-[12px] mb-4">Позиции заказа</h2>
      <div className="space-y-2">
        {order.items.map((item) => (
          <div key={item.id} className="y2k-box flex items-center justify-between">
            <div>
              <p className="text-[12px] font-bold text-gray-700">{item.title || `Книга #${item.book_id}`}</p>
              <p className="text-[11px] text-gray-400">{item.quantity} x {item.unit_price.toLocaleString('ru-RU')} ₽</p>
              {item.item_type === 'digital' && item.file_key && (
                <a href={`${IMG_BASE}${item.file_key}`} download className="y2k-btn-xp text-[9px] mt-1 inline-flex items-center gap-1 no-underline">
                  <Download size={10} /> СКАЧАТЬ
                </a>
              )}
            </div>
            <p className="font-bold text-[14px] text-gray-700 text-right">
              {(item.quantity * item.unit_price).toLocaleString('ru-RU')} ₽
            </p>
          </div>
        ))}
      </div>
    </div>
  )
}
