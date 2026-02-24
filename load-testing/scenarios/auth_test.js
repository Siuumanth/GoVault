import http from 'k6/http';
import { check, sleep } from 'k6';
import { createUser } from '../lib/auth.js';

const BASE_URL = 'http://localhost:9000';

export function setup() {
  // pre-create N tokens, one per VU
  const tokens = [];
  const vuCount = __ENV.VU_COUNT ? parseInt(__ENV.VU_COUNT) : 50;
  for (let i = 0; i < vuCount; i++) {
    tokens.push(createUser());
  }
  return { tokens };
}

export default function (data) {
  const token = data.tokens[__VU - 1] || data.tokens[0];

  const res = http.post(
    `${BASE_URL}/auth/login`,
    JSON.stringify({ email: `test@govault.com`, password: 'Test@1234' }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(res, {
    'login 200': (r) => r.status === 200,
    'has token': (r) => JSON.parse(r.body).token !== undefined,
  });

  sleep(1);
}