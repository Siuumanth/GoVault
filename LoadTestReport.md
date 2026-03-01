Contents:

- introduction (what is this, tools used, context, overview of architecture)
and sayying strictly this test is based on locally tested , on windows ,  ryzen 5, 16 gb ram so there is wsl overhead also , so is less efifcient than actual cloud 

- things to know before hand , virtual users vs actual users, mapping between them, what is p95 n stuff, why errors can happen in http  , like tcp connections , db pool exhaust n stuff 
- actual load tests, starting from 200 vu


# **1. Introduction**

### **Project Overview**

This report covers the local performance testing of **GoVault**, a cloud-native file storage system.
- **Core Tech:** Built using the **Go (Golang) `net/http` standard library**. By avoiding heavy frameworks, the backend remains extremely lightweight, providing high execution speed and low memory overhead.
- **Objective:** The tests compare two ways of uploading files: **Proxy-Based** (through the backend) and **Direct-to-S3** (using MinIO).
### **System Architecture**

The system uses a **microservices** design. To test this, I simulated a full production stack using **Docker Compose**, running multiple containers simultaneously:

- **4 Go Services:** Gateway, Auth, Upload, and Files.
- **3 PostgreSQL Instances:** Dedicated database for each core service.
- **Object Storage:** **MinIO** (S3-compatible) for file storage.
- **Observability:** **Prometheus** and **Grafana** for tracking metrics.
### **Tools Used**
- **k6:** For simulating concurrent users and file uploads.
- **Prometheus & Grafana:** For monitoring system health and speed.
- **Zap:** For high-performance structured logging.
### **Testing Environment & Limits**

> **Strictly Local Test:** These tests were run on a single machine (**AMD Ryzen 5, 16GB RAM**) using **Windows with WSL2**.
> 
> Running **10+ containers** (4 services, 3 databases, MinIO, Prometheus, and Grafana) on one machine creates significant CPU and I/O overhead due to the WSL2 virtualization layer. This makes the system less efficient than a real cloud environment, but it accurately shows how the architecture handles stress



---


# 2. Core Concepts & Terminology
### **1. Virtual Users (VUs) vs. Actual Users**

In `k6`, we use **Virtual Users (VUs)**.
- **The Difference:** A real user might click a button and then spend 30 seconds reading a page. A **VU** is a script that executes a loop as fast as possible.
- **The Mapping:** 50 VUs do not equal 50 people; they might represent 500+ real-world users because the VUs never "stop to think"—they just keep hitting the API.
### **2. Percentiles (p95, p99) – Why Averages Lie**

If 99 people get their file in 1 second, but 1 person takes 60 seconds, the "Average" looks okay, but the experience for that one person is broken.

- **p95 (95th Percentile):** This means 95% of requests were _faster_ than this value. It’s the industry standard for "normal" performance.
    
- **p99 (99th Percentile):** This shows the "worst-case" scenarios. If p99 is very high, it usually means your system is hitting a bottleneck (like a slow DB query or GC pause).

### **3. Why HTTP Errors Happen Under Load**

When a system fails under stress, it’s usually not because the "code is wrong," but because a resource limit was hit:

- **TCP Connection Limits:** The OS only has a certain number of "slots" (sockets) to talk to the internet. If you open 1,000 connections at once, the OS might start dropping them.
    
- **DB Connection Pool Exhaustion:** Each service has a "pool" (e.g., 20 connections) to its PostgreSQL instance. If 50 VUs try to write metadata at the exact same millisecond, 30 of them have to wait. If they wait too long, the request times out (504 Gateway Timeout).

- **CPU Throttling:** As the number of Goroutines increases, the CPU spends more time "context switching" (deciding which task to do next) than actually doing the work. This causes a massive spike in response times.
    
- **Memory (RAM) Pressure:** In the **Proxy-Based** approach, the backend handles file chunks in memory. If too many large chunks are processed at once, the system might run out of RAM, triggering the **Go Garbage Collector (GC)** to run more frequently, which slows down the entire service.
    
