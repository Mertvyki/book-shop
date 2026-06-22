import { useEffect, useState, type FormEvent } from 'react'
import { getAuthors, createAuthor, patchAuthor, deleteAuthor } from '../../api/books'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Modal } from '../../components/ui/Modal'
import { Plus, Edit2, Trash2 } from 'lucide-react'
import type { Author } from '../../types'

export const AdminAuthorsPage = () => {
  const [authors, setAuthors] = useState<Author[]>([])
  const [loading, setLoading] = useState(true)
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Author | null>(null)
  const [form, setForm] = useState({ name: '', bio: '', birth_year: '' })
  const [saving, setSaving] = useState(false)

  const fetch = () => { setLoading(true); getAuthors().then(setAuthors).finally(() => setLoading(false)) }
  useEffect(() => { fetch() }, [])

  const resetForm = () => { setForm({ name: '', bio: '', birth_year: '' }); setEditing(null) }
  const openCreate = () => { resetForm(); setShowModal(true) }
  const openEdit = (author: Author) => { setEditing(author); setForm({ name: author.name, bio: author.bio || '', birth_year: author.birth_year ? String(author.birth_year) : '' }); setShowModal(true) }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!form.name.trim()) return
    setSaving(true)
    try {
      if (editing) await patchAuthor(editing.id, { name: form.name, bio: form.bio || null, birth_year: form.birth_year ? Number(form.birth_year) : null })
      else await createAuthor({ name: form.name, bio: form.bio || null, birth_year: form.birth_year ? Number(form.birth_year) : null })
      setShowModal(false); resetForm(); fetch()
    } finally { setSaving(false) }
  }

  const handleDelete = async (id: number) => { if (!confirm('Удалить автора?')) return; await deleteAuthor(id); fetch() }

  return (
    <AdminLayout title="АВТОРЫ">
      <div className="flex justify-end mb-3">
        <Button onClick={openCreate}><Plus size={14} /> ДОБАВИТЬ АВТОРА</Button>
      </div>
      {loading ? <div className="y2k-loading">Загрузка...</div> : (
        <div className="y2k-box overflow-hidden p-0">
          <table className="y2k-table">
            <thead><tr><th>ID</th><th>ИМЯ</th><th>ГОД РОЖДЕНИЯ</th><th className="text-center">ДЕЙСТВИЯ</th></tr></thead>
            <tbody>
              {authors.map((a) => (
                <tr key={a.id}>
                  <td className="text-gray-400">{a.id}</td>
                  <td className="font-bold text-gray-700">{a.name}</td>
                  <td className="text-gray-500">{a.birth_year || '—'}</td>
                  <td className="text-center">
                    <div className="flex items-center justify-center gap-2">
                      <button className="text-gray-400 hover:text-[#ff8a00]" onClick={() => openEdit(a)}><Edit2 size={14} /></button>
                      <button className="text-gray-400 hover:text-red-500" onClick={() => handleDelete(a.id)}><Trash2 size={14} /></button>
                    </div>
                  </td>
                </tr>
              ))}
              {authors.length === 0 && <tr><td colSpan={4} className="text-center py-10 text-gray-400">Нет авторов</td></tr>}
            </tbody>
          </table>
        </div>
      )}
      <Modal open={showModal} onClose={() => { setShowModal(false); resetForm() }} title={editing ? 'РЕДАКТИРОВАТЬ АВТОРА' : 'НОВЫЙ АВТОР'}>
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Имя" value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} required />
          <Input label="Биография" value={form.bio} onChange={(e) => setForm({ ...form, bio: e.target.value })} />
          <Input label="Год рождения" type="number" value={form.birth_year} onChange={(e) => setForm({ ...form, birth_year: e.target.value })} />
          <Button type="submit" className="w-full" loading={saving}>{editing ? 'СОХРАНИТЬ' : 'СОЗДАТЬ'}</Button>
        </form>
      </Modal>
    </AdminLayout>
  )
}
