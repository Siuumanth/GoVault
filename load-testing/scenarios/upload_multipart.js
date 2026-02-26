import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';
import { getChunkSize, calculateChunks } from '../lib/file.js';
const binFile = open('../lib/test.wav', 'b');   // normal upload of 1 mb
const BASE_URL = 'http://localhost:9000/api/upload';
const PART_SIZE = getChunkSize('multipart');

// const binFile = encoding.b64decode(fileData, 'std', 'b');
const fileSize = binFile.byteLength;

export default function (currentUser) {
    // Safety check for user data
    if (!currentUser || !currentUser.token) return;

    const token = currentUser.token;
    const headers = {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
    };

    const totalParts = calculateChunks(fileSize, PART_SIZE);

    // 1. Create Multipart Session
    const sessionRes = http.post(
        `${BASE_URL}/multipart/session`,
        JSON.stringify({
            file_name: 'testfile.bin',
            file_size_bytes: fileSize,
            part_size_bytes: PART_SIZE,
        }),
        { 
            headers,
            tags: { name: 'multipart_session_init' } 
        }
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
        const s3Res = http.put(url, partData, {
            tags: { name: 's3_part_put' }
        });
        check(s3Res, { [`part ${partNumber} s3 200`]: (r) => r.status === 200 });

       const etag = s3Res.headers['Etag'] || s3Res.headers['etag'] || s3Res.headers['ETag'];

        // Register part in GoVault backend
        const partRes = http.post(
            `${BASE_URL}/multipart/part`,
            JSON.stringify({ 
                upload_uuid, 
                part_number: partNumber, 
                size_bytes: partData.byteLength, 
                etag 
            }),
            { 
                headers,
                tags: { name: 'register_part' }
            }
        );
        check(partRes, { [`part ${partNumber} registered`]: (r) => r.status === 200 });
    }
sleep(0.2)
    // 3. Complete Upload
    const completeRes = http.post(
        `${BASE_URL}/multipart/complete`,
        JSON.stringify({ upload_uuid }),
        { 
            headers,
            tags: { name: 'multipart_complete' }
        }
    );
    check(completeRes, { 'multipart complete 200': (r) => r.status === 200 });
}