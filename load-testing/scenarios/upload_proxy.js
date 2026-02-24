import http from 'k6/http';
import { check } from 'k6';
import { createUser } from '../lib/auth.js';
import { getChunkSize, calculateChunks } from '../lib/file.js';
  import { sha256 } from 'k6/crypto';

const BASE_URL = 'http://localhost:9000/api/upload';
const CHUNK_SIZE = getChunkSize('proxy');

export function setup() {
  const tokens = [];
  const vuCount = __ENV.VU_COUNT ? parseInt(__ENV.VU_COUNT) : 50;
  for (let i = 0; i < vuCount; i++) {
    tokens.push(createUser());
  }
  return { tokens };
}

export default function (data) {
  const token = data.tokens[__VU - 1] || data.tokens[0];
  const headers = {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  };

  // read test file
  const file = open('./assets/testfile.bin', 'b');
  const fileSize = file.byteLength;
  const totalChunks = calculateChunks(fileSize, CHUNK_SIZE);

  // 1. create session
  const sessionRes = http.post(
    `${BASE_URL}/proxy/session`,
    JSON.stringify({ file_name: 'testfile.bin', file_size_bytes: fileSize }),
    { headers }
  );

  check(sessionRes, { 'session 200': (r) => r.status === 200 });
  const { upload_uuid } = JSON.parse(sessionRes.body);

  // 2. upload chunks
  for (let i = 0; i < totalChunks; i++) {
    const start = i * CHUNK_SIZE;
    const end = Math.min(start + CHUNK_SIZE, fileSize);
    const chunk = file.slice(start, end);

    // compute checksum
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

function calculateChecksum(buffer) {
  // k6 doesn't have crypto.subtle, use k6/crypto
  return sha256(buffer, 'hex');
}