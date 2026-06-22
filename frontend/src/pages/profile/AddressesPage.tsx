import { useEffect, useState, type FormEvent } from 'react'
import { useAuth } from '../../contexts/AuthContext'
import { getAddresses, createAddress, deleteAddress, patchAddress } from '../../api/addresses'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Modal } from '../../components/ui/Modal'
import { Plus, Trash2, Star, Edit2 } from 'lucide-react'
import type { Address } from '../../types'

const emptyForm = { street_address: '', city: '', postal_code: '', country: 'Россия', is_default: false }

export const AddressesPage = () => {
  const { user } = useAuth()
  const [addresses, setAddresses] = useState<Address[]>([])
  const [loading, setLoading] = useState(true)
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Address | null>(null)
  const [form, setForm] = useState(emptyForm)
  const [saving, setSaving] = useState(false)

  const fetch = () => {
    if (!user) return
    setLoading(true)
    getAddresses(user.id).then(setAddresses).finally(() => setLoading(false))
  }

  useEffect(() => { fetch() }, [user])

  const resetForm = () => { setForm(emptyForm); setEditing(null) }

  const openCreate = () => { resetForm(); setShowModal(true) }

  const openEdit = (addr: Address) => {
    setEditing(addr)
    setForm({
      street_address: addr.street_address,
      city: addr.city,
      postal_code: addr.postal_code,
      country: addr.country,
      is_default: addr.is_default,
    })
    setShowModal(true)
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!user) return
    setSaving(true)
    try {
      if (editing) {
        await patchAddress(user.id, editing.id, {
          street_address: { Value: form.street_address, Set: true },
          city: { Value: form.city, Set: true },
          postal_code: { Value: form.postal_code, Set: true },
          country: { Value: form.country, Set: true },
          is_default: { Value: form.is_default, Set: true },
        })
      } else {
        await createAddress(user.id, form)
      }
      setShowModal(false); resetForm(); fetch()
    } finally { setSaving(false) }
  }

  const handleDelete = async (id: number) => {
    if (!user) return
    if (!confirm('Удалить адрес?')) return
    await deleteAddress(user.id, id)
    fetch()
  }

  const setDefault = async (addr: Address) => {
    if (!user) return
    await patchAddress(user.id, addr.id, { is_default: { Value: true, Set: true } })
    fetch()
  }

  if (loading) return <div className="y2k-loading">Загрузка...</div>

  return (
    <div className="max-w-[500px] mx-auto">
      <div className="flex items-center justify-between mb-5">
        <h1 className="y2k-title">Адреса доставки</h1>
        <Button size="sm" onClick={openCreate}><Plus size={14} /> Добавить</Button>
      </div>

      {addresses.length === 0 ? (
        <p className="y2k-loading">Нет сохранённых адресов</p>
      ) : (
        <div className="space-y-3">
          {addresses.map((addr) => (
            <div key={addr.id} className="y2k-box flex items-start justify-between">
              <div>
                <div className="flex items-center gap-2">
                  <p className="text-[12px] font-bold text-gray-700">{addr.street_address}</p>
                  {addr.is_default && <Star size={12} className="text-[#ff8a00] fill-[#ff8a00]" />}
                </div>
                <p className="text-[11px] text-gray-400">{addr.city}, {addr.postal_code}</p>
                <p className="text-[11px] text-gray-400">{addr.country}</p>
              </div>
              <div className="flex items-center gap-2">
                <button className="text-gray-300 hover:text-[#ff8a00]" onClick={() => openEdit(addr)} title="Редактировать"><Edit2 size={14} /></button>
                {!addr.is_default && (
                  <button className="text-gray-300 hover:text-[#ff8a00]" onClick={() => setDefault(addr)} title="Сделать основным"><Star size={14} /></button>
                )}
                <button className="text-gray-300 hover:text-red-500" onClick={() => handleDelete(addr.id)}><Trash2 size={14} /></button>
              </div>
            </div>
          ))}
        </div>
      )}

      <Modal open={showModal} onClose={() => { setShowModal(false); resetForm() }} title={editing ? 'Редактировать адрес' : 'Новый адрес'}>
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Улица, дом" value={form.street_address} onChange={(e) => setForm({ ...form, street_address: e.target.value })} required />
          <Input label="Город" value={form.city} onChange={(e) => setForm({ ...form, city: e.target.value })} required />
          <Input label="Почтовый индекс" value={form.postal_code} onChange={(e) => setForm({ ...form, postal_code: e.target.value })} required />
          <Input label="Страна" value={form.country} onChange={(e) => setForm({ ...form, country: e.target.value })} />
          <label className="flex items-center gap-2 text-[12px] text-gray-600">
            <input type="checkbox" checked={form.is_default} onChange={(e) => setForm({ ...form, is_default: e.target.checked })} />
            Основной адрес
          </label>
          <Button type="submit" className="w-full" loading={saving}>{editing ? 'Сохранить' : 'Создать'}</Button>
        </form>
      </Modal>
    </div>
  )
}
