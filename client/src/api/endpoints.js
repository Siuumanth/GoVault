export const ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    SIGNUP: '/auth/signup',
  },
  FILES: {
    OWNED: '/api/files/me/owned',
    SHARED: '/api/files/me/shared',
    SHORTCUTS: '/api/files/me/shortcuts',
    DETAILS: (id) => `/api/files/f/${id}`,
    DOWNLOAD: (id) => `/api/files/f/${id}/download`,
  },
  UPLOAD: {
    SESSION: '/api/upload/session',
    CHUNK: (id) => `/api/upload/chunk?id=${id}`,
    STATUS: '/api/upload/status',
  },
  SHARING: {
    LIST: (fileId) => `/api/files/f/${fileId}/shares`,
    ADD: (fileId) => `/api/files/f/${fileId}/shares`,
    REMOVE: (fileId, userId) => `/api/files/f/${fileId}/shares/${userId}`,
  },
  PUBLIC: {
    CREATE: (fileId) => `/api/files/f/${fileId}/public`,
    DELETE: (fileId) => `/api/files/f/${fileId}/public`,
  },
  SHORTCUT: {
    ADD: (fileId) => `/api/files/f/${fileId}/shortcut`,
    DELETE: (fileId) => `/api/files/f/${fileId}/shortcut`,
  },
};