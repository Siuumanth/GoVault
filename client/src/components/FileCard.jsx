export default function FileCard({ file, onDownload, onDelete, onShare, onShortcut }) {
  const formatSize = (bytes) => (bytes / (1024 * 1024)).toFixed(2) + " MB";

  return (
    <div className="bg-[#161b22] border border-[#30363d] p-4 rounded-xl hover:border-blue-500/50 transition-all group relative">
      <div className="flex justify-between items-start mb-3">
        <div className="text-4xl">📄</div>
        <button onClick={() => onShortcut(file.file_id)} className="text-gray-600 hover:text-yellow-500 transition-colors">⭐</button>
      </div>
      
      <h3 className="font-medium truncate text-white mb-1" title={file.name}>{file.name}</h3>
      <p className="text-[10px] text-gray-500 font-mono tracking-tighter uppercase">{file.mime_type || 'Unknown Type'}</p>
      <p className="text-xs text-gray-500 mt-1">{formatSize(file.size_bytes)}</p>
      
      <div className="mt-4 flex gap-1 opacity-0 group-hover:opacity-100 transition-all translate-y-2 group-hover:translate-y-0">
        <button onClick={() => onDownload(file.file_id, file.name)} className="flex-1 bg-blue-600/10 hover:bg-blue-600 text-blue-400 hover:text-white text-[10px] font-bold py-2 rounded-lg transition-colors">
          Download
        </button>
        <button onClick={() => onShare(file)} className="flex-1 bg-emerald-600/10 hover:bg-emerald-600 text-emerald-400 hover:text-white text-[10px] font-bold py-2 rounded-lg transition-colors">
          Share
        </button>
        <button onClick={() => onDelete(file.file_id)} className="px-2 bg-red-900/10 hover:bg-red-600 text-red-400 hover:text-white text-[10px] py-2 rounded-lg transition-colors">
          🗑️
        </button>
      </div>
    </div>
  );
}