import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';
import { getChunkSize, calculateChunks, calculateChecksum } from '../lib/file.js';

const BASE_URL = 'http://localhost:9000/api/upload';
const CHUNK_SIZE = getChunkSize('proxy');

// --- SHARED MEMORY: Loads file once for all VUs ---
const binFile = new SharedArray('test file proxy', function () {
    const data = open('../lib/test.wav', 'b');
    
    // VALIDATION: If this fails, k6 will stop before the test starts
    if (!data || data.byteLength === 0) {
        throw new Error("\n\n [!] CRITICAL: test.wav is empty or not found at ../lib/test.wav \n");
    }
    
    console.log(`\n [SUCCESS] Loaded ${data.byteLength} bytes into SharedArray \n`);
    return [data];
})[0];
export default function (currentUser) {
    // Safety check for user data
    if (!currentUser || !currentUser.token) return;

    const token = currentUser.token;
    const headers = {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
    };

    const fileSize = binFile.byteLength;
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