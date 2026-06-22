import type { ButtonHTMLAttributes, ReactNode } from 'react'

interface Props extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'danger' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  loading?: boolean
  children: ReactNode
}

const variants: Record<string, string> = {
  primary: 'y2k-btn',
  secondary: 'y2k-btn y2k-btn-secondary',
  danger: 'y2k-btn y2k-btn-danger',
  ghost: 'y2k-btn-ghost',
}

const sizes: Record<string, string> = {
  sm: 'text-[11px] px-3 py-1',
  md: 'text-[13px] px-4 py-1.5',
  lg: 'text-[14px] px-6 py-2',
}

export const Button = ({
  variant = 'primary',
  size = 'md',
  loading,
  children,
  className = '',
  disabled,
  ...props
}: Props) => (
  <button
    className={`${variants[variant]} ${sizes[size]} ${className}`}
    disabled={disabled || loading}
    {...props}
  >
    {loading && (
      <svg className="animate-spin h-4 w-4" fill="none" viewBox="0 0 24 24">
        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4z" />
      </svg>
    )}
    {children}
  </button>
)
