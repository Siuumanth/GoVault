import http from 'k6/http';
import { check } from 'k6';
import { getChunkSize, calculateChunks, calculateChecksum } from '../lib/file.js';

const BASE_URL = 'http://localhost:9000/api/upload';
const CHUNK_SIZE = getChunkSize('proxy');

// File must be opened here in the Init stage
const binFile = open('../lib/test.wav', 'b');

export default function (currentUser) {
  const token = currentUser.token;
  const headers = {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  };

  const fileSize = binFile.byteLength;
  const totalChunks = calculateChunks(fileSize, CHUNK_SIZE);

  // 1. create session
  const sessionRes = http.post(
    `${BASE_URL}/proxy/session`,
    JSON.stringify({ file_name: 'testfile.bin', file_size_bytes: fileSize }),
    { headers }
  );

  check(sessionRes, { 'session 200': (r) => r.status === 200 });
  const { upload_uuid } = sessionRes.json();

  // 2. upload chunks
  for (let i = 0; i < totalChunks; i++) {
    const start = i * CHUNK_SIZE;
    const end = Math.min(start + CHUNK_SIZE, fileSize);
    const chunk = binFile.slice(start, end);

    const checksum = calculateChecksum(chunk);

    const chunkRes = http.post(
      `${BASE_URL}/proxy/chunk?id=${i}`,
      chunk,
      {
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/octet-stream',
          'Upload-UUID': upload_uuid,
          Checksum: checksum,
        },
      }
    );

    check(chunkRes, { [`chunk ${i} 200`]: (r) => r.status === 200 });
  }
}