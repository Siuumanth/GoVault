export const SCALE = 0.2; 

export const UPLOAD_METHOD = 'proxy'; // swap to 'multipart' for S3
export const VU_COUNT = Math.round(100 * SCALE); // virtual users count 

export const loadOptions = {
  stages: [
    { duration: '1m', target: Math.round(200 * SCALE) },   // steady ramp
    { duration: '3m', target: Math.round(200 * SCALE) },   // sustained heavy load
    { duration: '1m', target: 0 },                         // ramp down
  ],
  setupTimeout: '8m',
};

export const stressOptions = {
  stages: [
    { duration: '1m', target: Math.round(250 * SCALE) },   // initial push
    { duration: '1m', target: Math.round(500 * SCALE) },   // breaking point search
    { duration: '1m', target: Math.round(750 * SCALE) },   // extreme load
    { duration: '2m', target: Math.round(1000 * SCALE) },  // peak stress
    { duration: '1m', target: 0 },                         // recovery
  ],
  setupTimeout: '8m',
};

export const spikeOptions = {
  stages: [
    { duration: '10s', target: 0 },
    { duration: '20s', target: Math.round(2000 * SCALE) }, // massive sudden spike
    { duration: '1m', target: Math.round(2000 * SCALE) },  // hold the flood
    { duration: '20s', target: 0 },                        // drop
  ],
  setupTimeout: '8m',
};

// change this to loadOptions / stressOptions / spikeOptions
export const options = stressOptions;