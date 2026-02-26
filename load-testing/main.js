import { options, UPLOAD_METHOD, SCALE } from './config/options.js';
import { createUser } from './lib/auth.js';
import authTest from './scenarios/auth_test.js';
import filesTest from './scenarios/files_test.js';
import proxyTest from './scenarios/upload_proxy.js';
import s3Test from './scenarios/upload_multipart.js';

export { options };

/**
 setup() runs ONCE before the test starts.
 We use it to populate the database with users so the test doesn't fail.
 */
export function setup() {
  const users = [];
  
  // Hard-calculate the peak based on your SCALE to be 100% sure
  const peakVus = Math.round(1000 * SCALE); // This matches your stressOptions peak
  
  console.log(`[SETUP] Starting... Target: ${peakVus} users`);

  for (let i = 1; i <= peakVus; i++) {
    const user = createUser(i);
    if (user && user.token) {
      users.push(user);
    }
  }

  console.log(`[SETUP] Done. Created ${users.length} users.`);
  return { users: users }; // Ensure the key is exactly 'users'
}

export default function (data) {
  // 1. Check if data exists
  if (!data || !data.users || data.users.length === 0) {
    // This only happens if setup() fails or returns empty
    return; 
  }

  // 2. Index safely. Use modulo (%) so if VU 401 wakes up but we have 400 users, 
  // it just grabs user 1 instead of crashing.
  const index = (__VU - 1) % data.users.length;
  const currentUser = data.users[index];

  if (!currentUser) return;

  // Run tests
  authTest(currentUser);
  filesTest(currentUser);

  if (UPLOAD_METHOD === 'proxy') {
    proxyTest(currentUser);
  } else {
    s3Test(currentUser);
  }
}