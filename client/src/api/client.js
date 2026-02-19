import { BASE_URL } from './constants';

export const request = async (endpoint, options = {}) => {
  const token = localStorage.getItem('gv_token');
  
  const headers = {
    ...options.headers,
  };

  // Only add Content-Type: json if we aren't sending a raw binary chunk
  if (!(options.body instanceof ArrayBuffer) && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json';
  }

  if (token) {
    headers['Authorization'] = `Bearer ${token}`;
  }

  const response = await fetch(`${BASE_URL}${endpoint}`, {
    ...options,
    headers,
  });

  if (response.status === 204) return null;
  if (!response.ok) {
    const error = await response.text();
    throw new Error(error || response.statusText);
  }

  const text = await response.text();
  if (!text || text.trim() === '') return null;
  try {
    return JSON.parse(text);
  } catch {
    return null;
  }
};