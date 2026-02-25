import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:9000';

export default function (currentUser) {
  // Now we use the ID that was actually created for this specific VU
  const res = http.post(
    `${BASE_URL}/auth/login`,
    JSON.stringify({ 
      email: `test-${currentUser.id}@govault.com`, 
      password: 'Test@1234' 
    }),
    { headers: { 'Content-Type': 'application/json' } }
  );

  check(res, {
    'login 200': (r) => r.status === 200,
    'has token': (r) => r.json().token !== undefined,
  });

  sleep(1);
}