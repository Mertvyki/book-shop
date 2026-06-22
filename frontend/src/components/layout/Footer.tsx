import { Link } from 'react-router-dom'

export const Footer = () => (
  <footer className="bg-white border-t-2 border-gray-250 mt-8">
    <div className="max-w-[1000px] mx-auto px-4 py-6">
      <div className="flex justify-between text-[10px]">
        <div>
          <p className="font-bold text-gray-600 uppercase mb-2 tracking-[2px] text-[11px]">BOOK<span className="text-[#ff8a00]">ZONE</span></p>
          <p className="text-gray-400 leading-5">
            Крупнейший онлайн-магазин книг.<br />
            Тысячи наименований. Лучшие авторы.<br />
            Новинки каждую неделю.
          </p>
        </div>
        <div>
          <p className="font-bold text-gray-500 uppercase mb-2 text-[10px] tracking-wide">Информация</p>
          <div className="space-y-1">
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">О нас</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Контакты</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Доставка</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Конфиденциальность</Link>
          </div>
        </div>
        <div>
          <p className="font-bold text-gray-500 uppercase mb-2 text-[10px] tracking-wide">Поддержка</p>
          <div className="space-y-1">
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">ЧаВо</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Партнёры</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Возврат</Link>
            <Link to="/" className="block text-gray-400 hover:text-[#ff8a00] no-underline">Помощь</Link>
          </div>
        </div>
      </div>
      <hr className="y2k-divider mt-5" />
      <p className="text-[9px] text-gray-400 text-center tracking-wider">
        &copy; 2003 BOOKZONE Corporation. Все права защищены.
      </p>
    </div>
  </footer>
)
