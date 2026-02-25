import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:9000';

/**
  @param {number} index - Optional index to create predictable credentials
 */
export function createUser(index = null) {
  // Use the index if provided, otherwise a random number
  const id = index !== null ? index : Math.floor(Math.random() * 1000000);
  
  const payload = JSON.stringify({
    username: `user_${id}`,
    email: `test-${id}@govault.com`,
    password: 'Test@1234',
  });

  const params = {
    headers: { 'Content-Type': 'application/json' },
  };

  const signupRes = http.post(`${BASE_URL}/auth/signup`, payload, params);

  check(signupRes, {
    'signup status is 200': (r) => r.status === 200,
  });
  
  sleep(0.1); // Short pause to avoid slamming the DB during setup

  const loginRes = http.post(`${BASE_URL}/auth/login`, payload, params);

  // Return an object with both the token and the unique ID used
  return {
    token: loginRes.json().token,
    id: id
  };
}