- **Disk I/O Bottlenecks:** Since I am running **10+ containers** on one machine, all of them are fighting for the same Hard Drive/SSD "write" speed. If MinIO is writing a file while PostgreSQL is saving metadata, one of them will eventually lag.
- 
- **WSL2 Context Switching:** Since I am on Windows, the CPU spends a lot of time "swapping" between the Linux kernel (WSL) and the Windows kernel. This adds "jitter" to the results.
    
### **4. Throughput vs. Latency**

- **Throughput:** How many requests per second (RPS) the system can finish.
    
- **Latency:** How long a single request takes (the "delay"). Under high load, throughput usually plateaus (stays flat) while latency spikes (goes up) as requests start queuing.

### **5. Go Concurrency & Request Flow**

Go handles traffic differently than many other languages, which affects how it behaves under load:

- **Goroutines:** Instead of heavy OS threads, Go uses **Goroutines**. For every incoming request, the `net/http` server spawns a new goroutine. These are tiny (starting at ~2KB), allowing **GoVault** to handle thousands of concurrent connections without eating all your RAM.
    
- **API Gateway Mounting:** The Gateway doesn't just "pass" or "forward" data. It **mounts** the request to a route, wraps it in middleware (Auth, Rate Limiting, Logging), and then **forwards** it to the correct service, the request wait to finish on the goroutine (like async await in javascript).
    
- **The "Wait" Factor:** Even though Goroutines are fast, if the Gateway is waiting for a response from the Upload Service, that Goroutine stays active. If the backend is slow, these Goroutines pile up, increasing memory usage and latency.



---




# **3. Test Methodology & Adjustments**

### **How I Ran the Tests**

I initially tried **7 MB files**, but my system hit a resource wall. Since `k6` Virtual Users (VUs) are isolated, they can't share a memory buffer. With hundreds of VUs, the test runner and backend tried to hold gigabytes of raw data simultaneously, causing the **Go Garbage Collector** to struggle and performance to tank. I switched to **1 MB files** to focus on architectural limits like network and database connection pools.

Every test I ran caused my **CPU and Memory to fully explode** as the machine fought to manage 10+ containers and the load generator at once.
#### **Test Configuration**

Each test ran for **6 minutes** using a scaled stress strategy. I used a `SCALE` variable to adjust the intensity across these stages:

```javascript
export const stressOptions = {
  stages: [
    { duration: '1m', target: Math.round(250 * SCALE) },   // Initial push
    { duration: '1m', target: Math.round(500 * SCALE) },   // Breaking point search
    { duration: '1m', target: Math.round(750 * SCALE) },   // Extreme load
    { duration: '2m', target: Math.round(1000 * SCALE) },  // Peak stress
    { duration: '1m', target: 0 },                         // Recovery
  ],
  setupTimeout: '8m',
  http_req_duration: { max: 60000 }, // 60 seconds req timeout
};
```

The stages represent different points in time in the 6 minute window, where the load increases by time, as we seen in the target key's value.

---

### **The VU Lifecycle**

I designed the `k6` script to ensure every Virtual User (VU) is fully prepared before hitting the API. This mimics a realistic, high-pressure user journey.
#### **Global Setup (Runs Once)**
Before traffic starts, the `setup()` function registers **1,000 unique users**. It performs a **Signup** and **Login** for each, pre-baking a list of valid JWTs. This prevents registration bottlenecks from skewing the actual stress test results.
#### **The Continuous VU Loop**
Once the test starts, each VU is assigned an account and loops through these four phases:
- **Phase 1: Session Refresh (Auth)** The VU logs in again to ensure the token is active and the Auth service can handle concurrent requests.
    
- **Phase 2: Home page Simulation (Browse)** The VU calls the "Owned Files" endpoint, adding realistic **Database Read** pressure while the uploads happen.
    
