import { BASE_URL } from './constants';

export const request = async (endpoint, options = {}) => {
  const token = localStorage.getItem('gv_token');

  // Public Paths Check: Don't add Auth if accessing a specific file resource
  const isPublicFileEndpoint = /^\/api\/files\/f\/[^/]+(\/download)?$/.test(endpoint);
  
  const headers = {
    ...options.headers,
  };

  // Only add JSON content type if not sending raw binary/buffer
  if (!(options.body instanceof ArrayBuffer) && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json';
  }

  if (token && !isPublicFileEndpoint) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  try {
    const response = await fetch(`${BASE_URL}${endpoint}`, { ...options, headers });

    // Handle Token Expiry / Unauthorized
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
    } catch (parseError) {
      return text; 
    }
  } catch (error) {
    console.error("API Request Failed:", error);
    throw error;
  }
};