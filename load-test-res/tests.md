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

# recent
# Test 3: - 200 users with auth, upload and fetch files, 1 mb file , proxy, 1 mb file
```bash
  █ TOTAL RESULTS

    checks_total.......: 68265   177.60543/s
    checks_succeeded...: 100.00% 68265 out of 68265
    checks_failed......: 0.00%   0 out of 68265

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=243.98ms min=766.5µs med=59.27ms max=2.33s p(90)=1s    p(95)=1.49s
      { expected_response:true }...: avg=243.98ms min=766.5µs med=59.27ms max=2.33s p(90)=1s    p(95)=1.49s
    http_req_failed................: 0.00% 0 out of 54852
    http_reqs......................: 54852 142.708754/s

    EXECUTION
    iteration_duration.............: avg=2.98s    min=2.03s   med=2.99s   max=4.72s p(90)=3.98s p(95)=4.11s
    iterations.....................: 13613 35.417018/s
    vus............................: 1     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 77 MB 201 kB/s
    data_sent......................: 14 GB 37 MB/s

---
---

# Test 4: - 300 users with auth, upload and fetch files, 1 mb file , proxy, 1 mb file

```bash
     scenarios: (100.00%) 1 scenario, 300 max VUs, 6m30s max duration (incl. graceful stop):
              * default: Up to 300 looping VUs for 6m0s over 5 stages (gracefulRampDown: 30s, gracefulStop: 30s)

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

# Test 5: - 400 users with auth, upload and fetch files, 1 mb file , proxy, 1 mb file
## Starting to see errors 
```bash
  █ TOTAL RESULTS

    checks_total.......: 70775  174.065595/s
    checks_succeeded...: 99.57% 70475 out of 70775
    checks_failed......: 0.42%  300 out of 70775

    ✗ signup status is 200
      ↳  25% — ✓ 100 / ✗ 300
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=940.43ms min=560.4µs med=187.05ms max=8.08s  p(90)=4.12s p(95)=5.5s
      { expected_response:true }...: avg=945.36ms min=560.4µs med=188.25ms max=8.08s  p(90)=4.13s p(95)=5.5s
    http_req_failed................: 0.52% 300 out of 57100
    http_reqs......................: 57100 140.432999/s

    EXECUTION
    iteration_duration.............: avg=5.81s    min=2.03s   med=6.45s    max=11.78s p(90)=9s    p(95)=9.48s
    iterations.....................: 14075 34.616365/s
    vus............................: 4     min=0            max=400
    vus_max........................: 400   min=400          max=400

    NETWORK
    data_received..................: 86 MB 211 kB/s
    data_sent......................: 15 GB 36 MB/s

running (6m46.6s), 000/400 VUs, 14075 complete and 0 interrupted iterations
default ✓ [======================================] 000/400 VUs  6m0s

```


# tEst 6- 500 users 
```bash
  █ TOTAL RESULTS

    checks_total.......: 68265   177.60543/s
    checks_succeeded...: 100.00% 68265 out of 68265
    checks_failed......: 0.00%   0 out of 68265

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=243.98ms min=766.5µs med=59.27ms max=2.33s p(90)=1s    p(95)=1.49s
      { expected_response:true }...: avg=243.98ms min=766.5µs med=59.27ms max=2.33s p(90)=1s    p(95)=1.49s
    http_req_failed................: 0.00% 0 out of 54852
    http_reqs......................: 54852 142.708754/s

    EXECUTION
    iteration_duration.............: avg=2.98s    min=2.03s   med=2.99s   max=4.72s p(90)=3.98s p(95)=4.11s
    iterations.....................: 13613 35.417018/s
    vus............................: 1     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 77 MB 201 kB/s
    data_sent......................: 14 GB 37 MB/s

running (6m24.4s), 000/200 VUs, 13613 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  6m0s
```





---










---

---

---

# S3 Multipart tests
# 1. 400 VU, same configs
```bash

  █ TOTAL RESULTS

    checks_total.......: 95656   234.7078/s
    checks_succeeded...: 100.00% 95656 out of 95656
    checks_failed......: 0.00%   0 out of 95656

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=630.3ms min=718µs med=172.13ms max=6.92s  p(90)=2.24s p(95)=4.15s
      { expected_response:true }...: avg=630.3ms min=718µs med=172.13ms max=6.92s  p(90)=2.24s p(95)=4.15s
    http_req_failed................: 0.00% 0 out of 82448
    http_reqs......................: 82448 202.29979/s

    EXECUTION
    iteration_duration.............: avg=6.02s   min=2.27s med=6.39s    max=10.56s p(90)=8.88s p(95)=9.16s
    iterations.....................: 13608 33.389476/s
    vus............................: 2     min=0          max=400
    vus_max........................: 400   min=400        max=400

    NETWORK
    data_received..................: 89 MB 219 kB/s
    data_sent......................: 14 GB 35 MB/s
```

---



# 2. 600 vu, same 
```bash

  █ TOTAL RESULTS

    checks_total.......: 88415   205.359687/s
    checks_succeeded...: 100.00% 88415 out of 88415
    checks_failed......: 0.00%   0 out of 88415

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=1.26s min=783.9µs med=302.77ms max=14.51s p(90)=4.81s  p(95)=8.65s
      { expected_response:true }...: avg=1.26s min=783.9µs med=302.77ms max=14.51s p(90)=4.81s  p(95)=8.65s
    http_req_failed................: 0.00% 0 out of 76470
    http_reqs......................: 76470 177.615283/s

    EXECUTION
    iteration_duration.............: avg=9.89s min=2.29s   med=9.72s    max=21.58s p(90)=16.28s p(95)=17.52s
    iterations.....................: 12545 29.138011/s
    vus............................: 2     min=0          max=600
    vus_max........................: 600   min=600        max=600

    NETWORK
    data_received..................: 74 MB 173 kB/s
    data_sent......................: 13 GB 31 MB/s

```


