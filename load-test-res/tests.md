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
```

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

    checks_total.......: 74690   178.28226/s
    checks_succeeded...: 100.00% 74690 out of 74690
    checks_failed......: 0.00%   0 out of 74690

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=1.21s min=921.4µs med=208.44ms max=9.71s p(90)=5.77s  p(95)=6.8s
      { expected_response:true }...: avg=1.21s min=921.4µs med=208.44ms max=9.71s p(90)=5.77s  p(95)=6.8s
    http_req_failed................: 0.00% 0 out of 60352
    http_reqs......................: 60352 144.057985/s

    EXECUTION
    iteration_duration.............: avg=6.92s min=2.03s   med=7.37s    max=13.1s p(90)=11.05s p(95)=11.61s
    iterations.....................: 14838 35.417756/s
    vus............................: 1     min=0          max=500
    vus_max........................: 500   min=500        max=500

    NETWORK
    data_received..................: 73 MB 174 kB/s
    data_sent......................: 16 GB 37 MB/s

```


---


---

# test 7 - 600 vu

```bash
 
  █ TOTAL RESULTS

    checks_total.......: 73832  164.981371/s
    checks_succeeded...: 84.56% 62438 out of 73832
    checks_failed......: 15.43% 11394 out of 73832

    ✓ signup status is 200
    ✗ login 200
      ↳  82% — ✓ 12533 / ✗ 2684
    ✗ has token
      ↳  82% — ✓ 12533 / ✗ 2684
    ✗ files 200
      ↳  82% — ✓ 12616 / ✗ 2586
    ✗ proxy session 200
      ↳  80% — ✓ 12268 / ✗ 2918
    ✗ proxy chunk 0 200
      ↳  95% — ✓ 11738 / ✗ 522

    HTTP
    http_req_duration..............: avg=2.09s  min=0s       med=317.4ms  max=1m0s   p(90)=8.44s  p(95)=10.02s
      { expected_response:true }...: avg=2.15s  min=631.19µs med=373.22ms max=59.76s p(90)=8.76s  p(95)=10.07s
    http_req_failed................: 14.67% 8710 out of 59365
    http_reqs......................: 59365  132.654121/s

    EXECUTION
    iteration_duration.............: avg=10.17s min=2s       med=10.04s   max=1m19s  p(90)=15.48s p(95)=16.84s
    iterations.....................: 15178  33.916015/s
    vus............................: 1      min=0             max=750
    vus_max........................: 750    min=750           max=750

    NETWORK
    data_received..................: 53 MB  119 kB/s
    data_sent......................: 13 GB  29 MB/s




running (7m27.5s), 000/750 VUs, 15178 complete and 42 interrupted iterations
default ✓ [======================================] 000/750 VUs  6m0s


```
# test 8  - 750 VU

```bash
  █ TOTAL RESULTS

    checks_total.......: 75985  170.352861/s
    checks_succeeded...: 99.97% 75966 out of 75985
    checks_failed......: 0.02%  19 out of 75985

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 15028 / ✗ 19
    ✓ proxy session 200
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=2.04s  min=751.7µs med=357.28ms max=22.53s p(90)=8.07s  p(95)=11.18s
      { expected_response:true }...: avg=2.04s  min=751.7µs med=357.14ms max=22.53s p(90)=8.07s  p(95)=11.19s
    http_req_failed................: 0.03% 19 out of 61688
    http_reqs......................: 61688 138.300024/s

    EXECUTION
    iteration_duration.............: avg=10.37s min=2.03s   med=10.11s   max=40.63s p(90)=17.44s p(95)=19.07s
    iterations.....................: 15047 33.734283/s
    vus............................: 10    min=0           max=750
    vus_max........................: 750   min=750         max=750

    NETWORK
    data_received..................: 66 MB 148 kB/s
    data_sent......................: 16 GB 35 MB/s


```