- **Phase 3: The Heavy Lift (Upload)**
    - **Proxy:** Uploads a **1MB file**, calculates **SHA-256 checksums**, and streams chunks through the Go backend.
    - **Multipart:** Fetches **Presigned URLs**, puts chunks directly to **MinIO**, and sends **ETags** back to GoVault.
        
- **Phase 4: Assembly (Complete)** The VU sends the final `Complete` signal to finalize metadata—the stage where most database timeouts typically occur.

---
### **Understanding k6 Test Results**
#### How a K6 test result looks like

```bash
  █ TOTAL RESULTS
    checks_total.......: 96480  205.640433/s
    checks_succeeded...: 98.35% 94893 out of 96480
    checks_failed......: 1.64%  1587 out of 96480

    ✓ signup status is 200
    ✓ login 200
    ✓ has token
    ✗ files 200
      ↳  99% — ✓ 13689 / ✗ 16
    ✗ multipart session 200
      ↳  99% — ✓ 13580 / ✗ 125
    ✓ part 1 s3 200
    ✗ part 1 registered
      ↳  96% — ✓ 13102 / ✗ 478
    ✗ multipart complete 200
      ↳  92% — ✓ 12607 / ✗ 968
    HTTP
    http_req_duration..............: avg=1.95s  min=0s   med=499.64ms max=30.75s p(90)=6.61s p(95)=10.27s
      { expected_response:true }...: avg=1.96s  min=0s   med=492.61ms max=30.75s p(90)=6.69s p(95)=10.4s
    http_req_failed................: 1.89% 1587 out of 83700
    http_reqs......................: 83700 178.400749/s

    EXECUTION
    iteration_duration.............: avg=14.14s min=2.3s med=13.25s   max=48.08s p(90)=26.1s p(95)=29.6s
    iterations.....................: 13700 29.2006/s
    vus............................: 22    min=0             max=925
    vus_max........................: 925   min=759           max=925

    NETWORK
    data_received..................: 73 MB 156 kB/s
    data_sent......................: 14 GB 30 MB/s
  

running (7m49.2s), 000/925 VUs, 13700 complete and 5 interrupted iterations
default ✓ [======================================] 000/925 VUs  6m0s
```

Below is a summary from a **925 VU multipart upload test**:
> 
> - **Checks Succeeded:** 98.35% (The percentage of individual API steps that passed).
> - **HTTP Req Duration:** p(95)=10.27s (95% of requests finished within 10 seconds).
> - **Data Sent:** 14 GB (Total data pushed through the gateway during this 6-minute window).

#### **Explaining the Non-Obvious Metrics**

Looking at the results, there are a few things that aren't immediately clear:

- **The "Checks" Hierarchy:** Notice that **Signup** and **Login** usually have 100% success. The failures only start appearing at the **Multipart Complete** or **S3 Registration** steps. This tells me the Auth service is fine, but the **Upload Service** or **Database** is where the bottleneck sits.
    
- **Iteration Duration vs. Request Duration:** * `http_req_duration` is the time for a _single_ API call (e.g., uploading one part).
    
    - `iteration_duration` is the time it takes for one VU to complete the _entire flow_ (Login -> Init Upload -> Upload Parts -> Complete). An average of **14s** means a single user takes that long to finish the whole process under heavy load.
        
- **Expected Response (true) vs. Failed:** The `p(95)` for "expected_response: true" only tracks the timing for successful requests. If this number is much lower than the total `p(95)`, it means the "errors" are happening very fast (like an immediate 504 timeout), while the "successes" are the ones actually dragging the speed down.
    
- **Data Sent (30 MB/s):** This is the actual aggregate throughput my local machine was handling. Since I am running the test script, the Gateway, the Upload Service, and MinIO all on one disk, this **30 MB/s** represents a massive amount of internal I/O "fighting" for the same hardware.
    




---




