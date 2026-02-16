
import { request } from './client';
import { ENDPOINTS } from './endpoints';

export const uploadApi = {
  createSession: (fileName, fileSize) => 
    request(ENDPOINTS.UPLOAD.SESSION, {
      method: 'POST',
      body: JSON.stringify({ 
        file_name: fileName, 
        file_size_bytes: fileSize 
      }),
    }),

  getStatus: (uploadUuid) => 
    request(`${ENDPOINTS.UPLOAD.STATUS}?upload_uuid=${uploadUuid}`),
};