# test 9  - 850 VU ()zap
```bash

  █ TOTAL RESULTS

    checks_total.......: 92092  195.747004/s
    checks_succeeded...: 72.33% 66617 out of 92092
    checks_failed......: 27.66% 25475 out of 92092

    ✓ signup status is 200
    ✗ login 200
      ↳  68% — ✓ 13314 / ✗ 6241
    ✗ has token
      ↳  68% — ✓ 13314 / ✗ 6241
    ✗ files 200
      ↳  68% — ✓ 13355 / ✗ 6176
    ✗ proxy session 200
      ↳  67% — ✓ 13110 / ✗ 6381
    ✗ proxy chunk 0 200
      ↳  96% — ✓ 12674 / ✗ 436

    HTTP
    http_req_duration..............: avg=1.84s min=0s med=274.21ms max=1m0s   p(90)=7.93s  p(95)=9.62s
      { expected_response:true }...: avg=2.1s  min=0s med=376.45ms max=59.93s p(90)=8.69s  p(95)=9.84s
    http_req_failed................: 26.20% 19234 out of 73387
    http_reqs......................: 73387  155.988418/s

    EXECUTION
    iteration_duration.............: avg=8.94s min=2s med=8.04s    max=2m2s   p(90)=15.13s p(95)=17.19s
    iterations.....................: 19491  41.429276/s
    vus............................: 1      min=0              max=850
    vus_max........................: 850    min=713            max=850

    NETWORK
    data_received..................: 58 MB  123 kB/s
    data_sent......................: 14 GB  29 MB/s




running (7m50.5s), 000/850 VUs, 19491 complete and 65 interrupted iterations
default ✓ [======================================] 000/850 VUs  6m0s
```

---

# test 10  - 925 VU
```bash
  █ TOTAL RESULTS

    checks_total.......: 86136  184.756761/s
    checks_succeeded...: 81.98% 70618 out of 86136
    checks_failed......: 18.01% 15518 out of 86136

    ✓ signup status is 200
    ✗ login 200
      ↳  78% — ✓ 13939 / ✗ 3893
    ✗ has token
      ↳  78% — ✓ 13939 / ✗ 3893
    ✗ files 200
      ↳  78% — ✓ 13937 / ✗ 3878
    ✗ proxy session 200
      ↳  78% — ✓ 13939 / ✗ 3854
    ✓ proxy chunk 0 200

    HTTP
    http_req_duration..............: avg=2.22s  min=0s      med=348.19ms max=1m0s   p(90)=9.41s  p(95)=11.52s
      { expected_response:true }...: avg=2.4s   min=752.2µs med=433.28ms max=59.93s p(90)=10.01s p(95)=12s
    http_req_failed................: 16.79% 11625 out of 69229
    http_reqs......................: 69229  148.49222/s

    EXECUTION
    iteration_duration.............: avg=10.65s min=2s      med=10.48s   max=2m2s   p(90)=17.91s p(95)=19.85s
    iterations.....................: 17793  38.164961/s
    vus............................: 5      min=0              max=924
    vus_max........................: 925    min=925            max=925

    NETWORK
    data_received..................: 59 MB  127 kB/s
    data_sent......................: 15 GB  31 MB/s

```

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

    checks_total.......: 100854  234.752307/s
    checks_succeeded...: 100.00% 100854 out of 100854
    checks_failed......: 0.00%   0 out of 100854

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=1.05s min=0s    med=232.61ms max=12.66s p(90)=3.97s  p(95)=8.05s
      { expected_response:true }...: avg=1.05s min=0s    med=232.61ms max=12.66s p(90)=3.97s  p(95)=8.05s
    http_req_failed................: 0.00% 0 out of 87132
    http_reqs......................: 87132 202.812363/s

    EXECUTION
    iteration_duration.............: avg=8.63s min=2.27s med=8.11s    max=17.66s p(90)=14.24s p(95)=15.18s
    iterations.....................: 14322 33.336531/s
    vus............................: 1     min=0          max=600
    vus_max........................: 600   min=600        max=600

    NETWORK
    data_received..................: 88 MB 204 kB/s
    data_sent......................: 15 GB 35 MB/s



```

---



# 4. 750 vu, same
```bash
   
  █ TOTAL RESULTS

    checks_total.......: 104262  233.95752/s
    checks_succeeded...: 100.00% 104262 out of 104262
    checks_failed......: 0.00%   0 out of 104262

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=1.37s  min=739.8µs med=289.22ms max=23.97s p(90)=5.22s  p(95)=9.02s
      { expected_response:true }...: avg=1.37s  min=739.8µs med=289.22ms max=23.97s p(90)=5.22s  p(95)=9.02s
    http_req_failed................: 0.00% 0 out of 90220
    http_reqs......................: 90220 202.448135/s

    EXECUTION
    iteration_duration.............: avg=10.56s min=2.27s   med=10.3s    max=41.11s p(90)=18.63s p(95)=20.36s
    iterations.....................: 14780 33.165412/s
    vus............................: 3     min=0          max=750
    vus_max........................: 750   min=750        max=750

    NETWORK
    data_received..................: 87 MB 195 kB/s
    data_sent......................: 16 GB 35 MB/s




