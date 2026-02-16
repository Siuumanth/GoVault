export default function Sidebar({ activeTab, setActiveTab }) {
  const tabs = [
    { id: 'owned', label: 'My Vault', icon: '📂' },
    { id: 'shared', label: 'Shared with Me', icon: '🤝' },
    { id: 'shortcuts', label: 'Shortcuts', icon: '⭐' },
  ];

  return (
    <aside className="w-64 border-r border-[#30363d] bg-[#161b22] flex flex-col">
      <div className="p-6 text-xl font-bold text-white border-b border-[#30363d]">
        GoVault <span className="text-blue-500 text-sm">v4</span>
      </div>
      <nav className="flex-1 p-4 space-y-2">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id)}
            className={`w-full flex items-center gap-3 px-4 py-3 rounded-xl transition-all ${
              activeTab === tab.id 
                ? 'bg-blue-600/10 text-blue-400 border border-blue-600/30' 
                : 'text-gray-400 hover:bg-[#21262d] hover:text-gray-200'
            }`}
          >
            <span>{tab.icon}</span>
            <span className="font-medium">{tab.label}</span>
          </button>
        ))}
      </nav>
    </aside>
  );
}