# **4. Performance Metrics & Analysis**

In this section, I break down the core performance differences between the **Proxy-Based** and **Direct S3 Multipart** strategies. I used my CSV datasets to compare how the system handles increasing concurrency.

### **Live Monitoring: Grafana Dashboard**

I monitored the system in real-time using **Prometheus and Grafana** to track **HTTP metrics** like request duration and error rates. This was crucial for identifying the exact moment the system hit its limit. Across all test scenarios, the latency spikes in my dashboard made it clear that the **CPU and memory were pinned at 95%+**.

The dashboard allowed me to see the "hockey stick" curve in real-time—where latency stayed flat and then suddenly shot up. Seeing these HTTP metrics helped me understand that the bottlenecks were physical resource limits, specifically CPU saturation and I/O wait, rather than just code inefficiency. It proved that while the architecture was stable, the local hardware had reached its absolute ceiling.

### **The Visualization Strategy**

I have generated a separate dashboard for every single test run. However, since the **RPS (Requests Per Second) panels** look nearly identical across tests until the breaking point, I have included the most representative charts below to show the overall trends.

- ##### 400 VUs - proxy uploading 
![](https://github.com/Siuumanth/GoVault/blob/main/load-test-res/proxy-uploads/5-400vu-1mb-proxy.png?raw=true)

- ##### 750 VUs - proxy uploading (latency capped in the graph)
  
`"They gave us he number they had, I think the true number is much much higher"` moment  (If u got the reference then let me know :) )

