import { useEffect, useState, type FormEvent } from 'react'
import { getCategories, createCategory, patchCategory, deleteCategory } from '../../api/books'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Modal } from '../../components/ui/Modal'
import { Plus, Edit2, Trash2 } from 'lucide-react'
import type { Category } from '../../types'

export const AdminCategoriesPage = () => {
  const [categories, setCategories] = useState<Category[]>([])
  const [loading, setLoading] = useState(true)
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Category | null>(null)
  const [form, setForm] = useState({ name: '', slug: '', description: '' })
  const [saving, setSaving] = useState(false)

  const fetch = () => { setLoading(true); getCategories().then(setCategories).finally(() => setLoading(false)) }
  useEffect(() => { fetch() }, [])

  const resetForm = () => { setForm({ name: '', slug: '', description: '' }); setEditing(null) }
  const openCreate = () => { resetForm(); setShowModal(true) }
  const openEdit = (cat: Category) => { setEditing(cat); setForm({ name: cat.name, slug: cat.slug, description: cat.description || '' }); setShowModal(true) }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    if (!form.name.trim()) return
    setSaving(true)
    try {
      if (editing) await patchCategory(editing.id, { name: form.name, slug: form.slug || undefined, description: form.description || null })
      else await createCategory({ name: form.name, slug: form.slug || undefined, description: form.description || null })
      setShowModal(false); resetForm(); fetch()
    } finally { setSaving(false) }
  }

  const handleDelete = async (id: number) => { if (!confirm('Удалить категорию?')) return; await deleteCategory(id); fetch() }

  return (
    <AdminLayout title="КАТЕГОРИИ">
      <div className="flex justify-end mb-3">
        <Button onClick={openCreate}><Plus size={14} /> ДОБАВИТЬ КАТЕГОРИЮ</Button>
      </div>
      {loading ? <div className="y2k-loading">Загрузка...</div> : (
        <div className="y2k-box overflow-hidden p-0">
          <table className="y2k-table">
            <thead><tr><th>ID</th><th>НАЗВАНИЕ</th><th>SLUG</th><th className="text-center">ДЕЙСТВИЯ</th></tr></thead>
            <tbody>
              {categories.map((c) => (
                <tr key={c.id}>
                  <td className="text-gray-400">{c.id}</td>
                  <td className="font-bold text-gray-700">{c.name}</td>
                  <td className="text-gray-500">{c.slug}</td>
                  <td className="text-center">
                    <div className="flex items-center justify-center gap-2">
                      <button className="text-gray-400 hover:text-[#ff8a00]" onClick={() => openEdit(c)}><Edit2 size={14} /></button>
                      <button className="text-gray-400 hover:text-red-500" onClick={() => handleDelete(c.id)}><Trash2 size={14} /></button>
                    </div>
                  </td>
                </tr>
              ))}
              {categories.length === 0 && <tr><td colSpan={4} className="text-center py-10 text-gray-400">Нет категорий</td></tr>}
            </tbody>
          </table>
        </div>
      )}
      <Modal open={showModal} onClose={() => { setShowModal(false); resetForm() }} title={editing ? 'РЕДАКТИРОВАТЬ КАТЕГОРИЮ' : 'НОВАЯ КАТЕГОРИЯ'}>
        <form onSubmit={handleSubmit} className="space-y-4">
          <Input label="Название" value={form.name} onChange={(e) => setForm({ ...form, name: e.target.value })} required />
          <Input label="Slug" value={form.slug} onChange={(e) => setForm({ ...form, slug: e.target.value })} placeholder="оставьте пустым для автогенерации" />
          <Input label="Описание" value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} />
          <Button type="submit" className="w-full" loading={saving}>{editing ? 'СОХРАНИТЬ' : 'СОЗДАТЬ'}</Button>
        </form>
      </Modal>
    </AdminLayout>
  )
}
