export const ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    SIGNUP: '/auth/signup',
  },
  FILES: {
    // These match r.Route("/", ...) with r.Get("/me/...")
    // listing files in each tab
    OWNED: '/api/files/me/owned',
    SHARED: '/api/files/me/shared',
    SHORTCUTS: '/api/files/me/shortcuts',
    
    // These match r.Route("/f/{fileID}", ...)
    DETAILS: (id) => `/api/files/f/${id}`,
    DOWNLOAD: (id) => `/api/files/f/${id}/download`,
    COPY: (id) => `/api/files/f/${id}/copy`,
    
    // Public/Private Toggles (owner only)
    PUBLIC: (id) => `/api/files/f/${id}/public`,

    // Shares - matches r.Route("/shares", ...) inside /f/{fileID}
    SHARES: (id) => `/api/files/f/${id}/shares`,
    MANAGE_USER: (fileId, userId) => `/api/files/f/${fileId}/shares/${userId}`,

    // Shortcuts - matches /f/{fileID}/shortcut
    SHORTCUT: (id) => `/api/files/f/${id}/shortcut`,
  },
  UPLOAD: {
    SESSION: '/api/upload/session',
    CHUNK: (id) => `/api/upload/chunk?id=${id}`,
    STATUS: '/api/upload/status',
  },
  SHARING: {
    LIST: (fileId) => `/api/files/f/${fileId}/shares`,  // get all shares for a file
    ADD: (fileId) => `/api/files/f/${fileId}/shares`,
    // Handles PATCH (Update) and DELETE (Remove)
    UPDATE: (fileId, userId) => `/api/files/f/${fileId}/shares/${userId}`,
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