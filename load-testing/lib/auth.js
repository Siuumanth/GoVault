import http from 'k6/http';
import { check } from 'k6';

const BASE_URL = 'http://localhost:9000';

export function createUser() {
  const id = Math.floor(Math.random() * 1000000);
  const email = `test-${id}@govault.com`;
  const username = `test-${id}`;
  const password = 'Test@1234';

  const signupRes = http.post(
    `${BASE_URL}/auth/signup`,
    JSON.stringify({ username, email, password }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(signupRes, { 'signup 200': (r) => r.status === 200 });

  const loginRes = http.post(
    `${BASE_URL}/auth/login`,
    JSON.stringify({ email, password }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(loginRes, { 'login 200': (r) => r.status === 200 });

  const token = JSON.parse(loginRes.body).token;
  return token;
}