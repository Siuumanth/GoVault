
export const BASE_URL = 'http://localhost:9000';
export const CHUNK_SIZE = 5 * 1024 * 1024; // 5MB matching Go shared.ChunkSizeBytes
export const UPLOAD_STATUS = {
  PENDING: 'pending',
  ASSEMBLING: 'assembling',
  UPLOADING: 'uploading',
  COMPLETED: 'completed',
  FAILED: 'failed'
};