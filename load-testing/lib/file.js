import { sha256 } from 'k6/crypto';
export function getChunkSize(method) {
  // 10MB for proxy, 100MB for direct S3
  return method === 'proxy' ? 10 * 1024 * 1024 : 100 * 1024 * 1024;
}

export function calculateChunks(fileSize, chunkSize) {
  return Math.ceil(fileSize / chunkSize);
}

// Checksum helper remains the same

export function calculateChecksum(buffer) {
  return sha256(buffer, 'hex');
}