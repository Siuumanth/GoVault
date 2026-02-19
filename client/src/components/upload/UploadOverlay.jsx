import { useState, useEffect, useRef } from 'react';

export default function UploadOverlay({ file, progress, logs, isUploading, onClose }) {
  const logEndRef = useRef(null);

  useEffect(() => {
    logEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [logs]);

  if (!file && !isUploading) return null;

  return (
    <div className="fixed inset-0 bg-black/80 backdrop-blur-md z-[100] flex items-center justify-center p-6">
      <div className="bg-[#161b22] border border-[#30363d] w-full max-w-2xl rounded-2xl shadow-2xl flex flex-col max-h-[80vh]">
        <div className="p-4 border-b border-[#30363d] flex justify-between items-center">
          <h3 className="text-white font-bold">Live Upload: {file?.name}</h3>
          {!isUploading && <button onClick={onClose} className="text-gray-500 hover:text-white text-xl">✕</button>}
        </div>

        <div className="p-6 overflow-y-auto flex-1 font-mono text-xs space-y-2 bg-[#0d1117]">
          {logs.map((log, i) => (
            <div key={i} className={log.type === 'error' ? 'text-red-400' : 'text-blue-400'}>
              <span className="opacity-40">[{log.time}]</span> {log.msg}
            </div>
          ))}
          <div ref={logEndRef} />
        </div>

        <div className="p-4 border-t border-[#30363d] bg-[#161b22]">
          <div className="flex justify-between text-xs text-gray-400 mb-2">
            <span>Progress</span>
            <span>{progress}%</span>
          </div>
          <div className="w-full bg-gray-800 h-2 rounded-full overflow-hidden">
            <div className="bg-blue-500 h-full transition-all" style={{ width: `${progress}%` }} />
          </div>
        </div>
      </div>
    </div>
  );
}