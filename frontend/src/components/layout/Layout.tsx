import { Outlet } from 'react-router-dom'
import { Header } from './Header'
import { Footer } from './Footer'

export const Layout = () => (
  <div className="min-h-screen flex flex-col bg-[#e5e5e5]">
    <Header />
    <main className="flex-1 max-w-[1000px] w-full mx-auto px-4 py-5">
      <Outlet />
    </main>
    <Footer />
  </div>
)
