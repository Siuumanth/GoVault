import { BASE_URL } from './constants';

export const request = async (endpoint, options = {}) => {
  const token = localStorage.getItem('gv_token');

  // Only skip auth for these two specific GET endpoints:
  // 1. GET /api/files/f/{id} (file details)
  // 2. GET /api/files/f/{id}/download (download)
  const method = (options.method || 'GET').toUpperCase();
  let isPublicFileEndpoint = false;
  
  if (method === 'GET') {
    if (endpoint.startsWith('/api/files/f/')) {
      const afterPrefix = endpoint.slice('/api/files/f/'.length);
      const parts = afterPrefix.split('/');
      // Check if it's exactly /api/files/f/{id} or /api/files/f/{id}/download
      if (parts.length === 1) {
        // /api/files/f/{id} - public
        isPublicFileEndpoint = true;
      } else if (parts.length === 2 && parts[1] === 'download') {
        // /api/files/f/{id}/download - public
        isPublicFileEndpoint = true;
      }
    }
  }
  
  const headers = { ...options.headers };

  if (!(options.body instanceof ArrayBuffer) && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json';
  }

  if (token && !isPublicFileEndpoint) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  try {
    const response = await fetch(`${BASE_URL}${endpoint}`, { ...options, headers });

    if (response.status === 401 && !isPublicFileEndpoint) {
      localStorage.clear();
      alert("Session expired. Please login again.");
      window.location.href = '/'; 
      return null;
    }

    if (response.status === 204) return null;

    if (!response.ok) {
      const errorData = await response.text();
      throw new Error(errorData || `Error: ${response.status} ${response.statusText}`);
    }

    const text = await response.text();
    if (!text || text.trim() === '') return null;
    
    try {
      return JSON.parse(text);
    } catch {
      return text;
    }
  } catch (error) {
    console.error("API Request Failed:", error);
    throw error;
  }
};