import { useState } from 'react';

export default function FileCard({ file, onDownload, onDelete, onShare, onRename }) {
  const [showMenu, setShowMenu] = useState(false);

  return (
    <div className="bg-[#161b22] border border-[#30363d] p-4 rounded-xl hover:border-blue-500/50 transition-all group relative">
      <div className="flex justify-between">
        <div className="text-3xl">📄</div>
        
        {/* 3-Dot Menu */}
        <div className="relative">
          <button 
            onClick={() => setShowMenu(!showMenu)}
            className="p-1 hover:bg-[#30363d] rounded text-gray-400 hover:text-white transition-colors"
          >
            ⋮
          </button>
          
          {showMenu && (
            <>
              <div className="fixed inset-0 z-10" onClick={() => setShowMenu(false)} />
              <div className="absolute right-0 mt-2 w-48 bg-[#161b22] border border-[#30363d] rounded-lg shadow-2xl z-20 overflow-hidden py-1">
                <button onClick={() => {onRename(file); setShowMenu(false)}} className="w-full text-left px-4 py-2 text-sm hover:bg-blue-600 text-gray-300 hover:text-white">Edit Name</button>
                <button onClick={() => {onShare(file); setShowMenu(false)}} className="w-full text-left px-4 py-2 text-sm hover:bg-blue-600 text-gray-300 hover:text-white">Manage Permissions</button>
                <div className="border-t border-[#30363d] my-1"></div>
                <button onClick={() => {onDelete(file.file_id); setShowMenu(false)}} className="w-full text-left px-4 py-2 text-sm hover:bg-red-600 text-red-400 hover:text-white">Delete File</button>
              </div>
            </>
          )}
        </div>
      </div>

      <div className="mt-3 cursor-pointer" onClick={() => onDownload(file.file_id, file.name)}>
        <h3 className="font-medium truncate text-white" title={file.name}>{file.name}</h3>
        <p className="text-xs text-gray-500 mt-1">{(file.size_bytes / 1024 / 1024).toFixed(2)} MB</p>
      </div>
    </div>
  );
}