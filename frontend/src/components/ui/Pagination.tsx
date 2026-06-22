import { Button } from './Button'

interface Props {
  page: number
  totalPages: number
  total: number
  onPageChange: (page: number) => void
}

export const Pagination = ({ page, totalPages, total, onPageChange }: Props) => {
  if (totalPages <= 1) return null

  return (
    <div className="flex items-center justify-between mt-5 px-1">
      <p className="text-[11px] text-gray-400 uppercase tracking-wide">
        Всего: <span className="font-bold text-gray-600">{total}</span>
      </p>
      <div className="flex items-center gap-2">
        <Button variant="secondary" size="sm" disabled={page <= 1} onClick={() => onPageChange(page - 1)}>
          &laquo;
        </Button>
        <span className="text-[11px] text-gray-600 font-bold px-2">
          {page} / {totalPages}
        </span>
        <Button variant="secondary" size="sm" disabled={page >= totalPages} onClick={() => onPageChange(page + 1)}>
          &raquo;
        </Button>
      </div>
    </div>
  )
}
