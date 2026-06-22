import { useState, type FormEvent } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../../contexts/AuthContext'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'

export const LoginPage = () => {
  const { login } = useAuth()
  const navigate = useNavigate()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      await login(email, password)
      navigate('/')
    } catch {
      setError('Неверный email или пароль')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-[70vh] flex items-center justify-center">
      <div className="w-full max-w-[400px] y2k-box-thick">
        <div className="text-center mb-6">
          <p className="y2k-title text-[18px] mb-1">BOOK<span className="text-[#ff8a00]">ZONE</span></p>
          <p className="text-[11px] text-gray-400 uppercase tracking-wide">Вход в аккаунт</p>
        </div>
        {error && <div className="mb-4 px-3 py-2 bg-red-50 border border-red-200 text-red-600 text-[11px] font-bold uppercase rounded-[4px]">{error}</div>}
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
          <Input label="Пароль" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
          <Button type="submit" className="w-full" loading={loading}>Войти</Button>
        </form>
        <p className="text-center text-[11px] text-gray-400 mt-5">
          Нет аккаунта?{' '}
          <Link to="/register" className="y2k-link">Зарегистрироваться</Link>
        </p>
      </div>
    </div>
  )
}
