import { useState } from 'react';

export default function UploadView({ onUpload, progress, isUploading, logs }) {
  const [dragActive, setDragActive] = useState(false);

  const handleDrag = (e) => {
    e.preventDefault();
    e.stopPropagation();
    if (e.type === "dragenter" || e.type === "dragover") setDragActive(true);
    else if (e.type === "dragleave") setDragActive(false);
  };

  const handleDrop = (e) => {
    e.preventDefault();
    e.stopPropagation();
    setDragActive(false);
    if (e.dataTransfer.files && e.dataTransfer.files[0]) {
      onUpload(e.dataTransfer.files[0]);
    }
  };

  return (
    <div className="max-w-4xl mx-auto space-y-6 animate-in fade-in duration-500">
      <h2 className="text-2xl font-bold text-white mb-6">Upload Manager</h2>

      {/* Drag & Drop Zone */}
      <div 
        onDragEnter={handleDrag} onDragLeave={handleDrag} onDragOver={handleDrag} onDrop={handleDrop}
        className={`relative border-2 border-dashed rounded-3xl p-12 transition-all flex flex-col items-center justify-center ${
          dragActive ? "border-blue-500 bg-blue-500/5" : "border-[#30363d] bg-[#161b22]"
        }`}
      >
        <div className="text-6xl mb-4">☁️</div>
        <p className="text-gray-300 font-medium">Drag and drop your file here</p>
        <p className="text-gray-500 text-sm mt-2">or</p>
        <input 
          type="file" id="file-upload" className="hidden" 
          onChange={(e) => e.target.files[0] && onUpload(e.target.files[0])} 
        />
        <label htmlFor="file-upload" className="mt-4 bg-blue-600 hover:bg-blue-500 text-white px-6 py-2 rounded-xl cursor-pointer font-bold transition-all">
          Browse Files
        </label>
      </div>

      {/* Progress & Logs */}
      {(isUploading || logs.length > 0) && (
        <div className="bg-[#161b22] border border-[#30363d] rounded-2xl overflow-hidden shadow-xl">
          <div className="p-4 border-b border-[#30363d] bg-[#0d1117] flex justify-between items-center">
            <span className="text-xs font-mono text-blue-400">STATUS: {isUploading ? 'UPLOADING...' : 'IDLE'}</span>
            <span className="text-xs font-mono text-gray-500">{progress}%</span>
          </div>
          
          {/* Progress Bar */}
          <div className="w-full bg-gray-800 h-1">
            <div className="bg-blue-500 h-full transition-all duration-300" style={{ width: `${progress}%` }} />
          </div>

          {/* Terminal Logs */}
          <div className="p-4 h-64 overflow-y-auto font-mono text-[10px] space-y-1 bg-black/20">
            {logs.map((log, i) => (
              <div key={i} className={log.type === 'error' ? 'text-red-400' : 'text-blue-300'}>
                <span className="opacity-40">[{log.time}]</span> {log.msg}
              </div>
            ))}
            {!isUploading && progress === 100 && (
              <div className="text-emerald-400 font-bold mt-2">✅ UPLOAD DONE</div>
            )}
          </div>
        </div>
      )}
    </div>
  );
}