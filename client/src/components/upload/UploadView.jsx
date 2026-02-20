import { useState } from 'react';

export default function UploadView({ onUpload, progress, isUploading, logs }) {
  const [dragActive, setDragActive] = useState(false);

  const handleDrop = (e) => {
    e.preventDefault();
    setDragActive(false);
    if (e.dataTransfer.files?.[0]) onUpload(e.dataTransfer.files[0]);
  };

  return (
    <div className="max-w-4xl mx-auto space-y-6">
      <div 
        onDragOver={(e) => {e.preventDefault(); setDragActive(true)}}
        onDragLeave={() => setDragActive(false)}
        onDrop={handleDrop}
        className={`border-2 border-dashed rounded-3xl p-12 flex flex-col items-center transition-all ${
          dragActive ? "border-blue-500 bg-blue-500/5" : "border-[#30363d] bg-[#161b22]"
        }`}
      >
        <div className="text-6xl mb-4">☁️</div>
        <p className="text-gray-300">Drag and drop files to upload</p>
        <input type="file" id="f-up" className="hidden" onChange={(e) => onUpload(e.target.files[0])} />
        <label htmlFor="f-up" className="mt-4 bg-blue-600 px-6 py-2 rounded-xl cursor-pointer font-bold">Browse</label>
      </div>

      {(isUploading || logs.length > 0) && (
        <div className="bg-[#161b22] border border-[#30363d] rounded-2xl overflow-hidden shadow-xl">
          <div className="p-4 bg-[#0d1117] border-b border-[#30363d] flex justify-between text-xs font-mono">
            <span className="text-blue-400">UPLOAD LOGS</span>
            <span>{progress}%</span>
          </div>
          <div className="w-full bg-gray-800 h-1"><div className="bg-blue-500 h-full" style={{width: `${progress}%`}} /></div>
          <div className="p-4 h-64 overflow-y-auto font-mono text-[10px] space-y-1 bg-black/20">
            {logs.map((l, i) => <div key={i} className="text-blue-300">[{l.time}] {l.msg}</div>)}
            {!isUploading && progress === 100 && <div className="text-emerald-400 font-bold mt-2">✅ UPLOAD DONE</div>}
          </div>
        </div>
      )}
    </div>
  );
}