import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';
import { getChunkSize, calculateChunks } from '../lib/file.js';

const BASE_URL = 'http://localhost:9000/api/upload';
const PART_SIZE = getChunkSize('multipart');

// --- SHARED MEMORY: Shared across all VUs to save RAM ---
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
    const totalParts = calculateChunks(fileSize, PART_SIZE);

    // 1. Create Multipart Session
    const sessionRes = http.post(
        `${BASE_URL}/multipart/session`,
        JSON.stringify({
            file_name: 'testfile.bin',
            file_size_bytes: fileSize,
            part_size_bytes: PART_SIZE,
        }),
        { headers }
    );

    check(sessionRes, { 'multipart session 200': (r) => r.status === 200 });
    
    if (sessionRes.status !== 200) return;
    const { upload_uuid, parts } = sessionRes.json();

    // 2. Upload Parts to S3 and Register Etags
    for (let i = 0; i < totalParts; i++) {
        const start = i * PART_SIZE;
        const end = Math.min(start + PART_SIZE, fileSize);
        const partData = binFile.slice(start, end);
        const partNumber = i + 1;
        
        const url = parts[i].url;

        // PUT directly to MinIO
        const s3Res = http.put(url, partData);
        check(s3Res, { [`part ${partNumber} s3 200`]: (r) => r.status === 200 });

        const etag = s3Res.headers['ETag'];

        // Register part in GoVault backend
        const partRes = http.post(
            `${BASE_URL}/multipart/part`,
            JSON.stringify({ 
                upload_uuid, 
                part_number: partNumber, 
                size_bytes: partData.byteLength, 
                etag 
            }),
            { headers }
        );
        check(partRes, { [`part ${partNumber} registered`]: (r) => r.status === 200 });
    }

    // 3. Complete Upload
    const completeRes = http.post(
        `${BASE_URL}/multipart/complete`,
        JSON.stringify({ upload_uuid }),
        { headers }
    );
    check(completeRes, { 'multipart complete 200': (r) => r.status === 200 });
}