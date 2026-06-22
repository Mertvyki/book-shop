interface Props {
  value: number
  onChange: (rating: number) => void
}

export const RatingInput = ({ value, onChange }: Props) => {
  const bars = [1, 2, 3, 4, 5]

  return (
    <div className="flex items-end gap-[3px] cursor-pointer" onClick={(e) => e.stopPropagation()}>
      {bars.map((i) => {
        const height = `${12 + i * 5}px`
        const filled = value >= i
        return (
          <div
            key={i}
            style={{ width: 20, height, backgroundColor: '#ddd', position: 'relative' }}
            onMouseEnter={() => onChange(i)}
          >
            <div
              style={{
                width: '100%',
                height: filled ? '100%' : 0,
                backgroundColor: i < 3 ? '#ff8a00' : '#ff6b00',
              }}
            />
          </div>
        )
      })}
    </div>
  )
}
