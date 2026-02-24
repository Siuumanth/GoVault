import { options, UPLOAD_METHOD } from './config/options.js';
import { setup as authSetup, default as authTest } from './scenarios/auth_test.js';
import { setup as filesSetup, default as filesTest } from './scenarios/files_test.js';
import { setup as proxySetup, default as proxyTest } from './scenarios/upload_proxy.js';
import { setup as s3Setup, default as s3Test } from './scenarios/upload_s3.js';

export { options };

export function setup() {
  if (UPLOAD_METHOD === 'proxy') return proxySetup();
  return s3Setup();
}

export default function (data) {
  authTest(data);
  filesTest(data);
  if (UPLOAD_METHOD === 'proxy') {
    proxyTest(data);
  } else {
    s3Test(data);
  }
}