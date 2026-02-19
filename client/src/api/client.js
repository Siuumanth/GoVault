import { BASE_URL } from './constants';

export const request = async (endpoint, options = {}) => {
  const token = localStorage.getItem('gv_token');
  
  // Public Paths Check: Don't add Auth if accessing a specific file resource
  // Regex matches: /api/files/f/[uuid] or /api/files/f/[uuid]/download
  const isPublicFileEndpoint = /^\/api\/files\/f\/[^/]+(\/download)?$/.test(endpoint);
  
  const headers = { ...options.headers };

  if (!(options.body instanceof ArrayBuffer) && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json';
  }

  // Only add token if it's NOT a public file endpoint and token exists
  if (token && !isPublicFileEndpoint) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  try {
    const response = await fetch(`${BASE_URL}${endpoint}`, { ...options, headers });

    // Handle Token Expiry / Unauthorized
    if (response.status === 401 && !isPublicFileEndpoint) {
      localStorage.clear();
      alert("Session expired. Please login again.");
      window.location.href = '/'; // Force redirect to login
      return;
    }

    if (response.status === 204) return null;
    if (!response.ok) {
      const error = await response.text();
      throw new Error(error || response.statusText);
    }

    return response.json();
  } catch (err) {
    console.error("Fetch Error:", err);
    throw err;
  }
};