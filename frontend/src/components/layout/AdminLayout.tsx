import { Link, useLocation } from 'react-router-dom'

const links = [
  { to: '/admin/books', label: 'КНИГИ' },
  { to: '/admin/orders', label: 'ЗАКАЗЫ' },
  { to: '/admin/authors', label: 'АВТОРЫ' },
  { to: '/admin/categories', label: 'КАТЕГОРИИ' },
  { to: '/admin/publishers', label: 'ИЗДАТЕЛИ' },
  { to: '/admin/employees', label: 'СОТРУДНИКИ' },
]

export const AdminLayout = ({ children, title }: { children: React.ReactNode; title: string }) => {
  const location = useLocation()

  return (
    <div>
      <div className="y2k-admin-topbar flex items-center justify-between mb-5">
        <div>
          <span className="text-[14px] font-bold tracking-[3px] text-[#003366]">BOOK<span className="text-[#ff8a00]">ZONE</span></span>
          <span className="text-[10px] text-gray-500 ml-3 uppercase tracking-wider font-bold">Панель управления</span>
        </div>
        <div className="flex items-center gap-2">
          <Link to="/" className="y2k-btn-xp text-[9px]">НА САЙТ</Link>
        </div>
      </div>

      <div className="flex gap-5 items-start">
        <div className="y2k-admin-sidebar shrink-0">
          <div className="px-3 py-2 border-b border-gray-300 text-[9px] font-bold text-gray-500 uppercase tracking-wider bg-gray-100/50">
            Навигация
          </div>
          {links.map((link) => {
            const isActive = location.pathname === link.to
            return (
              <Link
                key={link.to}
                to={link.to}
                className={`y2k-admin-sidebar-item ${isActive ? 'y2k-admin-sidebar-item-active' : ''}`}
              >
                {link.label}
              </Link>
            )
          })}
        </div>

        <div className="flex-1 min-w-0">
          <h1 className="y2k-title text-[15px] mb-4">{title}</h1>
          {children}
        </div>
      </div>
    </div>
  )
}
