import { useAuth } from '../../context/AuthContext';

export default function Navbar() {
  const { user, logout } = useAuth();

  return (
    <nav className="h-16 border-b border-[#30363d] bg-[#161b22] flex items-center justify-between px-8 z-10 shrink-0">
      {/* Left Section: Breadcrumbs/Logo */}
      <div className="flex items-center gap-4">
        <div className="text-gray-400 font-medium text-sm tracking-tight">
          Dashboard / <span className="text-white font-semibold">All Files</span>
        </div>
      </div>

      {/* Right Section: Search & Profile */}
      <div className="flex items-center gap-6">
        {/* Search Bar */}
        <div className="hidden md:block">
          <input 
            type="text" 
            placeholder="Search your vault..." 
            className="bg-[#0d1117] border border-[#30363d] rounded-md px-4 py-1.5 text-sm text-gray-300 focus:border-blue-500 focus:ring-1 focus:ring-blue-500 outline-none transition-all w-64"
          />
        </div>

        {/* User Profile & Logout */}
        <div className="flex items-center gap-4 border-l border-[#30363d] pl-6">
          <div className="text-right hidden sm:block">
            <p className="text-sm font-bold text-white leading-none capitalize">
              {user?.username || 'User'}
            </p>
            <p className="text-[11px] text-gray-500 mt-1 font-mono">
              {user?.email}
            </p>
          </div>

          <button 
            onClick={logout}
            className="flex items-center gap-2 px-3 py-1.5 rounded-md text-red-400 hover:bg-red-500/10 hover:text-red-300 transition-all border border-transparent hover:border-red-500/20"
            title="Logout of GoVault"
          >
            <span className="text-sm font-bold">Logout</span>
            <span className="text-lg">󰗼</span> {/* Or use a simple icon like ⎋ */}
          </button>
        </div>
      </div>
    </nav>
  );
}