import { type ReactNode, useEffect } from 'react'
interface Props {
  open: boolean
  onClose: () => void
  title: string
  children: ReactNode
}

export const Modal = ({ open, onClose, title, children }: Props) => {
  useEffect(() => {
    if (open) {
      document.body.style.overflow = 'hidden'
    } else {
      document.body.style.overflow = ''
    }
    return () => { document.body.style.overflow = '' }
  }, [open])

  if (!open) return null

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center">
      <div className="fixed inset-0 bg-black/40" onClick={onClose} />
      <div className="relative z-10 w-full max-w-lg mx-4 bg-white border-2 border-gray-250 rounded-[10px] shadow-lg">
        <div className="flex items-center justify-between px-5 py-3 border-b-2 border-gray-200 bg-gray-100 rounded-t-[9px]">
          <h2 className="y2k-subtitle text-[13px]">{title}</h2>
          <button onClick={onClose} className="text-gray-400 hover:text-red-600 font-bold text-lg leading-none">&times;</button>
        </div>
        <div className="px-5 py-4">{children}</div>
      </div>
    </div>
  )
}
