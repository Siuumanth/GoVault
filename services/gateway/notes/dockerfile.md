Perfect questions. This is **core Docker brain** stuff. I’ll keep it **short, precise, and mechanical**.

---

## 1️⃣ Why order matters (Docker cache, line-by-line)

### Key rule

**Each Dockerfile instruction creates a cached layer.**
If a layer changes → **all layers after it are rebuilt**.

---

### ✅ GOOD version (why it’s good)

```dockerfile
COPY go.mod go.sum .
RUN go mod download
COPY . .
```

#### What happens:

1. `COPY go.mod go.sum .`

   * Cache depends **only** on go.mod/go.sum
2. `RUN go mod download`

   * Downloads deps → cached
3. `COPY . .`

   * Your app code (changes often)

#### Result:

* You change **code** → only last layer rebuilds
* Dependencies are **reused from cache**
* Build is **fast**

---

### ❌ BAD version (why it’s bad)

```dockerfile
COPY . .
RUN go mod download
```

#### What happens:

1. `COPY . .`

   * Copies **everything** (code, README, comments)
2. Any code change → this layer changes
3. Cache invalidated
4. `RUN go mod download` runs **again**

#### Result:

* Every small change → deps re-download
* Slow builds
* Wasted bandwidth

---

### One-line rule (memorize this)

> Put **slow + stable steps first**, fast-changing stuff last.

---

## 2️⃣ How does the builder layer get deleted?

Short answer: **it’s never part of the final image**.

### What Docker does internally

* Each `FROM` starts a **new image**
* `AS builder` = temporary build stage
* Final image only contains layers from **last FROM**

```dockerfile
FROM golang AS builder   ← temporary
FROM alpine              ← final image
```

When build finishes:

* Docker **does not tag** the builder image
* Untagged images are **garbage-collectable**
* They disappear automatically (or via `docker prune`)

👉 Builder stage is **discarded by design**, not manually deleted.

---

## 3️⃣ What is `--no-cache` (and `apk add --no-cache`)?

### Docker `--no-cache`

```bash
docker build --no-cache .
```

* Ignores **all cached layers**
* Forces a clean rebuild
* Used for debugging / fresh deps

---

### Alpine `apk add --no-cache`

```dockerfile
RUN apk add --no-cache ca-certificates
```

Normally:

* `apk` downloads package index
* Stores it on disk ❌

With `--no-cache`:

* Downloads index
* Installs packages
* **Deletes index immediately**

Result:

* Smaller image
* No junk files

---

## TL;DR (interview-ready)

* Docker caches **each instruction**
* Cache invalidation is **top-down**
* Copy dependency files before source code
* Multi-stage builds discard builder automatically
* `--no-cache` avoids storing temp package data

If you want next:

* how Docker decides cache hits internally
* difference between Docker cache vs BuildKit
* how layer hashes are computed

Say the word.
