import { useState } from 'react';
import { BASE_URL, CHUNK_SIZE } from '../api/constants';
import { ENDPOINTS } from '../api/endpoints';

export const useUpload = () => {
  const [progress, setProgress] = useState(0);
  const [isUploading, setIsUploading] = useState(false);
  const [logs, setLogs] = useState([]);

  const addLog = (msg, type = 'info') => {
    setLogs(prev => [...prev, { time: new Date().toLocaleTimeString(), msg, type }]);
  };

  const calculateHash = async (buffer) => {
    const hashBuffer = await crypto.subtle.digest('SHA-256', buffer);
    return Array.from(new Uint8Array(hashBuffer)).map(b => b.toString(16).padStart(2, '0')).join('');
  };

  const uploadFile = async (file) => {
    setIsUploading(true);
    setLogs([]); // Reset logs for new upload
    setProgress(0);

    try {
      const token = localStorage.getItem('gv_token');
      addLog(`Initializing session for ${file.name}...`);

      const res = await fetch(`${BASE_URL}${ENDPOINTS.UPLOAD.SESSION}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
        body: JSON.stringify({ file_name: file.name, file_size_bytes: file.size })
      });
      
      const sessionData = await res.json();
      addLog(`Session created: ${sessionData.upload_uuid}`, 'success');

      for (let i = 0; i < sessionData.total_chunks; i++) {
        const start = i * CHUNK_SIZE;
        const chunk = file.slice(start, start + CHUNK_SIZE);
        const buffer = await chunk.arrayBuffer();
        const checksum = await calculateHash(buffer);

        addLog(`Uploading Chunk ${i} (Hash: ${checksum.substring(0,8)}...)`);

        const chunkRes = await fetch(`${BASE_URL}${ENDPOINTS.UPLOAD.CHUNK(i)}`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Upload-UUID': sessionData.upload_uuid,
            'Checksum': checksum,
            'Content-Type': 'application/octet-stream'
          },
          body: buffer
        });

        if (!chunkRes.ok) throw new Error(`Chunk ${i} failed`);
        setProgress(Math.round(((i + 1) / sessionData.total_chunks) * 100));
      }

      addLog("All chunks uploaded. Server is assembling...", 'success');
      alert('Upload successful. It might take some time before download is possible.');
      return true;
    } catch (err) {
      addLog(`Error: ${err.message}`, 'error');
      return false;
    } finally {
      setIsUploading(false);
    }
  };

  return { uploadFile, progress, isUploading, logs };
};