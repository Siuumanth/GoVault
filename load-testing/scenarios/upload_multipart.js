import http from 'k6/http';
import { check } from 'k6';
import { getChunkSize, calculateChunks } from '../lib/file.js';

const BASE_URL = 'http://localhost:9000/api/upload';
const PART_SIZE = getChunkSize('multipart');

// Open file at the top (Init stage)
const binFile = open('../lib/test.wav', 'b');

export default function (currentUser) {
  const token = currentUser.token;
  const headers = {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  };

  const fileSize = binFile.byteLength;
  const totalParts = calculateChunks(fileSize, PART_SIZE);

  // 1. create multipart session
  const sessionRes = http.post(
    `${BASE_URL}/multipart/session`,
    JSON.stringify({
      file_name: 'testfile.bin',
      file_size_bytes: fileSize,
      part_size_bytes: PART_SIZE,
    }),
    { headers }
  );

  check(sessionRes, { 'session 200': (r) => r.status === 200 });
  const session = sessionRes.json();
  const { upload_uuid, parts } = session;

  // 2. upload parts directly to presigned URLs + register etag
  for (let i = 0; i < totalParts; i++) {
    const start = i * PART_SIZE;
    const end = Math.min(start + PART_SIZE, fileSize);
    const part = binFile.slice(start, end);
    const partNumber = i + 1;
    
    // Ensure the part URL exists in the response
    const url = parts[i].url;

    // PUT directly to MinIO/S3
    const s3Res = http.put(url, part);
    check(s3Res, { [`part ${partNumber} s3 200`]: (r) => r.status === 200 });

    const etag = s3Res.headers['ETag'];

    // register part in backend
    const partRes = http.post(
      `${BASE_URL}/multipart/part`,
      JSON.stringify({ 
        upload_uuid, 
        part_number: partNumber, 
        size_bytes: part.byteLength, 
        etag 
      }),
      { headers }
    );
    check(partRes, { [`part ${partNumber} registered`]: (r) => r.status === 200 });
  }

  // 3. complete
  const completeRes = http.post(
    `${BASE_URL}/multipart/complete`,
    JSON.stringify({ upload_uuid }),
    { headers }
  );
  check(completeRes, { 'complete 200': (r) => r.status === 200 });
}