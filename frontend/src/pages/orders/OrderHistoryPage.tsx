import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { getOrders } from '../../api/orders'
import { Pagination } from '../../components/ui/Pagination'
import type { Order } from '../../types'

const statusColors: Record<string, string> = {
  pending: 'border-yellow-500 text-yellow-700',
  paid: 'border-blue-500 text-blue-700',
  shipped: 'border-purple-500 text-purple-700',
  delivered: 'border-green-600 text-green-700',
  cancelled: 'border-red-500 text-red-700',
}

const statusLabels: Record<string, string> = {
  pending: 'Ожидает',
  paid: 'Оплачен',
  shipped: 'Отправлен',
  delivered: 'Доставлен',
  cancelled: 'Отменён',
}

export const OrderHistoryPage = () => {
  const [orders, setOrders] = useState<Order[]>([])
  const [total, setTotal] = useState(0)
  const [totalPages, setTotalPages] = useState(0)
  const [page, setPage] = useState(1)
  const [loading, setLoading] = useState(true)
  const limit = 10

  useEffect(() => {
    setLoading(true)
    getOrders(limit, (page - 1) * limit)
      .then((res) => {
        setOrders(res.data)
        setTotal(res.total)
        setTotalPages(res.total_pages)
      })
      .finally(() => setLoading(false))
  }, [page])

  if (loading) return <div className="y2k-loading">Загрузка...</div>

  return (
    <div>
      <h1 className="y2k-title mb-5">Мои заказы</h1>
      {orders.length === 0 ? (
        <p className="y2k-loading">У вас ещё нет заказов</p>
      ) : (
        <>
          <div className="space-y-3">
            {orders.map((order) => (
              <Link key={order.id} to={`/orders/${order.id}`} className="y2k-box-thick block no-underline hover:border-[#ff8a00] transition-colors">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-[13px] font-bold text-gray-700">Заказ #{order.id}</p>
                    <p className="text-[11px] text-gray-400">{new Date(order.created_at).toLocaleDateString('ru-RU')}</p>
                  </div>
                  <div className="text-right">
                    <p className="y2k-price text-[15px]">{order.total_amount.toLocaleString('ru-RU')} ₽</p>
                    <span className={`y2k-badge ${statusColors[order.status] || ''}`}>
                      {statusLabels[order.status] || order.status}
                    </span>
                  </div>
                </div>
              </Link>
            ))}
          </div>
          <Pagination page={page} totalPages={totalPages} total={total} onPageChange={setPage} />
        </>
      )}
    </div>
  )
}
