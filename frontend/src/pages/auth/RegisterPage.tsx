import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { createUser } from '../../api/users'

export const RegisterPage = () => {
  const navigate = useNavigate()
  const [form, setForm] = useState({ email: '', password: '', full_name: '', phone_number: '' })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      await createUser(form)
      navigate('/login')
    } catch (err: unknown) {
      if (err && typeof err === 'object' && 'response' in err) {
        const axiosErr = err as { response?: { data?: { message?: string } } }
        setError(axiosErr.response?.data?.message || 'Ошибка регистрации')
      } else {
        setError('Ошибка регистрации')
      }
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-[70vh] flex items-center justify-center">
      <div className="w-full max-w-[400px] y2k-box-thick">
        <div className="text-center mb-6">
          <p className="y2k-title text-[18px] mb-1">BOOK<span className="text-[#ff8a00]">ZONE</span></p>
          <p className="text-[11px] text-gray-400 uppercase tracking-wide">Регистрация</p>
        </div>
        {error && <div className="mb-4 px-3 py-2 bg-red-50 border border-red-200 text-red-600 text-[11px] font-bold uppercase rounded-[4px]">{error}</div>}
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Полное имя" value={form.full_name} onChange={(e) => setForm({ ...form, full_name: e.target.value })} required />
          <Input label="Email" type="email" value={form.email} onChange={(e) => setForm({ ...form, email: e.target.value })} required />
          <Input label="Пароль" type="password" value={form.password} onChange={(e) => setForm({ ...form, password: e.target.value })} required />
          <Input label="Телефон" value={form.phone_number} onChange={(e) => setForm({ ...form, phone_number: e.target.value })} />
          <Button type="submit" className="w-full" loading={loading}>Зарегистрироваться</Button>
        </form>
        <p className="text-center text-[11px] text-gray-400 mt-5">
          Уже есть аккаунт?{' '}
          <Link to="/login" className="y2k-link">Войти</Link>
        </p>
      </div>
    </div>
  )
}
