# GoVault

**A Microservices-Based Cloud File Storage System**

---
# Dependencies: 
        
- **Language & Core Backend:** Go (Golang), net/http
- **Authentication & Security:** bcrypt, JWT
- **Resilience & Utilities:** Circuit breaker (gobreaker), UUID package, Zap
- **Database:** PostgreSQL
- **Object Storage:** AWS S3, MinIO
- **Containerization & Deployment:** Docker, Docker Compose
- **Testing & Observability:** k6 (JavaScript), Prometheus, Grafana
- **Frontend**: React js

Built using the **Go (Golang) `net/http` standard library**. By avoiding heavy frameworks, the backend remains extremely lightweight, providing high execution speed and low memory overhead.

---
## 1. System Overview

GoVault is a cloud-native file storage platform designed with a **microservices architecture**, focusing on reliability, scalability, and production-grade backend design.

![](https://github.com/Siuumanth/GoVault/blob/main/images/system-arch.png?raw=true)

The system supports:
- proxy Chunked and resumable uploads
- S3 multipart uploads
- File ownership & sharing
- Public/private access control
- Rate limiting & authentication
- Service isolation with per-service databases

It emphasizes **clean architecture principles**, operational awareness, and deployment realism.

---
## 2. Architecture

### High-Level Design

#### Microservices:
- **API Gateway**
- **Auth Service**
- **Upload Service**
- **Files Service**
### Architectural Highlights
- Microservices-based design
- Clear service boundaries
- Separate PostgreSQL instance per service
- Interface-driven internal architecture
- REST-based inter-service communication
- Circuit breaker–protected service calls
- Default Docker network for inter service communication

This design enforces **fault isolation**, **independent scalability**, and clean ownership of data.

---

## 3. Code Architecture & Design
The internal structure follows production-grade layering:
- **Handler Layer** – HTTP parsing & response handling
- **Service Layer** – Business logic encapsulation
- **Repository Layer** – Database abstraction by dependency injection

![](https://github.com/Siuumanth/GoVault/blob/main/images/layers.png?raw=true)
#### **Low-Level Control & Custom Implementation**
Unlike many modern projects that rely on heavy frameworks or ORMs, GoVault is built from the ground up for maximum performance and transparency:

- **Standard Library Only:** Built using the native `net/http` package. No external web frameworks (like Gin or Echo) were used, ensuring a tiny binary size and zero "magic."

- **Raw SQL & No ORM:** All database interactions are written in raw SQL. This allows for optimized queries, precise indexing control, and avoids the performance overhead of an ORM.

- **Handcrafted Data Flow:** Every component was coded from scratch, including:
    - **Custom DTOs:** Strict Data Transfer Objects for API contracts.
    
    - **Domain Models:** Clear separation between database entities and business logic.

    - **Pointer Optimization:** Efficient memory management by passing pointers to avoid unnecessary data copying.
    
    - **Manual Mapping:** Total control over how data transforms between the database, service, and transport layers.

#### **Design Principles Used**
- **Interface-driven development:** High decouplability between layers.
- **Dependency injection:** Components are injected via interfaces, making unit testing easy.
- **Middleware Adapters:** Custom-built wrappers for cross-cutting concerns like logging and auth.

---
## 4. API Gateway

The API Gateway acts as the single entry point into the system.
### Features:
- Ordered middleware chaining
- Request ID generation and structured logging with `zap`
- Security header injection
- CORS handling
- JWT authentication
- Fixed-window in-memory rate limiting
- Middleware adapters for handling and developing middlewares more efficiently.

The gateway handles all cross-cutting concerns, keeping internal services clean and focused on business logic.

---
## 5. Authentication Service

Authentication is fully stateless and token-based.
### Features
- Secure password hashing using bcrypt
- Stateless JWT issuance
- JWT validation at gateway level
---
## 6. Upload System

This is the strongest technical component of the system, saving files in the cloud and supporting 2 types of uploads.
### A. Proxy-Based Chunked Upload

#### Flow:
Frontend → Backend → S3

#### Features:
- Custom 10MB chunk slicing on frontend
- Raw byte transmission
- Chunk tracking stored in database
- Server-side validation
- File assembly after all chunks received
- Upload session state management
- Resumable upload support
- Proper state tracking for each step

This approach gives **full control over validation and metadata consistency**, at the cost of some memory/network overhead.

---
### B. Direct S3 Multipart Upload (Presigned URLs)

- Backend creates multipart session
- Presigned URLs are generated in the backend where client uploads chunks.
- Client uploads directly to S3
- ETags stored in database
- Backend completes multipart upload
- Fully resumable

![](https://github.com/Siuumanth/GoVault/blob/main/images/upload.png?raw=true)
## Why Two Upload Mechanisms?

GoVault implements **both proxy-based chunked uploads and direct S3 multipart uploads** to explore architectural tradeoffs between control and scalability.

- **Proxy-based chunk upload** gives the backend full control over validation, state tracking, and metadata consistency, making it easier to enforce business rules and integrity.
- **Direct S3 multipart upload (via presigned URLs)** improves scalability by offloading large file transfer directly to object storage, reducing backend memory and bandwidth overhead.

Implementing both approaches allowed comparison of:
- Backend load vs scalability
- Control vs infrastructure delegation
- Operational complexity vs performance efficiency
and these metrics were validated by load tests.

This was a deliberate design decision to understand real-world upload architecture patterns and trade offs used in production systems.

---
## 7. File Management System

### Features
- Metadata storage, retrieval and updates (only name for now)
- Soft delete mechanism
- File sharing (viewer/editor roles) by email 
- File downloads and shortcuts
- Public/private toggle
- Three logical views for clients:
    - Owned files
    - Shared files
    - Shortcuts
supported by pagination.
---

---
## 8. Database Design
- PostgreSQL per service
- UUID-based identification and internal ids for faster joins and fetching
- Indexed database on frequently read columns for performance
- Composite constraints for integrity
- Strict ownership relationships

![](https://github.com/Siuumanth/GoVault/blob/main/images/schema.png?raw=true)

Data integrity is enforced at database level — not just application level.

---
## 9. Containerization & Deployment

### Features
- Multi-stage Docker builds, and static go binaries
- Docker Compose orchestration
- Dedicated volumes per service
- Dependency-based startup ordering
- Environment-driven configuration
- Smoke-tests health checks
- MinIO container for local S3 testing
All services communicate via default Docker networking.

---
## 10. Observability & Testing

- **Load Testing:** Performed concurrent user simulation using k6 to evaluate upload throughput, response latency, and service stability under stress.
- **Metrics Collection:** Integrated Prometheus to scrape service-level metrics for request counts, latency, and error rates.
- **Visualization:** Built a basic Grafana dashboard to monitor performance trends and bottlenecks in HTTP.
- **Structured Logging:** Implemented structured logging using Uber’s zap logger for performance and machine readability.
- **Local Object Storage Testing:** Used MinIO (S3-compatible) within Docker Compose for realistic local testing of upload flows and multipart behavior.

---
## 11. Frontend (MVP - similar look to Google Drive)
- Chunk/part slicing logic
- Upload session handling
- File views (owned, shared, shortcuts)
- Public access toggling and sharing between users

Frontend is functional, but backend is the core focus.

---
