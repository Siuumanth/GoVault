# TEst 1: - 200 users with auth, upload and fetch files 
```bash
D:\code\Golang\GoVault\load-testing>k6 run main.js

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/


     execution: local
        script: main.js
        output: -

     scenarios: (100.00%) 1 scenario, 200 max VUs, 6m30s max duration (incl. graceful stop):
              * default: Up to 200 looping VUs for 6m0s over 5 stages (gracefulRampDown: 30s, gracefulStop: 30s)

INFO[0000] [SETUP] Starting... Target: 200 users         source=console
INFO[0023] [SETUP] Done. Created 200 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 43675   113.813394/s
    checks_succeeded...: 100.00% 43675 out of 43675
    checks_failed......: 0.00%   0 out of 43675

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=664.49ms min=806.7µs med=96.93ms max=6.41s p(90)=2.98s p(95)=3.85s
      { expected_response:true }...: avg=664.49ms min=806.7µs med=96.93ms max=6.41s p(90)=2.98s p(95)=3.85s
    http_req_failed................: 0.00% 0 out of 35180
    http_reqs......................: 35180 91.676135/s

    EXECUTION
    iteration_duration.............: avg=4.69s    min=2.04s   med=4.89s   max=8.89s p(90)=6.88s p(95)=7.36s
    iterations.....................: 8695  22.658442/s
    vus............................: 2     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 46 MB 120 kB/s
    data_sent......................: 14 GB 35 MB/s




running (6m23.7s), 000/200 VUs, 8695 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  6m0s
```


---
---


# Test 2: - 200 users with auth, upload and fetch files, 1 mb file , proxy
```bash
D:\code\Golang\GoVault\load-testing>k6 run main.js

         /\      Grafana   /‾‾/
    /\  /  \     |\  __   /  /
   /  \/    \    | |/ /  /   ‾‾\
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/


     execution: local
        script: main.js
        output: -

     scenarios: (100.00%) 1 scenario, 200 max VUs, 6m30s max duration (incl. graceful stop):
              * default: Up to 200 looping VUs for 6m0s over 5 stages (gracefulRampDown: 30s, gracefulStop: 30s)

INFO[0000] [SETUP] Starting... Target: 200 users         source=console
INFO[0022] [SETUP] Done. Created 200 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 62997  164.160969/s
    checks_succeeded...: 99.46% 62661 out of 62997
    checks_failed......: 0.53%  336 out of 62997

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✗ proxy session 200
      ↳  98% — ✓ 12437 / ✗ 153
    ✗ proxy chunk 0 200
      ↳  98% — ✓ 12254 / ✗ 183

    HTTP
    http_req_duration..............: avg=305.63ms min=0s      med=69.25ms max=3.4s  p(90)=977.75ms p(95)=1.33s
      { expected_response:true }...: avg=304.14ms min=763.4µs med=68.04ms max=3.4s  p(90)=976.86ms p(95)=1.33s
    http_req_failed................: 0.66% 336 out of 50607
    http_reqs......................: 50607 131.874441/s

    EXECUTION
    iteration_duration.............: avg=3.23s    min=2.03s   med=3.17s   max=6.45s p(90)=4.5s     p(95)=4.85s
    iterations.....................: 12590 32.807699/s
    vus............................: 4     min=0            max=200
    vus_max........................: 200   min=200          max=200

    NETWORK
    data_received..................: 71 MB 184 kB/s
    data_sent......................: 13 GB 34 MB/s




running (6m23.8s), 000/200 VUs, 12590 complete and 0 interrupted iterati
```


# Test 3: - 200 users with auth, upload and fetch files, 1 mb file , proxy, 1 mb file
```bash


     scenarios: (100.00%) 1 scenario, 200 max VUs, 6m30s max duration (incl. graceful stop):
              * default: Up to 200 looping VUs for 6m0s over 5 stages (gracefulRampDown: 30s, gracefulStop: 30s)

INFO[0000] [SETUP] Starting... Target: 200 users         source=console
INFO[0023] [SETUP] Done. Created 200 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 58870   153.332042/s
    checks_succeeded...: 100.00% 58870 out of 58870
    checks_failed......: 0.00%   0 out of 58870

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=363.02ms min=1.91ms med=79.27ms max=5.25s p(90)=1.34s p(95)=1.99s
      { expected_response:true }...: avg=363.02ms min=1.91ms med=79.27ms max=5.25s p(90)=1.34s p(95)=1.99s
    http_req_failed................: 0.00% 0 out of 47336
    http_reqs......................: 47336 123.290734/s

    EXECUTION
    iteration_duration.............: avg=3.46s    min=2.03s  med=3.39s   max=7.91s p(90)=4.86s p(95)=5.28s
    iterations.....................: 11734 30.562225/s
    vus............................: 2     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 65 MB 170 kB/s
    data_sent......................: 12 GB 32 MB/s




running (6m23.9s), 000/200 VUs, 11734 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  6m0s
```



---


# Test 4: - 300 users with auth, upload and fetch files, 1 mb file , proxy, 1 mb file
```bash
     scenarios: (100.00%) 1 scenario, 300 max VUs, 6m30s max duration (incl. graceful stop):
              * default: Up to 300 looping VUs for 6m0s over 5 stages (gracefulRampDown: 30s, gracefulStop: 30s)

INFO[0000] [SETUP] Starting... Target: 300 users         source=console
INFO[0034] [SETUP] Done. Created 300 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 66285   167.433335/s
    checks_succeeded...: 100.00% 66285 out of 66285
    checks_failed......: 0.00%   0 out of 66285

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=651.78ms min=1.31ms med=131.9ms max=6.95s  p(90)=2.63s p(95)=3.39s
      { expected_response:true }...: avg=651.78ms min=1.31ms med=131.9ms max=6.95s  p(90)=2.63s p(95)=3.39s
    http_req_failed................: 0.00% 0 out of 53388
    http_reqs......................: 53388 134.856014/s

    EXECUTION
    iteration_duration.............: avg=4.64s    min=2.03s  med=4.69s   max=10.32s p(90)=7.05s p(95)=7.7s
    iterations.....................: 13197 33.335109/s
    vus............................: 1     min=0          max=300
    vus_max........................: 300   min=300        max=300

    NETWORK
    data_received..................: 70 MB 178 kB/s
    data_sent......................: 14 GB 35 MB/s




running (6m35.9s), 000/300 VUs, 13197 complete and 0 interrupted iterations
default ✓ [======================================] 000/300 VUs  6m0s
```

