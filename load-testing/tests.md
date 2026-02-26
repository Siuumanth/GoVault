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
INFO[0071] [SETUP] Done. Created 200 users.              source=console


  █ TOTAL RESULTS

    checks_total.......: 61720   142.648716/s
    checks_succeeded...: 100.00% 61720 out of 61720
    checks_failed......: 0.00%   0 out of 61720

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✓ files 200
    ✓ proxy session 200

    HTTP
    http_req_duration..............: avg=212.93ms min=563.4µs med=10.97ms max=2.93s p(90)=813.61ms p(95)=1.18s
      { expected_response:true }...: avg=212.93ms min=563.4µs med=10.97ms max=2.93s p(90)=813.61ms p(95)=1.18s
    http_req_failed................: 0.00% 0 out of 46540
    http_reqs......................: 46540 107.564343/s

    EXECUTION
    iteration_duration.............: avg=2.64s    min=2.12s   med=2.48s   max=4.95s p(90)=3.41s    p(95)=3.69s
    iterations.....................: 15380 35.546618/s
    vus............................: 3     min=0          max=199
    vus_max........................: 200   min=200        max=200

    NETWORK
    data_received..................: 19 MB 44 kB/s
    data_sent......................: 15 MB 35 kB/s




running (7m12.7s), 000/200 VUs, 15380 complete and 0 interrupted iterations
default ✓ [===================================] 000/200 VUs  6m0s
```