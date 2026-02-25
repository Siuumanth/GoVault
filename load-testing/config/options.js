export const UPLOAD_METHOD = 'proxy'; // swap to 'multipart' for S3
export const VU_COUNT = 100; // virtual users count 

export const loadOptions = {
  stages: [
    { duration: '30s', target: 50 },   // ramp up
    { duration: '1m', target: 50 },    // hold
    { duration: '30s', target: 0 },    // ramp down
  ],
};

export const stressOptions = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '30s', target: 200 },
    { duration: '30s', target: 300 },
    { duration: '30s', target: 400 },
    { duration: '30s', target: 0 },
  ],
};

export const spikeOptions = {
  stages: [
    { duration: '10s', target: 0 },
    { duration: '10s', target: 500 },  // sudden spike
    { duration: '30s', target: 500 },  // hold
    { duration: '10s', target: 0 },    // drop
  ],
};

// change this to loadOptions / stressOptions / spikeOptions
export const options = stressOptions;