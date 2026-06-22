interface Props {
  rating: number
  max?: number
  size?: 'sm' | 'md'
}

export const RatingDisplay = ({ rating, max = 5, size = 'sm' }: Props) => {
  const barCount = max
  const bars: { filled: boolean; partial: boolean; pct: number }[] = Array.from({ length: barCount }, (_, i) => {
    const filled = rating >= i + 1
    const partial = !filled && rating > i && rating < i + 1
    return { filled: filled || partial, partial, pct: partial ? (rating - i) * 100 : 0 }
  })

  const w = size === 'sm' ? 14 : 20
  const gap = size === 'sm' ? 2 : 3

  return (
    <div className="flex items-end" style={{ gap }}>
      {bars.map((bar, i) => {
        const height = size === 'sm' ? `${8 + (i + 1) * 3}px` : `${12 + (i + 1) * 5}px`
        const fillColor = i < 3 ? '#ff8a00' : '#ff6b00'
        return (
          <div
            key={i}
            style={{ width: w, height, backgroundColor: '#ddd', position: 'relative' }}
          >
            <div
              style={{
                width: '100%',
                height: bar.filled ? '100%' : 0,
                backgroundColor: fillColor,
                clipPath: bar.partial ? `inset(${100 - bar.pct}% 0 0 0)` : undefined,
              }}
            />
          </div>
        )
      })}
    </div>
  )
}
