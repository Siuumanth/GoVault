import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:9000';

// Remove the setup() from here!

export default function (currentUser) {
  // Use the token passed from main.js
  const res = http.get(`${BASE_URL}/api/files/me/owned`, {
    headers: { Authorization: `Bearer ${currentUser.token}` },
  });

  check(res, {
    'files 200': (r) => r.status === 200,
  });

  sleep(1);
}