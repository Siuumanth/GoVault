import { options, UPLOAD_METHOD } from './config/options.js';
import { createUser } from './lib/auth.js';
import authTest from './scenarios/auth_test.js';
import filesTest from './scenarios/files_test.js';
import proxyTest from './scenarios/upload_proxy.js';
import s3Test from './scenarios/upload_multipart.js';

export { options };

/**
 * setup() runs ONCE before the test starts.
 * We use it to populate the database with users so the test doesn't fail.
 */
export function setup() {
  const users = [];
  
  // Get the highest 'target' from your stages (e.g., 400)
  const maxVus = Math.max(...options.stages.map(s => s.target));
  
  console.log(`[SETUP] Creating ${maxVus} test users in GoVault...`);

  for (let i = 1; i <= maxVus; i++) {
    // We pass 'i' so credentials are predictable (test-1, test-2, etc.)
    users.push(createUser(i)); 
  }

  // This 'users' array is passed to the default function as 'data'
  return { users };
}

/**
 * default() is the code each Virtual User (VU) actually runs.
 */
export default function (data) {
  // Grab the specific user assigned to this VU based on its ID (__VU)
  const currentUser = data.users[__VU - 1];

  // Run the sub-tests
  authTest(currentUser);
  filesTest(currentUser);

  // Switch upload strategy based on your config/options.js
  if (UPLOAD_METHOD === 'proxy') {
    proxyTest(currentUser);
  } else {
    s3Test(currentUser);
  }
}