running (7m25.6s), 000/750 VUs, 14780 complete and 12 interrupted iterations
default ✓ [======================================] 000/750 VUs  6m0s
```

# 5. 200 vu
```bash
  █ TOTAL RESULTS

    checks_total.......: 92852   241.915286/s
    checks_succeeded...: 100.00% 92852 out of 92852
    checks_failed......: 0.00%   0 out of 92852

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ multipart session 200
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=144.41ms min=567.6µs med=51.03ms max=2.76s p(90)=311.58ms p(95)=776.31ms
      { expected_response:true }...: avg=144.41ms min=567.6µs med=51.03ms max=2.76s p(90)=311.58ms p(95)=776.31ms
    http_req_failed................: 0.00% 0 out of 79816
    http_reqs......................: 79816 207.951476/s

    EXECUTION
    iteration_duration.............: avg=3.07s    min=2.27s   med=2.91s   max=6.41s p(90)=4.06s    p(95)=4.45s
    iterations.....................: 13236 34.484887/s
    vus............................: 3     min=0          max=200
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 95 MB 247 kB/s
    data_sent......................: 14 GB 36 MB/s




running (6m23.8s), 000/200 VUs, 13236 complete and 0 interrupted iterations
default ✓ [======================================] 000/200 VUs  6m0s


```


# 1000 vu test 2, many fails

```bash
 █  █ TOTAL RESULTS

    checks_total.......: 135488 285.354581/s
    checks_succeeded...: 67.36% 91274 out of 135488
    checks_failed......: 32.63% 44214 out of 135488

    ✓ signup status is 200
    ✗ login 200
      ↳  53% — ✓ 12921 / ✗ 11020
    ✗ has token
      ↳  53% — ✓ 12921 / ✗ 11020
    ✗ files 200
      ↳  53% — ✓ 12741 / ✗ 11184
    ✗ multipart session 200
      ↳  54% — ✓ 12923 / ✗ 10990
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    HTTP
    http_req_duration..............: avg=1.37s min=0s      med=225.58ms max=1m0s  p(90)=2.62s  p(95)=10.88s
      { expected_response:true }...: avg=1.74s min=574.2µs med=369.07ms max=59.5s p(90)=6.95s  p(95)=11.8s
    http_req_failed................: 29.49% 33194 out of 112547
    http_reqs......................: 112547 237.037982/s

    EXECUTION
    iteration_duration.............: avg=8.61s min=2s      med=3.2s     max=2m2s  p(90)=18.48s p(95)=20.19s
    iterations.....................: 23912  50.361646/s
    vus............................: 2      min=0               max=1000
    vus_max........................: 1000   min=1000            max=1000

    NETWORK
    data_received..................: 73 MB  154 kB/s
    data_sent......................: 14 GB  29 MB/s




```

# 7. 850 VU
4 users failed to upload
```bash
 
  █ TOTAL RESULTS

    checks_total.......: 89099  194.703672/s
    checks_succeeded...: 99.36% 88529 out of 89099
    checks_failed......: 0.63%  570 out of 89099

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 12561 / ✗ 59
    ✗ multipart session 200
      ↳  99% — ✓ 12591 / ✗ 29
    ✓ part 1 s3 200
    ✗ part 1 registered
      ↳  98% — ✓ 12430 / ✗ 160
    ✗ multipart complete 200
      ↳  97% — ✓ 12267 / ✗ 322

    HTTP
    http_req_duration..............: avg=1.93s  min=738.1µs med=440.77ms max=46.58s p(90)=6.4s   p(95)=10.97s
      { expected_response:true }...: avg=1.93s  min=738.1µs med=437.62ms max=46.58s p(90)=6.43s  p(95)=10.99s
    http_req_failed................: 0.73% 570 out of 77329
    http_reqs......................: 77329 168.983269/s

    EXECUTION
    iteration_duration.............: avg=14.07s min=2.28s   med=13.57s   max=1m22s  p(90)=25.24s p(95)=29.17s
    iterations.....................: 12618 27.573496/s
    vus............................: 1     min=0            max=850
    vus_max........................: 850   min=850          max=850

    NETWORK
    data_received..................: 69 MB 151 kB/s
    data_sent......................: 14 GB 29 MB/s




running (7m37.6s), 000/850 VUs, 12618 complete and 2 interrupted iterations
default ✓ [======================================] 000/850 VUs  6m0s


