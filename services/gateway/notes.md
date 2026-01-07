Folder structure:
```bash

gateway/
├── cmd/
│   └── gateway/
│
├── internal/
│   ├── config/
│
│   ├── router/
│
│   ├── middleware/
│
│   └── proxy/
│
└── deploy/


```


# What main func does 

Great question — this is **exactly the right level of curiosity** about Go’s runtime.
I’ll explain this **step by step**, cleanly, and correct one small misconception on the way.

---

## 1️⃣ What happens when `main()` runs (ONCE)

You’re 100% correct here.

When you run the program:

```go
func main() {
    load env
    build deps
    build gateway
    build middleware chain
    http.ListenAndServe(":8080", finalHandler)
}
```

All of this happens **once** for the **entire lifetime of the process**.

After this point:

* No new gateways are created
* No new middleware chains are built
* No new routers are created

Everything is **fixed in memory**.

Think of it as building a **machine**, not running a job.

---

## 2️⃣ What `http.ListenAndServe` actually does

Internally (simplified):

* Starts a TCP listener on `:8080`
* Accepts incoming connections
* For **each incoming connection**, it does:

```go
go serveConnection(conn, handler)
```

👉 **Yes — Go starts a new goroutine** per connection (and often per request).

---

## 3️⃣ Important correction: NO copies are made

This is the most important part.

> ❌ Go does NOT copy your gateway or middleware
> ❌ Go does NOT clone handlers per request

### What actually happens

* There is **ONE** `finalHandler` in memory
* Every goroutine gets a **reference to the same handler**
* That handler is **read-only structure + shared state**

So:

```text
main goroutine
   |
   +--> finalHandler (one object in memory)
           |
           +--> goroutine #1 calls ServeHTTP
           +--> goroutine #2 calls ServeHTTP
           +--> goroutine #3 calls ServeHTTP
```

All goroutines call:

```go
finalHandler.ServeHTTP(w, r)
```

on the **same object**.

---

## 4️⃣ Why this is SAFE (this is crucial)

Because of **how you designed your code**.

### Middleware + handlers are:

* stateless (mostly)
* use local variables
* use request-scoped data (`*http.Request`, `Context`)
* avoid shared mutable state

Each request has:

* its own `http.Request`
* its own `ResponseWriter`
* its own `context.Context`

So concurrency is safe **by design**.

---

## 5️⃣ What IS shared across goroutines

Let’s be explicit.

### Shared (same memory)

* middleware structs
* router
* gateway
* config values
* env vars
* rate limiter state (⚠️ needs mutex)
* in-memory caches

### Per-request (NOT shared)

* `*http.Request`
* `http.ResponseWriter`
* `context.Context`
* local variables inside middleware
* stack variables

---

## 6️⃣ How middleware executes per request

For **each request goroutine**, this happens:

```
ServeHTTP
  ↓
logging middleware
  ↓
security headers
  ↓
cors
  ↓
jwt
  ↓
rate limiter
  ↓
router
  ↓
proxy
```

This is a **call stack**, not a chain of goroutines.

Each middleware runs **synchronously** inside the request goroutine.

---

## 7️⃣ Where goroutines DO matter for you

### 1️⃣ Rate limiter

You did the right thing:

```go
mu sync.Mutex
```

Because:

* multiple goroutines
* shared `map`
* concurrent writes

Without mutex → panic.

---

### 2️⃣ Background tasks

Example:

```go
go rl.resetVisitorCount()
```

This is a **separate goroutine**, running independently of requests.

---

## 8️⃣ Mental model you should lock in 🔒

> **main builds the server once
> Go spawns goroutines for requests
> handlers are shared
> request data is isolated
> shared state must be protected**

If you remember only this, you’re solid.

---

## 9️⃣ Why Go scales so well here

* goroutines are cheap
* stacks start small
* no thread-per-request overhead
* no handler copying
* no object cloning

That’s why Go is popular for gateways and servers.

---

## 10️⃣ Final one-line summary

> **Your gateway is built once.
> Every request runs in its own goroutine, calling the same handler with different request data.
> Only shared mutable state needs protection.**

