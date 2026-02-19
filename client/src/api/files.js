
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

  // Sharing endpoints
  getShares: (fileId) => request(ENDPOINTS.SHARING.LIST(fileId)),
  
  addShares: (fileId, recipients) => 
    request(ENDPOINTS.SHARING.ADD(fileId), {
      method: 'POST',
      body: JSON.stringify({ recipients }),
    }),

  removeShare: (fileId, userId) => 
    request(ENDPOINTS.SHARING.REMOVE(fileId, userId), {
      method: 'DELETE',
    }),

  // Public access endpoints
  makePublic: (fileId) => 
    request(ENDPOINTS.PUBLIC.CREATE(fileId), {
      method: 'POST',
    }),

  removePublic: (fileId) => 
    request(ENDPOINTS.PUBLIC.DELETE(fileId), {
      method: 'DELETE',
    }),

  // Shortcut endpoints
  addShortcut: (fileId) => 
    request(ENDPOINTS.SHORTCUT.ADD(fileId), {
      method: 'POST',
    }),

  removeShortcut: (fileId) => 
    request(ENDPOINTS.SHORTCUT.DELETE(fileId), {
      method: 'DELETE',
    }),
};