```

# applying zap logging : 


# 8. 925 vu
```bash
924/950 worked
 
  █ TOTAL RESULTS

    checks_total.......: 96480  205.640433/s
    checks_succeeded...: 98.35% 94893 out of 96480
    checks_failed......: 1.64%  1587 out of 96480

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 13689 / ✗ 16
    ✗ multipart session 200
      ↳  99% — ✓ 13580 / ✗ 125
    ✓ part 1 s3 200
    ✗ part 1 registered
      ↳  96% — ✓ 13102 / ✗ 478
    ✗ multipart complete 200
      ↳  92% — ✓ 12607 / ✗ 968

    HTTP
    http_req_duration..............: avg=1.95s  min=0s   med=499.64ms max=30.75s p(90)=6.61s p(95)=10.27s
      { expected_response:true }...: avg=1.96s  min=0s   med=492.61ms max=30.75s p(90)=6.69s p(95)=10.4s
    http_req_failed................: 1.89% 1587 out of 83700
    http_reqs......................: 83700 178.400749/s

    EXECUTION
    iteration_duration.............: avg=14.14s min=2.3s med=13.25s   max=48.08s p(90)=26.1s p(95)=29.6s
    iterations.....................: 13700 29.2006/s
    vus............................: 22    min=0             max=925
    vus_max........................: 925   min=759           max=925

    NETWORK
    data_received..................: 73 MB 156 kB/s
    data_sent......................: 14 GB 30 MB/s


running (7m49.2s), 000/925 VUs, 13700 complete and 5 interrupted iterations
default ✓ [======================================] 000/925 VUs  6m0s

```

# 925 p2 - no authing

```bash

  █ TOTAL RESULTS

    checks_total.......: 84994  182.177038/s
    checks_succeeded...: 99.43% 84515 out of 84994
    checks_failed......: 0.56%  479 out of 84994

    ✓ signup status is 200
    ✓ has valid session token
    ✗ files 200
      ↳  98% — ✓ 13864 / ✗ 242
    ✗ multipart session 200
      ↳  98% — ✓ 13952 / ✗ 154
    ✗ part 1 s3 200
      ↳  99% — ✓ 13886 / ✗ 36
    ✓ part 1 registered
    ✗ multipart complete 200
      ↳  99% — ✓ 13864 / ✗ 47

    HTTP
    http_req_duration..............: avg=2.26s  min=0s      med=411.97ms max=1m0s   p(90)=7.97s p(95)=9.9s
      { expected_response:true }...: avg=2.17s  min=807.2µs med=413.18ms max=59.92s p(90)=7.9s  p(95)=9.83s
    http_req_failed................: 0.66% 479 out of 71813
    http_reqs......................: 71813 153.924744/s

    EXECUTION
    iteration_duration.............: avg=13.71s min=2s      med=12.27s   max=2m6s   p(90)=23.7s p(95)=28.27s
    iterations.....................: 14065 30.14707/s
    vus............................: 28    min=0            max=925
    vus_max........................: 925   min=738          max=925

    NETWORK
    data_received..................: 71 MB 151 kB/s
    data_sent......................: 15 GB 31 MB/s
```



# 9. 1000 VU 
```bash


# 6. 1000 vu
```bash
978/1000 users worked 
  
  █ TOTAL RESULTS

    checks_total.......: 96890  203.452379/s
    checks_succeeded...: 98.02% 94973 out of 96890
    checks_failed......: 1.97%  1917 out of 96890

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 13712 / ✗ 41
    ✗ multipart session 200
      ↳  99% — ✓ 13648 / ✗ 105
    ✓ part 1 s3 200
    ✗ part 1 registered
      ↳  95% — ✓ 13038 / ✗ 591
    ✗ multipart complete 200
      ↳  91% — ✓ 12435 / ✗ 1180

    HTTP
    http_req_duration..............: avg=2.14s  min=813.1µs med=509.23ms max=37.26s p(90)=6.57s  p(95)=11.97s
      { expected_response:true }...: avg=2.15s  min=813.1µs med=496.83ms max=37.26s p(90)=6.62s  p(95)=12.13s
    http_req_failed................: 2.27% 1917 out of 84137
    http_reqs......................: 84137 176.673267/s

    EXECUTION
    iteration_duration.............: avg=15.26s min=2.28s   med=13.46s   max=53.56s p(90)=27.91s p(95)=30.6s
    iterations.....................: 13720 28.809646/s
    vus............................: 2     min=0             max=1000
    vus_max........................: 1000  min=979           max=1000

    NETWORK
    data_received..................: 72 MB 151 kB/s
    data_sent......................: 14 GB 30 MB/s




running (7m56.2s), 0000/1000 VUs, 13720 complete and 33 interrupted iterations
default ✓ [======================================] 0000/1000 VUs  6m0s

```

