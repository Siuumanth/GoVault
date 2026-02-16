export const ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    SIGNUP: '/auth/signup',
  },
  FILES: {
    OWNED: '/api/files/owned',
    SHARED: '/api/files/shared',
    SHORTCUTS: '/api/files/shortcuts',
    DETAILS: (id) => `/api/files/${id}`,
    DOWNLOAD: (id) => `/api/files/${id}/download`,
  },
  UPLOAD: {
    SESSION: '/api/upload/session',
    CHUNK: (id) => `/api/upload/chunk?id=${id}`,
    STATUS: '/api/upload/status',
  }
};