---


# 3. 600 vu, same
```bash

  █ TOTAL RESULTS

    checks_total.......: 96290  224.18221/s
    checks_succeeded...: 99.99% 96282 out of 96290
    checks_failed......: 0.00%  8 out of 96290

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 13665 / ✗ 5
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✗ multipart complete 200
      ↳  99% — ✓ 13667 / ✗ 3

    HTTP
    http_req_duration..............: avg=1.13s min=773.1µs med=272.91ms max=13.42s p(90)=3.95s  p(95)=6.02s
      { expected_response:true }...: avg=1.13s min=773.1µs med=272.87ms max=13.42s p(90)=3.95s  p(95)=6.02s
    http_req_failed................: 0.00% 8 out of 83220
    http_reqs......................: 83220 193.752659/s

    EXECUTION
    iteration_duration.............: avg=9.08s min=2.28s   med=8.6s     max=23.76s p(90)=14.99s p(95)=16.77s
    iterations.....................: 13670 31.82647/s
    vus............................: 4     min=0          max=600
    vus_max........................: 600   min=600        max=600

    NETWORK
    data_received..................: 82 MB 191 kB/s
    data_sent......................: 14 GB 34 MB/s

```

---



# 4. 750 vu, same
```bash
   █ TOTAL RESULTS

    checks_total.......: 95740  214.104232/s
    checks_succeeded...: 99.98% 95723 out of 95740
    checks_failed......: 0.01%  17 out of 95740

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 13566 / ✗ 4
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✗ multipart complete 200
      ↳  99% — ✓ 13557 / ✗ 13

    HTTP
    http_req_duration..............: avg=1.52s min=551µs med=383.41ms max=17.19s p(90)=5.52s  p(95)=9.18s
      { expected_response:true }...: avg=1.52s min=551µs med=383.38ms max=17.19s p(90)=5.52s  p(95)=9.18s
    http_req_failed................: 0.02% 17 out of 82920
    http_reqs......................: 82920 185.434749/s

    EXECUTION
    iteration_duration.............: avg=11.5s min=2.29s med=11.71s   max=28.26s p(90)=18.23s p(95)=19.57s
    iterations.....................: 13570 30.346714/s
    vus............................: 12    min=0           max=750
    vus_max........................: 750   min=750         max=750

    NETWORK
    data_received..................: 77 MB 172 kB/s
    data_sent......................: 14 GB 32 MB/s




running (7m27.2s), 000/750 VUs, 13570 complete and 0 interrupted iterations
default ✓ [======================================] 000/750 VUs  6m0s
```

# 5. 200 vu
```bash
INFO[0000] [SETUP] Starting... Target: 200 users         source=console
INFO[0022] [SETUP] Done. Created 200 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 87441   227.969747/s
    checks_succeeded...: 100.00% 87441 out of 87441
    checks_failed......: 0.00%   0 out of 87441

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=176.24ms min=0s    med=58.78ms max=2.82s p(90)=456.01ms p(95)=871.42ms
      { expected_response:true }...: avg=176.24ms min=0s    med=58.78ms max=2.82s p(90)=456.01ms p(95)=871.42ms
    http_req_failed................: 0.00% 0 out of 75178
    http_reqs......................: 75178 195.998555/s

    EXECUTION
    iteration_duration.............: avg=3.26s    min=2.27s med=3.06s   max=6.55s p(90)=4.52s    p(95)=4.98s
    iterations.....................: 12463 32.492617/s
    vus............................: 5     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 89 MB 231 kB/s
    data_sent......................: 13 GB 34 MB/s




running (6m23.6s), 000/200 VUs, 12463 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  6m0s


```


# 6. 1000 vu
```bash


  █ TOTAL RESULTS

    checks_total.......: 89013  187.144066/s
    checks_succeeded...: 99.80% 88839 out of 89013
    checks_failed......: 0.19%  174 out of 89013

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 12564 / ✗ 13
    ✓ multipart session 200
    ✗ part 1 s3 200
      ↳  99% — ✓ 12496 / ✗ 76
    ✓ part 1 registered
    ✗ multipart complete 200
      ↳  99% — ✓ 12476 / ✗ 85

    HTTP
    http_req_duration..............: avg=2.32s  min=0s      med=529.13ms max=1m0s   p(90)=8.25s  p(95)=13.32s
      { expected_response:true }...: avg=2.31s  min=613.4µs med=528.8ms  max=59.39s p(90)=8.24s  p(95)=13.31s
    http_req_failed................: 0.22% 174 out of 77436
    http_reqs......................: 77436 162.804173/s

    EXECUTION
    iteration_duration.............: avg=16.77s min=2.29s   med=14.04s   max=1m42s  p(90)=28.37s p(95)=37.6s
    iterations.....................: 12561 26.408689/s
    vus............................: 10    min=0            max=1000
    vus_max........................: 1000  min=1000         max=1000

    NETWORK
    data_received..................: 67 MB 140 kB/s
    data_sent......................: 13 GB 28 MB/s




running (7m55.6s), 0000/1000 VUs, 12561 complete and 16 interrupted iterations
default ✓ [======================================] 0000/1000 VUs  6m0s

```