Perfect, now we can be **100% concrete** with *your exact code*.
No theory, no hand-waving.

---

## Short direct answer

👉 **`rw.WriteHeader` runs inside the API gateway, when `httputil.ReverseProxy` writes the downstream service’s response back to the client.**

It runs **here**, indirectly:

```go
proxy.ServeHTTP(w, r)
```

where `w` is actually **your wrapped `rw`**.

---

## Exact call chain in *your* code (line by line mentally)

### 1️⃣ Request enters gateway

```go
server := &http.Server{
    Handler: finalGateway,
}
```

`finalGateway` = **middleware chain + router**

---

### 2️⃣ Logging middleware wraps ResponseWriter

From `logging.mw.go`:

```go
rw := &responseWriter{
    ResponseWriter: w,
    status: http.StatusOK,
}

next.ServeHTTP(rw, r)
```

⚠️ **Important**
From this point onward, **everyone below sees `rw`**, not the original `w`.

---

### 3️⃣ Router dispatches to reverse proxy

Eventually this runs:

```go
r.Mount("/", http.StripPrefix("/auth", proxies.Auth))
```

`proxies.Auth` is:

```go
proxy.NewReverseProxy("http://localhost:9001")
```

---

### 4️⃣ Reverse proxy executes

From `reverse.go`:

```go
proxy.ServeHTTP(w, r)
```

But remember 👇
`w` === **`rw`**

---

### 5️⃣ What `httputil.ReverseProxy` does internally

Inside Go stdlib (simplified):

```go
resp, _ := transport.RoundTrip(outReq)

// COPY RESPONSE TO CLIENT
w.WriteHeader(resp.StatusCode)
io.Copy(w, resp.Body)
```

💥 **THIS LINE IS THE ANSWER**

```go
w.WriteHeader(resp.StatusCode)
```

Since `w` is `rw` →

```go
rw.WriteHeader(resp.StatusCode)
```

---

### 6️⃣ Your override runs

From `logging.mw.go`:

```go
func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}
```

✔ status captured
✔ real header still sent to client
✔ logger prints correct status

---

## Key misconception you had (now fixed)

> ❌ “The microservice should be calling `rw.WriteHeader`”

No.

### Reality:

* **Microservice** writes → **its own ResponseWriter**
* **Gateway HTTP client** receives bytes
* **Gateway reverse proxy** re-writes the response
* **Gateway ResponseWriter = rw**

They are **two different HTTP servers**, two different writers.

---

## One-sentence mental model (lock this in)

> **`rw.WriteHeader` runs when the gateway replays the downstream HTTP response to the client — not when the microservice generates it.**

---

## Bonus: why this design is clean

* Logging is **gateway-local**
* Services are **unaware**
* Works for **any backend language**
* Same mechanism used by:

  * NGINX
  * Envoy
  * Traefik
  * Go’s `ReverseProxy`

---

If you want next:

* I can walk you through **ReverseProxy source code**
* Or show **when logging breaks (Hijacker / WebSockets)**
* Or help you add **response size / upstream latency** to logs

But yeah — this question was 🔥 and very senior-level thinking.
