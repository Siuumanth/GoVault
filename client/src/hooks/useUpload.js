
import { useState } from 'react';
import { ENDPOINTS } from '../api/endpoints';
import { CHUNK_SIZE, BASE_URL } from '../api/constants';

export const useUpload = () => {
  const [progress, setProgress] = useState(0);
  const [isUploading, setIsUploading] = useState(false);

  const calculateHash = async (buffer) => {
    const hashBuffer = await crypto.subtle.digest('SHA-256', buffer);
    return Array.from(new Uint8Array(hashBuffer))
      .map(b => b.toString(16).padStart(2, '0')).join('');
  };

  const uploadFile = async (file) => {
    if (!file) return;
    setIsUploading(true);
    const storageKey = `gv_up_${file.name}_${file.size}`;
    let session = JSON.parse(localStorage.getItem(storageKey));

    try {
      const token = localStorage.getItem('gv_token');

      if (!session) {
        const res = await fetch(`${BASE_URL}${ENDPOINTS.UPLOAD.SESSION}`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json', 'Authorization': `Bearer ${token}` },
          body: JSON.stringify({ file_name: file.name, file_size_bytes: file.size })
        });
        session = { ...(await res.json()), completed: {} };
        localStorage.setItem(storageKey, JSON.stringify(session));
      }

      for (let i = 0; i < session.total_chunks; i++) {
        if (session.completed[i]) continue;

        const start = i * CHUNK_SIZE;
        const chunk = file.slice(start, start + CHUNK_SIZE);
        const buffer = await chunk.arrayBuffer();
        const checksum = await calculateHash(buffer);

        const chunkRes = await fetch(`${BASE_URL}${ENDPOINTS.UPLOAD.CHUNK(i)}`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`,
            'Upload-UUID': session.upload_uuid,
            'Checksum': checksum,
            'Content-Type': 'application/octet-stream'
          },
          body: buffer
        });

        if (!chunkRes.ok) throw new Error(`Chunk ${i} failed`);

        session.completed[i] = true;
        localStorage.setItem(storageKey, JSON.stringify(session));
        setProgress(Math.round(((i + 1) / session.total_chunks) * 100));
      }

      localStorage.removeItem(storageKey);
      return true;
    } catch (err) {
      alert(err.message);
      return false;
    } finally {
      setIsUploading(false);
      setProgress(0);
    }
  };

  return { uploadFile, progress, isUploading };
};