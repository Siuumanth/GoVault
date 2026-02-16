
import { request } from './client';
import { ENDPOINTS } from './endpoints';

export const authApi = {
  login: (email, password) => 
    request(ENDPOINTS.AUTH.LOGIN, {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    }),

  signup: (username, email, password) => 
    request(ENDPOINTS.AUTH.SIGNUP, {
      method: 'POST',
      body: JSON.stringify({ username, email, password }),
    }),
};