
Below is a **clean, production-grade folder structure** for GoVault using **Go `net/http`**, **raw SQL (DAO pattern)**, and **separate microservices**.
This structure is intentionally **simple, scalable, and interview-friendly**.

---

# ✅ **Top-Level Project Structure**

```
govault/
│
├── services/
│   ├── gateway/
│   ├── upload/
│   ├── metadata/
│   ├── sharing/
│   └── preview/
│
├── pkg/
│   ├── logger/
│   ├── config/
│   ├── utils/
│   └── middleware/
│
├── deployments/
│   ├── docker/
│   ├── k8s/         (optional)
│   └── compose.yaml
│
└── README.md
```

Each microservice is **fully independent**, with its own `main.go`, HTTP handlers, DAO layer, and config.

---

# ✅ **Microservice Folder Structure Template**

Every service follows this exact blueprint:

```
services/<service-name>/
│
├── cmd/
│   └── <service-name>/
│       └── main.go
│
├── internal/
│   ├── http/
│   │   ├── handlers/       # All HTTP handlers
│   │   └── router.go
│   │
│   ├── dao/                # Pure SQL (DAO pattern)
│   │   ├── models.go
│   │   ├── queries.go
│   │   └── dao.go
│   │
│   ├── core/               # Business logic (“service layer”)
│   │   ├── upload.go
│   │   ├── sharing.go
│   │   └── metadata.go
│   │
│   ├── storage/            # S3 / MinIO adapters if needed
│   │   ├── s3.go
│   │   └── minio.go
│   │
│   ├── config/
│   │   └── config.go
│   │
│   └── app.go              # Wire up handlers, DB, router
│
└── go.mod
```

Below is the structure for *each* microservice.

---

# ⭐ **1. Gateway Service**

```
services/gateway/
├── cmd/gateway/main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   ├── auth.go
│   │   │   └── proxy.go
│   │   └── router.go
│   ├── config/config.go
│   └── app.go
└── go.mod
```

---

# ⭐ **2. Upload Service**

```
services/upload/
├── cmd/upload/main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   ├── chunk.go
│   │   │   └── session.go
│   │   └── router.go
│   │
│   ├── dao/               # raw SQL
│   │   ├── models.go
│   │   ├── dao.go
│   │   └── queries.go
│   │
│   ├── core/               # business logic
│   │   ├── chunk_service.go
│   │   └── upload_session.go
│   │
│   ├── storage/
│   │   ├── s3.go
│   │   ├── minio.go
│   │   └── interface.go
│   │
│   ├── config/config.go
│   └── app.go
└── go.mod
```

---

# ⭐ **3. Metadata Service**

```
services/metadata/
├── cmd/metadata/main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   ├── metadata.go
│   │   └── router.go
│   │
│   ├── dao/
│   │   ├── models.go
│   │   ├── dao.go
│   │   └── queries.go
│   │
│   ├── core/
│   │   └── metadata_service.go
│   │
│   ├── config/config.go
│   └── app.go
└── go.mod
```

---

# ⭐ **4. Sharing Service**

```
services/sharing/
├── cmd/sharing/main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   └── sharing.go
│   │   └── router.go
│   │
│   ├── dao/
│   │   ├── models.go
│   │   ├── dao.go
│   │   └── queries.go
│   │
│   ├── core/
│   │   ├── sharing.go
│   │   └── signed_urls.go
│   │
│   ├── config/config.go
│   └── app.go
└── go.mod
```

---

# ⭐ **5. Preview Service**

```
services/preview/
├── cmd/preview/main.go
├── internal/
│   ├── http/
│   │   ├── handlers/
│   │   │   ├── generate.go
│   │   │   └── status.go
│   │   └── router.go
│   │
│   ├── dao/
│   │   ├── models.go
│   │   └── dao.go
│   │
│   ├── core/
│   │   ├── preview_generator.go
│   │   └── worker.go
│   │
│   ├── storage/ (optional)
│   ├── config/config.go
│   └── app.go
└── go.mod
```

---

# 📦 **Shared Packages (`pkg/`)**

Reusable utilities for all services:

```
pkg/
├── config/        # env loading
│   └── env.go
├── logger/        # zap/logrus wrapper
│   └── logger.go
├── middleware/    # CORS, auth, rate-limit
│   └── auth.go
└── utils/         # helpers: UUID, hashing, etc.
    └── hash.go
```

---

# 📌 DAO Pattern (Raw SQL) — Folder Responsibilities

### **`dao/models.go`**

Structs that reflect DB rows.

### **`dao/queries.go`**

SQL statements:

```go
const CreateUploadSession = `
INSERT INTO upload_sessions (id, user_id, total_chunks)
VALUES ($1, $2, $3)
RETURNING id, created_at;
`
```

### **`dao/dao.go`**

Execute SQL & scan results.

```go
func (d *DAO) CreateUploadSession(ctx context.Context, s UploadSession) (*UploadSession, error) {
    row := d.db.QueryRowContext(ctx, CreateUploadSession, s.ID, s.UserID, s.TotalChunks)
    err := row.Scan(&s.ID, &s.CreatedAt)
    return &s, err
}
```

---

# 🧠 Why This Structure Works

✔ Completely isolates each microservice
✔ Perfect for Docker Compose + AWS
✔ DAO pattern keeps SQL clean and testable
✔ Handlers → Service Layer → DAO = clean architecture
✔ Easy to extend later (SQS, CloudFront, versioning, etc.)

---

# If you want, I can also generate for you:

✅ **The full `docker-compose.yaml`**
✅ **Directory skeleton as actual empty files**
✅ **Starter code for main.go for each service**
✅ **Upload-service chunked upload flow diagram**

Just tell me what you want next.
