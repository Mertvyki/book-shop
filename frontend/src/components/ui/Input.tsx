import type { InputHTMLAttributes } from 'react'

interface Props extends InputHTMLAttributes<HTMLInputElement> {
  label?: string
  error?: string
}

export const Input = ({ label, error, className = '', ...props }: Props) => (
  <div>
    {label && <label className="y2k-label">{label}</label>}
    <input
      className={`y2k-input ${error ? 'border-red-500' : ''} ${className}`}
      {...props}
    />
    {error && <p className="text-[11px] text-red-600 mt-1">{error}</p>}
  </div>
)
