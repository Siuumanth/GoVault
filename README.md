# GoVault

**Cloud-Native File Storage System with Resumable Uploads and Service Isolation**

---
# Dependencies: 
        
- **Language & Core Backend:** Go (Golang), net/http
- **Authentication & Security:** bcrypt, JWT
- **Resilience & Utilities:** Circuit breaker (custom / gobreaker), UUID package
- **Database:** PostgreSQL
- **Object Storage:** AWS S3, MinIO
- **Containerization & Deployment:** Docker, Docker Compose
- **Testing & Observability:** k6 (JavaScript), Prometheus, Grafana
- **Frontend**: React js

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
- Default Docker network for service communication

This design enforces **fault isolation**, **independent scalability**, and clean ownership of data.

---
## 3. API Gateway

The API Gateway acts as the single entry point into the system.
### Features
- Ordered middleware chaining
- Request ID generation and logging
- Security header injection
- CORS handling
- JWT authentication
- Fixed-window in-memory rate limiting
- Middleware adapters for handling and developing middlewares more efficiently.

The gateway handles all cross-cutting concerns, keeping internal services clean and focused on business logic.

---
## 4. Authentication Service

Authentication is fully stateless and token-based.
### Features
- Secure password hashing using bcrypt
- Stateless JWT issuance
- JWT validation at gateway level
- Token-based inter-service authorization

No session storage. Horizontal scaling is naturally supported.


---

## 5. Upload System

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

## Why Two Upload Mechanisms?

GoVault implements **both proxy-based chunked uploads and direct S3 multipart uploads** to explore architectural tradeoffs between control and scalability.

- **Proxy-based chunk upload** gives the backend full control over validation, state tracking, and metadata consistency, making it easier to enforce business rules and integrity.
    
- **Direct S3 multipart upload (via presigned URLs)** improves scalability by offloading large file transfer directly to object storage, reducing backend memory and bandwidth overhead.
    

Implementing both approaches allowed comparison of:
- Backend load vs scalability
- Control vs infrastructure delegation
- Operational complexity vs performance efficiency

This was a deliberate design decision to understand real-world upload architecture patterns used in production systems.


---

## 6. File Management System

### Features
- Metadata storage, retrieval and updates
- Soft delete mechanism
- Ownership validation
- File sharing (viewer/editor roles) and downloads
- Public/private toggle
- Three logical views for clients:
    - Owned files
    - Shared files
    - Shortcuts
supported by pagination.

This demonstrates proper relational modeling and access control design.

---
## 7. Code Architecture & Design
The internal structure follows production-grade layering:
- **Handler Layer** – HTTP parsing & response handling
- **Service Layer** – Business logic encapsulation
- **Repository Layer** – Database abstraction by dependency injection
### Design Principles Used
- Interface-driven development
- Dependency injection through interfaces
- Business logic separated from transport layer
- Repository abstraction to isolate persistence logic
- Middleware adapters for cross-cutting concerns

---
## 8. Database Design
- PostgreSQL per service
- UUID-based primary keys
- Indexed database on frequently read columns for performance
- Composite constraints for integrity
- Strict ownership relationships

Data integrity is enforced at database level — not just application level.
That’s important.

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
This is real-world DevOps exposure.

---
## 10. Observability & Testing
- Load testing using k6
- Prometheus metrics collection
- Grafana dashboard
- MinIO for local object storage testing

This shows operational thinking — not just feature building.

---
## 11. Frontend (MVP)
- Chunk slicing logic
- Upload session handling
- File views (owned, shared, shortcuts)
- Public access toggling

Frontend is functional, but backend is the core focus.

---
