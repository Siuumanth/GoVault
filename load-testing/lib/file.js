import { open } from 'k6/experimental/fs';

let file;

export async function getFile() {
  if (!file) {
    file = await open('./test.wav'); // put your test file here
  }
  return file;
}

export function getChunkSize(method) {
  if (method === 'proxy') return 10 * 1024 * 1024;      // 10MB
  if (method === 'multipart') return 100 * 1024 * 1024; // 100MB
}

export function calculateChunks(fileSize, chunkSize) {
  return Math.ceil(fileSize / chunkSize);
}