import { useEffect, useState, type FormEvent } from 'react'
import { Link } from 'react-router-dom'
import { useAuth } from '../../contexts/AuthContext'
import { patchUser, getUser } from '../../api/users'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'

export const ProfilePage = () => {
  const { user } = useAuth()
  const [fullName, setFullName] = useState(user?.full_name || '')
  const [phone, setPhone] = useState(user?.phone_number || '')
  const [email, setEmail] = useState(user?.email || '')
  const [oldPassword, setOldPassword] = useState('')
  const [newPassword, setNewPassword] = useState('')
  const [saving, setSaving] = useState(false)
  const [loading, setLoading] = useState(true)
  const [profile, setProfile] = useState(user)

  useEffect(() => {
    if (!user) return
    getUser(user.id)
      .then((u) => {
        setProfile(u)
        setFullName(u.full_name)
        setPhone(u.phone_number || '')
        setEmail(u.email)
      })
      .finally(() => setLoading(false))
  }, [user])

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!user || !profile) return
    setSaving(true)
    try {
      const payload: Record<string, unknown> = {}
      if (fullName !== profile.full_name) payload.full_name = fullName
      if ((phone || null) !== profile.phone_number) payload.phone_number = phone || null
      if (email !== profile.email) payload.email = email
      if (oldPassword && newPassword) {
        payload.old_password = oldPassword
        payload.password = newPassword
      }
      if (Object.keys(payload).length === 0) return
      const updated = await patchUser(user.id, payload)
      localStorage.setItem('user', JSON.stringify(updated))
      setProfile(updated)
      setOldPassword('')
      setNewPassword('')
      alert('Профиль обновлён')
    } catch {
      alert('Ошибка при сохранении')
    } finally {
      setSaving(false)
    }
  }

  if (loading) return <div className="y2k-loading">Загрузка...</div>

  return (
    <div className="max-w-[500px] mx-auto">
      <h1 className="y2k-title mb-5">Профиль</h1>

      <div className="y2k-box-thick mb-5">
        <div className="mb-3">
          <span className="y2k-label">Email</span>
          <p className="text-[13px] font-bold text-gray-700">{profile?.email}</p>
        </div>
        <div>
          <span className="y2k-label">Роль</span>
          <p className="text-[13px] font-bold text-gray-700">{profile?.role === 'admin' ? 'Администратор' : profile?.role === 'employee' ? 'Сотрудник' : 'Пользователь'}</p>
        </div>
      </div>

      <form onSubmit={handleSubmit} className="y2k-box-thick space-y-4">
        <h2 className="y2k-subtitle text-[12px]">Редактировать профиль</h2>
        <Input label="Полное имя" value={fullName} onChange={(e) => setFullName(e.target.value)} required />
        <Input label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} />
        <Input label="Телефон" value={phone} onChange={(e) => setPhone(e.target.value)} />
        <hr className="y2k-divider" />
        <h3 className="text-[11px] font-bold text-gray-500 uppercase tracking-wide">Смена пароля</h3>
        <Input label="Старый пароль" type="password" value={oldPassword} onChange={(e) => setOldPassword(e.target.value)} />
        <Input label="Новый пароль" type="password" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} />
        <Button type="submit" loading={saving}>Сохранить</Button>
      </form>

      <div className="mt-4">
        <Link to="/profile/addresses" className="y2k-link text-[12px]">
          Управление адресами &rarr;
        </Link>
      </div>
    </div>
  )
}
