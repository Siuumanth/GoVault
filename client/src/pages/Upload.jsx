import { useState, useRef, useEffect } from 'react';
import { useUpload } from '../hooks/useUpload';
import Navbar from '../components/layout/Navbar';
import Sidebar from '../components/layout/Sidebar';

export default function Upload() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [logs, setLogs] = useState([]);
  const { uploadFile, progress, isUploading } = useUpload();
  const logEndRef = useRef(null);

  // Auto-scroll logs to bottom
  useEffect(() => {
    logEndRef.current?.scrollIntoView({ behavior: 'smooth' });
  }, [logs]);

  const addLog = (msg, type = 'info') => {
    const timestamp = new Date().toLocaleTimeString();
    setLogs(prev => [...prev, { timestamp, msg, type }]);
  };

  const handleStartUpload = async () => {
    if (!selectedFile) return;
    addLog(`Starting upload for: ${selectedFile.name}`, 'start');
    
    // Pass a callback to useUpload if you want real-time chunk logs
    // For now, we simulate the log flow based on the process
    const success = await uploadFile(selectedFile);
    
    if (success) addLog("File assembly started on server!", "success");
    else addLog("Upload failed. Check connection.", "error");
  };

  return (
    <div className="flex h-screen bg-[#0d1117] text-gray-200 overflow-hidden">
      <Sidebar activeTab="upload" />
      <div className="flex-1 flex flex-col">
        <Navbar />
        <main className="p-8 overflow-y-auto">
          <div className="max-w-4xl mx-auto space-y-6">
            <h2 className="text-2xl font-bold text-white">Chunked Upload Manager</h2>
            
            {/* File Selection Card */}
            <div className="bg-[#161b22] border border-[#30363d] p-6 rounded-2xl shadow-lg">
              <input 
                type="file" 
                onChange={(e) => setSelectedFile(e.target.files[0])}
                className="block w-full text-sm text-gray-400 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-blue-600/10 file:text-blue-400 hover:file:bg-blue-600/20 cursor-pointer"
              />
              <button 
                onClick={handleStartUpload}
                disabled={!selectedFile || isUploading}
                className="mt-6 w-full bg-blue-600 hover:bg-blue-500 disabled:bg-gray-800 text-white font-bold py-3 rounded-xl transition-all"
              >
                {isUploading ? `Uploading ${progress}%` : 'Execute Multi-Stage Upload'}
              </button>
            </div>

            {/* Log Terminal */}
            <div className="bg-[#010409] border border-[#30363d] rounded-xl overflow-hidden shadow-inner">
              <div className="bg-[#161b22] px-4 py-2 border-b border-[#30363d] flex justify-between">
                <span className="text-xs font-mono text-gray-500 uppercase">Process Logs</span>
                <span className="text-xs text-blue-500 font-mono">Gateway: 9000</span>
              </div>
              <div className="h-64 overflow-y-auto p-4 font-mono text-xs space-y-1">
                {logs.map((log, i) => (
                  <div key={i} className={
                    log.type === 'error' ? 'text-red-400' : 
                    log.type === 'success' ? 'text-emerald-400' : 'text-blue-300'
                  }>
                    <span className="opacity-50">[{log.timestamp}]</span> {log.msg}
                  </div>
                ))}
                <div ref={logEndRef} />
                {logs.length === 0 && <div className="text-gray-600 italic">No activity recorded...</div>}
              </div>
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}