import { useEffect, useState, type FormEvent } from 'react'
import { getBooks, getAuthors, getCategories, getPublishers, createBook, deleteBook, patchBook, createAuthor, createCategory, createPublisher } from '../../api/books'
import { AdminLayout } from '../../components/layout/AdminLayout'
import { Button } from '../../components/ui/Button'
import { Input } from '../../components/ui/Input'
import { Modal } from '../../components/ui/Modal'
import { Pagination } from '../../components/ui/Pagination'
import type { Book, Author, Category, Publisher } from '../../types'
import { Plus, Edit2, Trash2 } from 'lucide-react'

const emptyForm = {
  title: '', price: '', book_type: 'physical', description: '', isbn: '',
  stock_quantity: '', publisher_id: '', author_ids: [] as number[], category_ids: [] as number[],
}

export const AdminBooksPage = () => {
  const [books, setBooks] = useState<Book[]>([])
  const [total, setTotal] = useState(0); const [totalPages, setTotalPages] = useState(0)
  const [page, setPage] = useState(1)
  const [authors, setAuthors] = useState<Author[]>([])
  const [categories, setCategories] = useState<Category[]>([])
  const [publishers, setPublishers] = useState<Publisher[]>([])
  const [loading, setLoading] = useState(true)
  const [showModal, setShowModal] = useState(false)
  const [editing, setEditing] = useState<Book | null>(null)
  const [form, setForm] = useState(emptyForm)
  const [cover, setCover] = useState<File | null>(null)
  const [bookFile, setBookFile] = useState<File | null>(null)
  const [saving, setSaving] = useState(false)
  const [showQuickAuthor, setShowQuickAuthor] = useState(false)
  const [quickAuthorName, setQuickAuthorName] = useState('')
  const [showQuickCategory, setShowQuickCategory] = useState(false)
  const [quickCategoryName, setQuickCategoryName] = useState('')
  const [showQuickPublisher, setShowQuickPublisher] = useState(false)
  const [quickPublisherName, setQuickPublisherName] = useState('')

  const fetchBooks = () => {
    setLoading(true)
    getBooks({ page, limit: 10 })
      .then((res) => { setBooks(res.data); setTotal(res.total); setTotalPages(res.total_pages) })
      .finally(() => setLoading(false))
  }

  useEffect(() => { fetchBooks(); getAuthors().then(setAuthors); getCategories().then(setCategories); getPublishers().then(setPublishers) }, [page])

  const resetForm = () => { setForm(emptyForm); setCover(null); setBookFile(null); setEditing(null) }
  const openCreate = () => { resetForm(); setShowModal(true) }

  const openEdit = (book: Book) => {
    setEditing(book)
    setForm({
      title: book.title, price: String(book.price), book_type: book.book_type,
      description: book.description || '', isbn: book.isbn || '',
      stock_quantity: book.stock_quantity !== null ? String(book.stock_quantity) : '',
      publisher_id: book.publisher?.id ? String(book.publisher.id) : '',
      author_ids: book.authors.map((a) => a.id), category_ids: book.categories.map((c) => c.id),
    })
    setShowModal(true)
  }

  const setAuthorIDs = (ids: number[]) => { setForm((f) => ({ ...f, author_ids: ids })) }
  const setCategoryIDs = (ids: number[]) => { setForm((f) => ({ ...f, category_ids: ids })) }

  const handleQuickAuthor = async () => {
    if (!quickAuthorName.trim()) return
    try { await createAuthor({ name: quickAuthorName }); setQuickAuthorName(''); setShowQuickAuthor(false); getAuthors().then(setAuthors) } catch { alert('Ошибка при создании автора') }
  }
  const handleQuickCategory = async () => {
    if (!quickCategoryName.trim()) return
    try { await createCategory({ name: quickCategoryName }); setQuickCategoryName(''); setShowQuickCategory(false); getCategories().then(setCategories) } catch { alert('Ошибка при создании категории') }
  }
  const handleQuickPublisher = async () => {
    if (!quickPublisherName.trim()) return
    try { await createPublisher(quickPublisherName); setQuickPublisherName(''); setShowQuickPublisher(false); getPublishers().then(setPublishers) } catch { alert('Ошибка при создании издателя') }
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setSaving(true)
    try {
      if (editing) {
        const payload: Record<string, unknown> = {
          title: form.title, price: Number(form.price),
          book_type: form.book_type,
        }
        if (form.description) payload.description = form.description
        if (form.isbn) payload.isbn = form.isbn
        if (form.publisher_id) payload.publisher_id = Number(form.publisher_id)
        if (form.stock_quantity) payload.stock_quantity = Number(form.stock_quantity)
        payload.author_ids = form.author_ids; payload.category_ids = form.category_ids
        await patchBook(editing.id, payload, cover || undefined, bookFile || undefined)
      } else {
        const fd = new FormData()
        fd.append('title', form.title); fd.append('price', form.price); fd.append('book_type', form.book_type)
        if (form.description) fd.append('description', form.description)
        if (form.isbn) fd.append('isbn', form.isbn)
        if (form.stock_quantity) fd.append('stock_quantity', form.stock_quantity)
        if (form.publisher_id) fd.append('publisher_id', form.publisher_id)
        fd.append('author_ids', form.author_ids.join(',')); fd.append('category_ids', form.category_ids.join(','))
        if (cover) fd.append('cover_image', cover)
        if (bookFile) fd.append('book_file', bookFile)
        await createBook(fd)
      }
      setShowModal(false); resetForm(); fetchBooks()
    } finally { setSaving(false) }
  }

  const handleDelete = async (id: number) => { if (!confirm('Удалить книгу?')) return; await deleteBook(id); fetchBooks() }

  return (
    <AdminLayout title="КНИГИ">
      <div className="flex justify-end mb-3">
        <Button onClick={openCreate}><Plus size={14} /> ДОБАВИТЬ КНИГУ</Button>
      </div>
      {loading ? <div className="y2k-loading">Загрузка...</div> : (
        <>
          <div className="y2k-box overflow-hidden p-0">
            <table className="y2k-table">
              <thead><tr><th>НАЗВАНИЕ</th><th>ТИП</th><th className="text-right">ЦЕНА</th><th className="text-center">ДЕЙСТВИЯ</th></tr></thead>
              <tbody>
                {books.map((book) => (
                  <tr key={book.id}>
                    <td className="font-bold text-gray-700">{book.title}</td>
                    <td className="text-gray-500">{book.book_type === 'digital' ? 'ЦИФРОВАЯ' : 'ПЕЧАТНАЯ'}</td>
                    <td className="text-right y2k-price text-[13px]">{book.price.toLocaleString('ru-RU')} ₽</td>
                    <td className="text-center">
                      <div className="flex items-center justify-center gap-2">
                        <button className="text-gray-400 hover:text-[#ff8a00]" onClick={() => openEdit(book)}><Edit2 size={14} /></button>
                        <button className="text-gray-400 hover:text-red-500" onClick={() => handleDelete(book.id)}><Trash2 size={14} /></button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
          <Pagination page={page} totalPages={totalPages} total={total} onPageChange={setPage} />
        </>
      )}

      <Modal open={showModal} onClose={() => { setShowModal(false); resetForm() }} title={editing ? 'РЕДАКТИРОВАТЬ КНИГУ' : 'НОВАЯ КНИГА'}>
        <form onSubmit={handleSubmit} className="space-y-4 max-h-[70vh] overflow-y-auto">
          <Input label="Название" value={form.title} onChange={(e) => setForm({ ...form, title: e.target.value })} required />
          <Input label="Цена" type="number" value={form.price} onChange={(e) => setForm({ ...form, price: e.target.value })} required />
          <div><label className="y2k-label">Тип</label><select className="y2k-select" value={form.book_type} onChange={(e) => setForm({ ...form, book_type: e.target.value })}><option value="physical">ПЕЧАТНАЯ</option><option value="digital">ЦИФРОВАЯ</option></select></div>
          <Input label="Описание" value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} />
          <Input label="ISBN" value={form.isbn} onChange={(e) => setForm({ ...form, isbn: e.target.value })} />
          {form.book_type === 'physical' && <Input label="Количество на складе" type="number" value={form.stock_quantity} onChange={(e) => setForm({ ...form, stock_quantity: e.target.value })} />}
          <div>
            <label className="y2k-label">Авторы</label>
            <select multiple className="y2k-select h-32" value={form.author_ids.map(String)} onChange={(e) => setAuthorIDs(Array.from(e.target.selectedOptions, (o) => Number(o.value)))}>
              {authors.length === 0 && <option disabled>Нет авторов</option>}
              {authors.map((a) => (<option key={a.id} value={a.id}>{a.name}</option>))}
            </select>
            <button type="button" className="y2k-link mt-1 block" onClick={() => setShowQuickAuthor(true)}>+ Быстро добавить автора</button>
          </div>
          <div>
            <label className="y2k-label">Категории</label>
            <select multiple className="y2k-select h-32" value={form.category_ids.map(String)} onChange={(e) => setCategoryIDs(Array.from(e.target.selectedOptions, (o) => Number(o.value)))}>
              {categories.length === 0 && <option disabled>Нет категорий</option>}
              {categories.map((c) => (<option key={c.id} value={c.id}>{c.name}</option>))}
            </select>
            <button type="button" className="y2k-link mt-1 block" onClick={() => setShowQuickCategory(true)}>+ Быстро добавить категорию</button>
          </div>
          <div>
            <label className="y2k-label">Издатель</label>
            <select multiple className="y2k-select h-32" value={form.publisher_id ? [form.publisher_id] : []} onChange={(e) => {
              const selected = Array.from(e.target.selectedOptions, (o) => o.value)
              setForm({ ...form, publisher_id: selected[selected.length - 1] || '' })
            }}>
              {publishers.length === 0 && <option disabled>Нет издателей</option>}
              {publishers.map((p) => (<option key={p.id} value={p.id}>{p.name}</option>))}
            </select>
            <button type="button" className="y2k-link mt-1 block" onClick={() => setShowQuickPublisher(true)}>+ Быстро добавить издателя</button>
          </div>
          <div>
            <label className="y2k-label">Обложка</label>
            <label className="y2k-file-btn">
              {cover ? cover.name : 'ВЫБРАТЬ ОБЛОЖКУ'}
              <input type="file" accept="image/*" onChange={(e) => setCover(e.target.files?.[0] || null)} />
            </label>
            {cover && <span className="text-[10px] text-gray-400 ml-2">{cover.name}</span>}
          </div>
          {form.book_type === 'digital' && (
            <div>
              <label className="y2k-label">Файл книги</label>
              <label className="y2k-file-btn">
                {bookFile ? bookFile.name : 'ВЫБРАТЬ ФАЙЛ'}
                <input type="file" onChange={(e) => setBookFile(e.target.files?.[0] || null)} />
              </label>
              {bookFile && <span className="text-[10px] text-gray-400 ml-2">{bookFile.name}</span>}
            </div>
          )}
          <Button type="submit" className="w-full" loading={saving}>{editing ? 'СОХРАНИТЬ' : 'СОЗДАТЬ'}</Button>
        </form>
      </Modal>

      <Modal open={showQuickAuthor} onClose={() => setShowQuickAuthor(false)} title="БЫСТРОЕ ДОБАВЛЕНИЕ АВТОРА">
        <form onSubmit={(e) => { e.preventDefault(); handleQuickAuthor() }} className="space-y-4">
          <Input label="Имя автора" value={quickAuthorName} onChange={(e) => setQuickAuthorName(e.target.value)} required />
          <div className="flex gap-2"><Button type="submit" loading={saving}>СОЗДАТЬ</Button><Button variant="secondary" type="button" onClick={() => setShowQuickAuthor(false)}>ОТМЕНА</Button></div>
        </form>
      </Modal>
      <Modal open={showQuickCategory} onClose={() => setShowQuickCategory(false)} title="БЫСТРОЕ ДОБАВЛЕНИЕ КАТЕГОРИИ">
        <form onSubmit={(e) => { e.preventDefault(); handleQuickCategory() }} className="space-y-4">
          <Input label="Название категории" value={quickCategoryName} onChange={(e) => setQuickCategoryName(e.target.value)} required />
          <div className="flex gap-2"><Button type="submit" loading={saving}>СОЗДАТЬ</Button><Button variant="secondary" type="button" onClick={() => setShowQuickCategory(false)}>ОТМЕНА</Button></div>
        </form>
      </Modal>
      <Modal open={showQuickPublisher} onClose={() => setShowQuickPublisher(false)} title="БЫСТРОЕ ДОБАВЛЕНИЕ ИЗДАТЕЛЯ">
        <form onSubmit={(e) => { e.preventDefault(); handleQuickPublisher() }} className="space-y-4">
          <Input label="Название издателя" value={quickPublisherName} onChange={(e) => setQuickPublisherName(e.target.value)} required />
          <div className="flex gap-2"><Button type="submit" loading={saving}>СОЗДАТЬ</Button><Button variant="secondary" type="button" onClick={() => setShowQuickPublisher(false)}>ОТМЕНА</Button></div>
        </form>
      </Modal>
    </AdminLayout>
  )
}
