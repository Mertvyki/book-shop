import { useEffect, useState, type FormEvent } from 'react'
import { getUsers, createUser, patchUser, deleteUser } from '../../api/users'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Trash2 } from 'lucide-react'
import type { User } from '../../types'

export const AdminEmployeesPage = () => {
  const [employees, setEmployees] = useState<User[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState('')
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [fullName, setFullName] = useState('')
  const [creating, setCreating] = useState(false)

  const fetchEmployees = () => {
    setLoading(true)
    getUsers()
      .then((res: unknown) => { const r = res as { data: User[] }; setEmployees(r.data.filter((u) => u.role === 'employee')) })
      .catch(() => setError('Не удалось загрузить сотрудников'))
      .finally(() => setLoading(false))
  }

  useEffect(() => { fetchEmployees() }, [])

  const handleCreate = async (e: FormEvent) => {
    e.preventDefault()
    if (!email || !password || !fullName) return
    setCreating(true); setError('')
    try {
      const user = await createUser({ email, password, full_name: fullName })
      await patchUser(user.id, { role: 'employee' })
      setEmail(''); setPassword(''); setFullName(''); fetchEmployees()
    } catch (err) {
      setError((err && typeof err === 'object' && 'response' in err)
        ? ((err as { response?: { data?: { message?: string } } }).response?.data?.message || 'Ошибка при создании')
        : 'Ошибка при создании')
    } finally { setCreating(false) }
  }

  const handleDelete = async (user: User) => {
    if (!confirm(`Удалить сотрудника ${user.email}?`)) return
    try { setError(''); await deleteUser(user.id); fetchEmployees() }
    catch (err) {
      setError((err && typeof err === 'object' && 'response' in err)
        ? ((err as { response?: { data?: { message?: string } } }).response?.data?.message || 'Ошибка при удалении')
        : 'Ошибка при удалении')
    }
  }

  return (
    <AdminLayout title="СОТРУДНИКИ">
      {error && <div className="mb-4 px-3 py-2 bg-red-50 border border-red-200 text-red-600 text-[10px] font-bold uppercase rounded-[4px]">{error}</div>}
      <div className="y2k-box-thick mb-5">
        <h2 className="y2k-subtitle text-[11px] mb-3">ДОБАВИТЬ СОТРУДНИКА</h2>
        <form onSubmit={handleCreate} className="grid grid-cols-1 md:grid-cols-4 gap-4 items-end">
          <Input label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          <Input label="Пароль" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          <Input label="Имя" value={fullName} onChange={(e) => setFullName(e.target.value)} required />
          <Button type="submit" loading={creating}>СОЗДАТЬ</Button>
        </form>
      </div>
      {loading ? <div className="y2k-loading">Загрузка...</div> : (
        <div className="y2k-box overflow-hidden p-0">
          <table className="y2k-table">
            <thead><tr><th>EMAIL</th><th>ИМЯ</th><th className="text-center">ДЕЙСТВИЯ</th></tr></thead>
            <tbody>
              {employees.map((user) => (
                <tr key={user.id}>
                  <td className="font-bold text-gray-700">{user.email}</td>
                  <td className="text-gray-500">{user.full_name}</td>
                  <td className="text-center">
                    <button className="text-gray-400 hover:text-red-500" onClick={() => handleDelete(user)}><Trash2 size={14} /></button>
                  </td>
                </tr>
              ))}
              {employees.length === 0 && <tr><td colSpan={3} className="text-center py-10 text-gray-400">Нет сотрудников</td></tr>}
            </tbody>
          </table>
        </div>
      )}
    </AdminLayout>
  )
}
