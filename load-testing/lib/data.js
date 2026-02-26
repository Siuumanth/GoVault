// import { SharedArray } from 'k6/data';
// import encoding from 'k6/encoding';

// // ✅ Survives JSON serialization — stored as base64 string
// export const fileData = new SharedArray('test file', function () {
//     const raw = open('../lib/test.wav', 'b');
//     return [encoding.b64encode(raw)]; // store as string, not binary
// })[0];

