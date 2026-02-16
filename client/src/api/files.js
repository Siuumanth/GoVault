
import { request } from './client';
import { ENDPOINTS } from './endpoints';

export const filesApi = {
  getOwned: () => request(ENDPOINTS.FILES.OWNED),
  
  getShared: () => request(ENDPOINTS.FILES.SHARED),
  
  getShortcuts: () => request(ENDPOINTS.FILES.SHORTCUTS),

  getDownloadUrl: (fileId) => request(ENDPOINTS.FILES.DOWNLOAD(fileId)),

  rename: (fileId, newName) => 
    request(ENDPOINTS.FILES.DETAILS(fileId), {
      method: 'PATCH',
      body: JSON.stringify({ name: newName }),
    }),

  delete: (fileId) => 
    request(ENDPOINTS.FILES.DETAILS(fileId), {
      method: 'DELETE',
    }),
};