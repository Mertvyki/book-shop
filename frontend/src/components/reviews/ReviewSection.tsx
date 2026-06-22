import { useEffect, useState } from 'react'
import { getBookReviews, getUserReview, upsertReview } from '../../api/reviews'
import { useAuth } from '../../contexts/AuthContext'
import { Button } from '../ui/Button'
import { RatingDisplay } from './RatingDisplay'
import { RatingInput } from './RatingInput'
import type { Review } from '../../types'

interface Props {
  bookId: number
  purchased: boolean
}

export const ReviewSection = ({ bookId, purchased }: Props) => {
  const { isAuthenticated } = useAuth()
  const [reviews, setReviews] = useState<Review[]>([])
  const [myReview, setMyReview] = useState<Review | null>(null)
  const [loading, setLoading] = useState(true)
  const [showForm, setShowForm] = useState(false)
  const [rating, setRating] = useState(5)
  const [title, setTitle] = useState('')
  const [body, setBody] = useState('')
  const [saving, setSaving] = useState(false)

  const loadReviews = () => {
    setLoading(true)
    getBookReviews(bookId).then(setReviews).catch(() => {}).finally(() => setLoading(false))
  }

  const loadMyReview = () => {
    if (!isAuthenticated) return
    getUserReview(bookId).then((r) => {
      setMyReview(r)
      setRating(r.rating)
      setTitle(r.title || '')
      setBody(r.body || '')
    }).catch(() => {})
  }

  useEffect(() => {
    loadReviews()
    loadMyReview()
  }, [bookId, isAuthenticated])

  const handleSubmit = async () => {
    setSaving(true)
    try {
      await upsertReview(bookId, { rating, title: title || null, body: body || null })
      setShowForm(false)
      loadReviews()
      loadMyReview()
    } catch {
      alert('Не удалось сохранить отзыв')
    } finally {
      setSaving(false)
    }
  }

  const handleEdit = () => {
    if (myReview) {
      setRating(myReview.rating)
      setTitle(myReview.title || '')
      setBody(myReview.body || '')
    }
    setShowForm(true)
  }

  return (
    <div className="mt-8 pt-6 border-t-2 border-gray-200">
      <h2 className="y2k-title text-[16px] mb-4">ОТЗЫВЫ</h2>

      {showForm && (
        <div className="y2k-box p-4 mb-5">
          <h3 className="y2k-subtitle text-[13px] mb-3">
            {myReview ? 'РЕДАКТИРОВАТЬ ОТЗЫВ' : 'НАПИСАТЬ ОТЗЫВ'}
          </h3>
          <div className="mb-3">
            <label className="text-[11px] text-gray-500 block mb-1">Оценка</label>
            <RatingInput value={rating} onChange={setRating} />
          </div>
          <div className="mb-3">
            <label className="text-[11px] text-gray-500 block mb-1">Заголовок</label>
            <input
              className="y2k-input w-full"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Краткое описание"
            />
          </div>
          <div className="mb-3">
            <label className="text-[11px] text-gray-500 block mb-1">Отзыв</label>
            <textarea
              className="y2k-input w-full h-24 resize-none"
              value={body}
              onChange={(e) => setBody(e.target.value)}
              placeholder="Поделитесь впечатлениями о книге..."
            />
          </div>
          <div className="flex gap-2">
            <Button size="sm" onClick={handleSubmit} loading={saving}>
              {myReview ? 'СОХРАНИТЬ' : 'ОТПРАВИТЬ'}
            </Button>
            <Button size="sm" variant="secondary" onClick={() => setShowForm(false)}>
              ОТМЕНА
            </Button>
          </div>
        </div>
      )}

      {isAuthenticated && purchased && !myReview && !showForm && (
        <Button size="sm" className="mb-4" onClick={() => { setRating(5); setTitle(''); setBody(''); setShowForm(true) }}>
          НАПИСАТЬ ОТЗЫВ
        </Button>
      )}

      {isAuthenticated && myReview && !showForm && (
        <div className="y2k-box p-3 mb-4">
          <div className="flex items-center justify-between mb-1">
            <span className="text-[11px] font-bold text-gray-600">ВАШ ОТЗЫВ</span>
            <Button size="sm" variant="ghost" onClick={handleEdit}>РЕДАКТИРОВАТЬ</Button>
          </div>
          <RatingDisplay rating={myReview.rating} />
          {myReview.title && <p className="text-[12px] font-bold mt-1">{myReview.title}</p>}
          {myReview.body && <p className="text-[11px] text-gray-600 mt-1">{myReview.body}</p>}
        </div>
      )}

      {loading ? (
        <p className="text-[11px] text-gray-400">Загрузка отзывов...</p>
      ) : reviews.length === 0 ? (
        <p className="text-[11px] text-gray-400">Пока нет отзывов</p>
      ) : (
        <div className="space-y-3">
          {reviews.map((review) => (
            <div key={review.id} className="y2k-box p-3">
              <div className="flex items-center gap-2 mb-1">
                <span className="text-[11px] font-bold text-gray-600">{review.user_name}</span>
                <span className="text-[10px] text-gray-400">
                  {new Date(review.created_at).toLocaleDateString('ru-RU')}
                </span>
              </div>
              <RatingDisplay rating={review.rating} />
              {review.title && <p className="text-[12px] font-bold mt-1">{review.title}</p>}
              {review.body && <p className="text-[11px] text-gray-600 mt-1">{review.body}</p>}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
