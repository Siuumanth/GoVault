import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';
import { getChunkSize, calculateChunks, calculateChecksum } from '../lib/file.js';
import {UPLOAD_METHOD} from '../config/options.js';

const binFile = open('../lib/test.wav', 'b');   // normal upload of 1 mb 
const BASE_URL = 'http://localhost:9000/api/upload';
const CHUNK_SIZE = getChunkSize(UPLOAD_METHOD);

const fileSize = binFile.byteLength;

export default function (currentUser) {
    // Safety check for user data
    if (!currentUser || !currentUser.token) return;
    // FORCED DEBUG
   
   // console.log(`DEBUG: VU ${__VU} sees ${actualSize} bytes`);

    if (fileSize === 0) {
        // This stops the test immediately if the memory is empty
        throw new Error("SharedArray is empty inside the VU loop!");
    }

    const token = currentUser.token;
    const headers = {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
    };

    const totalChunks = calculateChunks(fileSize, CHUNK_SIZE);

    // 1. Create Proxy Session
    const sessionRes = http.post(
        `${BASE_URL}/proxy/session`,
        JSON.stringify({ file_name: 'testfile.bin', file_size_bytes: fileSize }),
        { headers }
    );

    check(sessionRes, { 
        'proxy session 200': (r) => r.status === 200 
    });

    if (sessionRes.status !== 200) return;
    const { upload_uuid } = sessionRes.json();

    // 2. Upload Chunks
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

        check(chunkRes, { [`proxy chunk ${i} 200`]: (r) => r.status === 200 });
        
        // Optional: Small sleep to stay under your 100k/min rate limit if VUs are high
        // sleep(0.1); 
    }
}