![](https://github.com/Siuumanth/GoVault/blob/main/load-test-res/proxy-uploads/7-600vu.png?raw=true)

- ##### 400 VUs - multipart uploading
![](https://github.com/Siuumanth/GoVault/blob/main/load-test-res/multipart-uploads/1-400vu-1mb.png?raw=true)

- ##### 1000 VUs - multipart uploading 
![](https://github.com/Siuumanth/GoVault/blob/main/load-test-res/multipart-uploads/6-1000vu.png?raw=true)

### **The "Truth": k6 vs. Prometheus**

While I used **Grafana** for real-time monitoring, I relied on **k6 CSV exports** as the "absolute truth" for two main reasons:

- **Prometheus Bucket Smoothing:** Prometheus groups response times into "buckets" (e.g., 0-500ms). If a request takes 501ms or 999ms, they are lumped together, making the p95 graph less precise. **k6** records the exact millisecond of every single request.
    
- **"Invisible" Errors:** At peak stress, I saw many **Connection Refused** errors. Because these happened at the OS level (TCP port exhaustion), the requests never reached my Go code. Prometheus couldn't "see" them, but **k6** caught them because it tracks the request from the outside.


---

### **FINALLY!! (after an eternity of yapping) - The Comparative Results Table**
| **VU Count** | **Strategy**  | **Throughput (RPS)** | **Median Latency** | **p95 Latency** | **Error Rate** |
| ------------ | ------------- | -------------------- | ------------------ | --------------- | -------------- |
| **200**      | Proxy         | 142                  | 59ms               | 1.49s           | 0.00%          |
| **200**      | **Multipart** | **208**              | **51ms**           | **0.78s**       | 0.00%          |
|              |               |                      |                    |                 |                |
| **400**      | Proxy         | 140                  | 188ms              | 5.50s           | 0.00%          |
| **400**      | **Multipart** | **202**              | **172ms**          | **4.15s**       | 0.00%          |
|              |               |                      |                    |                 |                |
| **600**      | Proxy         | 126                  | 287ms              | 9.91s           | 0.01%          |
| **600**      | **Multipart** | **203**              | **233ms**          | **8.05s**       | 0.00%          |
|              |               |                      |                    |                 |                |
| 750          | Proxy         | 132                  | 373ms              | 10s             | 5%             |
| 750          | Multipart     | 202                  | 492ms              | 9.02s           | 0.00           |
|              |               |                      |                    |                 |                |
| **850**      | Proxy         | 155                  | 357ms              | 9.84s           | 14.0%          |
| **850**      | **Multipart** | **170**              | **437ms**          | **10.5s**       | **0.73%**      |
|              |               |                      |                    |                 |                |
|              |               |                      |                    |                 |                |
| 925          | **Multipart** | **178**              | **492ms**          | **10.4**        | **1.89%**      |
| 1000         | **Multipart** | **176**              | **496ms**          | **11.9**        | **2.27%**      |


---

# Visualizations for Easier Understanding:

## Proxy Uploads:

![](https://github.com/Siuumanth/GoVault/blob/main/images/proxy1.png?raw=true)

![](https://github.com/Siuumanth/GoVault/blob/main/images/proxy2.png?raw=true)


## S3 Multipart Uploads:

![](https://github.com/Siuumanth/GoVault/blob/main/images/mp1.png?raw=true)


![](https://github.com/Siuumanth/GoVault/blob/main/images/mp2.png?raw=true)

## Side by Side 

![](https://github.com/Siuumanth/GoVault/blob/main/images/bossTE.png?raw=true)

![](https://github.com/Siuumanth/GoVault/blob/main/images/bossTP.png?raw=true)





---
## **Deep Dive Analysis**

#### **The Efficiency Gap**

Right from the start (**200 VUs**), I observed that the **Direct S3 Multipart** strategy is significantly faster. It achieved **208 RPS** compared to the Proxy's **142 RPS**. Because the backend doesn't have to "touch" the file bytes in the Multipart flow, I saved a massive amount of CPU cycles, resulting in nearly half the p95 latency (**0.78s vs 1.49s**).

#### **The Breaking Point (The 750-850 VU Wall)**
The most revealing data point is the **Error Rate** at higher loads:

- **Proxy Failure:** At **850 VUs**, my Proxy strategy hit a wall with a **14% error rate**. The backend was so overwhelmed by managing file chunks in memory while simultaneously handling API logic that it began dropping requests.
    
- **Multipart Resilience:** At the same **850 VUs**, the Multipart strategy was still rock solid with an error rate of only **0.73%**. It successfully offloaded the heavy lifting to MinIO/S3, allowing the Go services to remain responsive.

One of my tests for 1000 VUs when running a couple other applications on my pc during the time of the test said this:
```bash
✓ signup status is 200
    ✗ login 200
      ↳  53% — ✓ 12921 / ✗ 11020
    ✗ has token
      ↳  53% — ✓ 12921 / ✗ 11020
    ✗ files 200
      ↳  53% — ✓ 12741 / ✗ 11184
    ✗ multipart session 200
      ↳  54% — ✓ 12923 / ✗ 10990
    ✓ part 1 s3 200
    ✓ part 1 registered
    ✓ multipart complete 200

    http_req_failed................: 29.49% 33194 out of 112547
```

---

#### **Throughput Saturation**

In my **Proxy** tests, I noticed the **Law of Diminishing Returns**. Even as I increased VUs from 200 to 500, the throughput stayed flat around **140 RPS**. This told me the system was "saturated"—adding more users didn't result in more work being done; it only increased the time users had to wait in the queue.

#### **Latency vs. Errors**

By **1000 VUs** in the Multipart test, the throughput dropped to **177 RPS** and the p95 latency climbed to **11.9s**. While the error rate stayed low (**2.27%**), the experience became sluggish. This confirms that while my architecture is stable, the local **WSL2 overhead** and single-disk I/O create a hard ceiling for performance that only a true cloud environment could resolve.

#### **The Latency Gap: Median vs. p95**

Across both architectures, I saw a massive divergence between Median and p95 latencies as load increased. 

- **Proxy Uploads:** At 400 VUs, the Median is **188ms**, but the p95 is **5.5 seconds**. This extreme gap is caused by **Head-of-Line Blocking**. Because the Go backend is busy buffering and checksumming chunks in memory, "unlucky" requests get stuck in the queue, making the experience terrible for the top 5% of users.
    
- **Multipart Uploads:** The gap is tighter and more predictable. At 400 VUs, the p95 is **4.15s**. Since the Go service isn't "touching" the file data, it doesn't get bogged down as easily. The eventual spike at 1000 VUs (**11.9s**) is mostly due to **TCP Connection Exhaustion**—the OS simply runs out of "slots" to handle new users.

#### **The "Bcrypt" CPU Trap**

Initially, I used **Bcrypt** for password hashing during the Auth phase. I quickly realized this was a massive bottleneck for a stress test.

- **The Problem:** Bcrypt is designed to be "slow" to prevent brute-force attacks, but at 500+ VUs, it was eating 70-80% of my CPU just to handle logins. This left almost no resources for the actual file uploading or database management.
    
- **The Fix:** For the purpose of this load test, I switched to **SHA-256**. While less secure for passwords in a real-world production app, it is much faster and allowed the CPU to focus on the architectural stress of the storage system rather than getting stuck on expensive math.

---





---

## **5. The Engineering Journey: Optimization & Lessons**

This project wasn't just about collecting numbers; it was about watching a distributed system break and re-engineering it in real-time. Here is what I learned through the "book of tests."

### **1. The WSL2 "Mount" Bottleneck**

Early on, I realized that **bind mounts** between Windows and WSL2 were a silent performance killer. The overhead of the Windows file system talking to the Linux kernel meant my databases and storage were lagging before the test even started. I moved everything to **Dedicated Named Volumes**, keeping all I/O strictly within the Linux environment, which stabilized my baseline latency.


### **2. The Database & Concurrency Wall**

As I pushed past 600 VUs, the system didn't just slow down—it started failing.

- **The Problem:** My Go services would hang because they ran out of **PostgreSQL connections**, and single heavy uploads would "lock" the DB, causing everything else to time out.
    
- **The Solution:** I learned that a production system needs a **Connection Pooler (like PgBouncer)** to manage DB traffic, **Horizontal Scaling** (adding more service replicas), and **Aggressive Timeouts** so one slow request doesn't kill the whole app.

### **3. The k6 Experience: Power vs. Overhead**

I was blown away by `k6`. Its engine is written in **Go**, which makes it incredibly fast, but the scripts are **JavaScript**, making it easy to write complex scenarios. However, I learned that `k6` is a double-edged sword:

- It spawns worker processes (VUs) that execute loops as fast as possible, which **explodes your local memory**.
    
- I hit a **"Metric Wall"** where the load generator itself became the bottleneck, proving that for massive tests, you need to distribute `k6` itself across multiple machines.

### **4. The "Zap" Factor: Structured Logging**

It wasn't just my imagination—switching to **Uber’s Zap Logger** actually decreased my error rates at **850+ VUs**. Standard logging often involves heavy string formatting and synchronous I/O. Zap is "zero-allocation" and extremely fast. By reducing the CPU time spent on logging, I freed up just enough cycles for the services to process actual requests instead of timing out.

### **5. Hardware Realities & Distributed Failure**

This project was a masterclass in **Resource Saturation**. I saw firsthand how fast computers are—managing thousands of processes and gigabytes of data every second—but I also saw their absolute limit.

- **Root Cause:** 99% of my failures weren't "bad code"; they were **resource exhaustion**. When CPU and RAM hit that 95% wall, the laws of physics take over, and the system begins to drift.
    
- **Theoretical Mastery:** I now have a good understanding of how distributed systems fail (TCP exhaustion, head-of-line blocking, disk I/O contention) and, more importantly, the architectural patterns needed to overcome them in a production-grade environment.
    
---

To wrap up the analysis, I want to talk about a major "ghost in the machine" moment I hit during a 850 VU stress test.

### **The Ghost Iteration Problem**

I found a confusing gap between my test logs and my actual storage. Even though **k6 reported 13,425 iterations finished**, my **Postgres** and **MinIO** storage only showed **624 unique user folders**.

Basically, over 200 of my 850 Virtual Users seemed to "finish" their work without actually leaving a single trace in the database or the cloud. It was like they never existed.

---
### **The Root Causes**

After digging into the architecture, I found three main reasons for this "Success Gap":

#### **1. The "Race to the Bottom" (Resource Starvation)**

Since my Go backend handles everything—Auth and File Metadata—850 VUs were all fighting for a tiny pool of **Database Connections**.

- The "fast" users (the ones who got a connection first) would finish multiple loops.
- The "slow" users (the ones waiting in line) were completely starved of resources.
    
    Because k6 counts an iteration as "complete" even if the script just finishes its loop after a failure, the total count looked high, but the actual data never made it to the disk.

#### **2. The OS Wall (TCP Exhaustion)**

At 850 VUs, my laptop hit its limit for **Open Files** and **Network Ports**.

- Many users reported "Completed" because the k6 script reached the end, but the network request was actually **Refused** before it even touched my Go code.
    
- Since the error happened at the **OS level** and not the **App level**, the system didn't even get a chance to log a 500 error. The requests just died silently in the background.

#### **3. The Metadata Sync Lag**

In the **Multipart** flow, there’s a tiny gap where MinIO gets the file, but the Go backend hasn't officially saved the "Success" info to Postgres yet.

- Under 95%+ CPU load, those database writes were simply timing out.
    
- MinIO would have the file parts, but because the "Complete" call never reached the database, the folder was never officially "finalized." The data was there, but the system didn't know it.
    
---

### **Theoretical Fix: The Redis "Shock Absorber"**

To solve the **"Ghost Iteration"** and **Database Starvation** issues, the next logical step in this architecture is adding a **Redis layer**.

- **Rate Limiting:** Instead of letting 850 VUs slam the PostgreSQL instance all at once, I could use Redis to implement a global **Rate Limiter**. This would "reject" excess traffic at the Gateway before it ever hits the expensive Database or Disk I/O.
    
- **Metadata Caching:** By caching user sessions and file metadata in Redis, I could reduce the number of "Reads" hitting Postgres by up to 90%. This would leave the database "Write" connections free to handle the final file assembly, preventing the timeouts that caused my data mismatch.




---


# **6. Conclusion**

The performance testing of GoVault proved one thing: **Architecture matters more than raw hardware.** Even on a single local machine with the overhead of WSL2, the shift from a Proxy-based model to a Direct S3 Multipart strategy transformed how the system handled pressure.

### **The Final Verdict**

- **The S3 Advantage:** By treating the Go backend as a **coordinator** (handling metadata) rather than a **data mover** (handling file bytes), I was able to maintain a steady throughput of **~175 RPS** even as the virtual user count climbed to 1,000.
    
- **Engineering Trade-offs:** I saw firsthand how every CPU cycle counts. Switching from **Bcrypt to SHA-256** for the test was a necessary move to keep the CPU free for the storage logic, and moving to **Zap Logging** proved that even "small" overheads can cause a system to tip over at 850+ VUs.
    
- **The Reality of Limits:** My tests hit a physical wall. The 95%+ CPU saturation and TCP port exhaustion showed that while the Go code was efficient, the local hardware reached its absolute limit.
    

---

### **Final Thoughts**

This marks the end of **3 months of hard grinding**. What started as a cloud-native file storage idea turned into a deep-dive into how distributed systems live, breathe, and eventually break. Seeing the "Ghost Iterations" and "Metric Walls" in real-time made me realize that building for scale isn't just about writing code—it's about managing resources and predicting failures.

Distributed systems are incredibly fun (and frustrating), and watching GoVault survive a 1,000 VU storm was the perfect way to wrap up this journey.