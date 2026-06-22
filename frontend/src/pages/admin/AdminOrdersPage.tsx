import { useEffect, useState } from 'react'
import { getAllOrders, patchOrderStatus } from '../../api/orders'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Pagination } from '../../components/ui/Pagination'
import type { Order } from '../../types'

const statuses = ['pending', 'paid', 'shipped', 'delivered', 'cancelled']
const statusColors: Record<string, string> = {
  pending: 'border-yellow-500 text-yellow-700',
  paid: 'border-blue-500 text-blue-700',
  shipped: 'border-purple-500 text-purple-700',
  delivered: 'border-green-600 text-green-700',
  cancelled: 'border-red-500 text-red-700',
}
const statusLabels: Record<string, string> = {
  pending: 'ОЖИДАЕТ', paid: 'ОПЛАЧЕН', shipped: 'ОТПРАВЛЕН', delivered: 'ДОСТАВЛЕН', cancelled: 'ОТМЕНЁН',
}

export const AdminOrdersPage = () => {
  const [orders, setOrders] = useState<Order[]>([])
  const [total, setTotal] = useState(0)
  const [totalPages, setTotalPages] = useState(0)
  const [page, setPage] = useState(1)
  const [loading, setLoading] = useState(true)
  const [changing, setChanging] = useState<number | null>(null)

  const fetchOrders = () => {
    setLoading(true)
    getAllOrders(10, (page - 1) * 10)
      .then((res) => { setOrders(res.data); setTotal(res.total); setTotalPages(res.total_pages) })
      .finally(() => setLoading(false))
  }

  useEffect(() => { fetchOrders() }, [page])

  const handleStatusChange = async (order: Order, newStatus: string) => {
    setChanging(order.id)
    try { await patchOrderStatus(order.id, newStatus, order.version); fetchOrders() }
    finally { setChanging(null) }
  }

  if (loading) return <AdminLayout title="ЗАКАЗЫ"><div className="y2k-loading">Загрузка...</div></AdminLayout>

  return (
    <AdminLayout title="ЗАКАЗЫ">
      <div className="y2k-box overflow-hidden p-0">
        <table className="y2k-table">
          <thead><tr><th>ID</th><th>ПОЛЬЗОВАТЕЛЬ</th><th className="text-right">СУММА</th><th className="text-center">СТАТУС</th><th className="text-center">ДЕЙСТВИЕ</th></tr></thead>
          <tbody>
            {orders.map((order) => (
              <tr key={order.id}>
                <td className="font-bold text-gray-700">#{order.id}</td>
                <td className="text-gray-500">#{order.user_id}</td>
                <td className="text-right font-bold text-gray-700">{order.total_amount.toLocaleString('ru-RU')} ₽</td>
                <td className="text-center"><span className={`y2k-badge ${statusColors[order.status]}`}>{statusLabels[order.status]}</span></td>
                <td className="text-center">
                  <select className="y2k-select text-[10px]" value={order.status} onChange={(e) => handleStatusChange(order, e.target.value)} disabled={changing === order.id}>
                    {statuses.map((s) => (<option key={s} value={s}>{statusLabels[s]}</option>))}
                  </select>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
      <Pagination page={page} totalPages={totalPages} total={total} onPageChange={setPage} />
    </AdminLayout>
  )
}
