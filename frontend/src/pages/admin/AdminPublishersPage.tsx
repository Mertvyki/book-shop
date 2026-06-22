import { useEffect, useState, type FormEvent } from 'react'
import { getPublishers, createPublisher, patchPublisher, deletePublisher } from '../../api/books'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Modal } from '../../components/ui/Modal'
import { Plus, Edit2, Trash2 } from 'lucide-react'
import type { Publisher } from '../../types'

export const AdminPublishersPage = () => {
  const [publishers, setPublishers] = useState<Publisher[]>([])
  const [loading, setLoading] = useState(true)
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Publisher | null>(null)
  const [name, setName] = useState('')
  const [saving, setSaving] = useState(false)

  const fetch = () => { setLoading(true); getPublishers().then(setPublishers).finally(() => setLoading(false)) }
  useEffect(() => { fetch() }, [])

  const resetForm = () => { setName(''); setEditing(null) }
  const openCreate = () => { resetForm(); setShowModal(true) }
  const openEdit = (p: Publisher) => { setEditing(p); setName(p.name); setShowModal(true) }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!name.trim()) return
    setSaving(true)
    try {
      if (editing) await patchPublisher(editing.id, name)
      else await createPublisher(name)
      setShowModal(false); resetForm(); fetch()
    } finally { setSaving(false) }
  }

  const handleDelete = async (id: number) => { if (!confirm('Удалить издателя?')) return; await deletePublisher(id); fetch() }

  return (
    <AdminLayout title="ИЗДАТЕЛИ">
      <div className="flex justify-end mb-3">
        <Button onClick={openCreate}><Plus size={14} /> ДОБАВИТЬ ИЗДАТЕЛЯ</Button>
      </div>
      {loading ? <div className="y2k-loading">Загрузка...</div> : (
        <div className="y2k-box overflow-hidden p-0">
          <table className="y2k-table">
            <thead><tr><th>ID</th><th>НАЗВАНИЕ</th><th className="text-center">ДЕЙСТВИЯ</th></tr></thead>
            <tbody>
              {publishers.map((p) => (
                <tr key={p.id}>
                  <td className="text-gray-400">{p.id}</td>
                  <td className="font-bold text-gray-700">{p.name}</td>
                  <td className="text-center">
                    <div className="flex items-center justify-center gap-2">
                      <button className="text-gray-400 hover:text-[#ff8a00]" onClick={() => openEdit(p)}><Edit2 size={14} /></button>
                      <button className="text-gray-400 hover:text-red-500" onClick={() => handleDelete(p.id)}><Trash2 size={14} /></button>
                    </div>
                  </td>
                </tr>
              ))}
              {publishers.length === 0 && <tr><td colSpan={3} className="text-center py-10 text-gray-400">Нет издателей</td></tr>}
            </tbody>
          </table>
        </div>
      )}
      <Modal open={showModal} onClose={() => { setShowModal(false); resetForm() }} title={editing ? 'РЕДАКТИРОВАТЬ ИЗДАТЕЛЯ' : 'НОВЫЙ ИЗДАТЕЛЬ'}>
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Название" value={name} onChange={(e) => setName(e.target.value)} required />
          <Button type="submit" className="w-full" loading={saving}>{editing ? 'СОХРАНИТЬ' : 'СОЗДАТЬ'}</Button>
        </form>
      </Modal>
    </AdminLayout>
  )
}
