export default function FileCard({ file, onDownload, onDelete, onShare, onShortcut, onRename, onTogglePublic, isPublic, isShortcut }) {
  const formatSize = (bytes) => {
    if (bytes < 1024) return bytes + " B";
    if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(2) + " KB";
    if (bytes < 1024 * 1024 * 1024) return (bytes / (1024 * 1024)).toFixed(2) + " MB";
    return (bytes / (1024 * 1024 * 1024)).toFixed(2) + " GB";
  };

  const formatDate = (dateString) => {
    if (!dateString) return '';
    const date = new Date(dateString);
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(yesterday.getDate() - 1);
    
    if (date.toDateString() === today.toDateString()) return 'Today';
    if (date.toDateString() === yesterday.toDateString()) return 'Yesterday';
    
    const diffDays = Math.floor((today - date) / (1000 * 60 * 60 * 24));
    if (diffDays < 7) return `${diffDays} days ago`;
    
    return date.toLocaleDateString('en-US', { month: 'short', day: 'numeric', year: date.getFullYear() !== today.getFullYear() ? 'numeric' : undefined });
  };

  return (
    <div className="bg-[#161b22] border border-[#30363d] p-4 rounded-xl hover:border-blue-500/50 transition-all group relative">
      {/* Date Header */}
      {file.created_at && (
        <div className="text-[10px] text-gray-500 mb-2 font-mono">
          {formatDate(file.created_at)}
        </div>
      )}
      
      <div className="flex justify-between items-start mb-3">
        <div className="text-4xl">📄</div>
        <div className="flex gap-1">
          <button 
            onClick={() => onShortcut && onShortcut(file.file_id)} 
            className={`transition-colors ${isShortcut ? 'text-yellow-500' : 'text-gray-600 hover:text-yellow-500'}`}
            title={isShortcut ? 'Remove from shortcuts' : 'Add to shortcuts'}
          >
            ⭐
          </button>
          {isPublic && (
            <span className="text-green-500 text-xs" title="Public file">🌐</span>
          )}
        </div>
      </div>
      
      <h3 className="font-medium truncate text-white mb-1" title={file.name}>{file.name}</h3>
      <p className="text-[10px] text-gray-500 font-mono tracking-tighter uppercase">{file.mime_type || 'Unknown Type'}</p>
      <p className="text-xs text-gray-500 mt-1">{formatSize(file.size_bytes)}</p>
      
      <div className="mt-4 flex gap-1 opacity-0 group-hover:opacity-100 transition-all translate-y-2 group-hover:translate-y-0">
        <button onClick={() => onDownload(file.file_id, file.name)} className="flex-1 bg-blue-600/10 hover:bg-blue-600 text-blue-400 hover:text-white text-[10px] font-bold py-2 rounded-lg transition-colors">
          Download
        </button>
        {onRename && (
          <button onClick={() => onRename(file)} className="px-2 bg-gray-600/10 hover:bg-gray-600 text-gray-400 hover:text-white text-[10px] py-2 rounded-lg transition-colors" title="Rename">
            ✏️
          </button>
        )}
        {onShare && (
          <button onClick={() => onShare(file)} className="flex-1 bg-emerald-600/10 hover:bg-emerald-600 text-emerald-400 hover:text-white text-[10px] font-bold py-2 rounded-lg transition-colors">
            Share
          </button>
        )}
        {onTogglePublic && (
          <button 
            onClick={() => onTogglePublic(file)} 
            className={`px-2 text-[10px] py-2 rounded-lg transition-colors ${
              isPublic 
                ? 'bg-green-600/10 hover:bg-green-600 text-green-400 hover:text-white' 
                : 'bg-purple-600/10 hover:bg-purple-600 text-purple-400 hover:text-white'
            }`}
            title={isPublic ? 'Remove public access' : 'Make public'}
          >
            {isPublic ? '🔒' : '🌐'}
          </button>
        )}
        <button onClick={() => onDelete(file.file_id)} className="px-2 bg-red-900/10 hover:bg-red-600 text-red-400 hover:text-white text-[10px] py-2 rounded-lg transition-colors">
          🗑️
        </button>
      </div>
    </div>
  );
}