import http from 'k6/http';
import { check, sleep } from 'k6';

const BASE_URL = 'http://localhost:9000';
export default function (currentUser) {
      const res = http.post(
        `${BASE_URL}/auth/login`,
        JSON.stringify({ 
          email: `test-${currentUser.id}@govault.com`, 
          password: 'Test@1234' 
        }),
        { headers: { 'Content-Type': 'application/json' } }
      );

      // ONLY try to parse if status is 200
      let token;
      if (res.status === 200) {
        token = res.json().token;
      }

      check(res, {
        'login 200': (r) => r.status === 200,
        'has token': () => token !== undefined,
      });

  // check(currentUser, {
  //       'has valid session token': (u) => u.token !== undefined && u.token.length > 0,
  //   });

  sleep(1);
}