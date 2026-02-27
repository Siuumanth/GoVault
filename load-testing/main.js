import { options, UPLOAD_METHOD, SCALE } from './config/options.js';
import { createUser } from './lib/auth.js';
import authTest from './scenarios/auth_test.js';
import filesTest from './scenarios/files_test.js';
import proxyTest from './scenarios/upload_proxy.js';
import s3Test from './scenarios/upload_multipart.js';

export { options };   // exporting options from main.js to the k6 runtime, these are our vu count n stuff 

/**
 setup() runs ONCE before the test starts.
 We use it to populate the database with users so the test doesn't fail.
 */
// set up users
export function setup() {
  const users = [];  // list of our current users 
  
  // Hard-calculate the peak based on your SCALE to be 100% sure
  const peakVus = Math.round(1000 * SCALE); // This matches your stressOptions peak
  
  console.log(`[SETUP] Starting... Target: ${peakVus} users`);

  for (let i = 1; i <= peakVus; i++) {
    const user = createUser(i);   // sign up every user 
    if (user && user.token) {
      users.push(user);
    }
  }

  console.log(`[SETUP] Done. Created ${users.length} users.`);
  return { users: users }; // Ensure the key is exactly 'users'
}
// k6 runs setup first, then sees what is being returned
// passes that data to data in the below default functoin

export default function (data) {   // data = same data returned from setup
  // 1. Check if data exists
  if (!data || !data.users || data.users.length === 0) {
    // This only happens if setup() fails or returns empty
    return; 
  }

  // 2. Index safely. Use modulo (%) so if VU 401 wakes up but we have 400 users, 
  // it just grabs user 1 instead of crashing.
  // In k6, __VU is a global variable that represents the Virtual User (VU) ID

  // the number of vus spwaned depend on what u exported from options 
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