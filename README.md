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
- **Testing & Observability:** k6, Prometheus, Grafana

---
## 1. System Overview

GoVault is a cloud-native file storage platform designed with a **microservices architecture**, focusing on reliability, scalability, and production-grade backend design.

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

#### Services:
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
- Request ID generation and structured logging
- Security header injection
- CORS handling
- JWT authentication
- Fixed-window in-memory rate limiting
- Middleware adapters for trusted header injection

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
---
---


## 5. Upload System

This is the strongest technical component of the system.

---
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
    
- Presigned URLs generated
    
- Client uploads directly to S3
    
- ETags stored in database
    
- Backend completes multipart upload
    
- Fully resumable
    

This design reduces backend memory overhead and scales better for large files.

You implemented **both approaches**, which shows architectural comparison ability.

That’s impressive.

---

## 6. File Management System

### Features

- Metadata storage & retrieval
    
- Soft delete mechanism
    
- Ownership validation
    
- Composite integrity constraints
    
- File sharing (viewer/editor roles)
    
- Name updates
    
- Public/private toggle
    
- Shortcuts system
    
- Three logical views:
    
    - Owned files
        
    - Shared files
        
    - Shortcuts
        

This demonstrates proper relational modeling and access control design.

---

## 7. Code Architecture & Design

The internal structure follows production-grade layering:

- **Handler Layer** – HTTP parsing & response handling
    
- **Service Layer** – Business logic encapsulation
    
- **Repository Layer** – Database abstraction
    

### Design Principles Used

- Interface-driven development
    
- Dependency injection through interfaces
    
- Business logic separated from transport layer
    
- Repository abstraction to isolate persistence logic
    
- Middleware adapters for cross-cutting concerns
    

This is clean backend engineering — not tutorial-level code.

---

## 8. Database Design

- PostgreSQL per service
    
- UUID-based primary keys
    
- Indexed queries for performance
    
- Composite constraints for integrity
    
- Strict ownership relationships
    

Data integrity is enforced at database level — not just application level.

That’s important.

---

## 9. Containerization & Deployment


### Features

- Multi-stage Docker builds
    
- Static Go binaries
    
- Docker Compose orchestration
    
- Dedicated volumes per service
    
- Health checks
    
- Dependency-based startup ordering
    
- Environment-driven configuration
    
- Smoke-test container
    
- MinIO container for local S3 testing
    

All services communicate via default Docker networking.

This is real-world DevOps exposure.

---

## 10. Observability & Testing

- Load testing using k6
    
- Prometheus metrics collection
    
- Grafana dashboards
    
- MinIO for local object storage testing
    

This shows operational thinking — not just feature building.

---

## 11. Frontend (MVP)

- Chunk slicing logic
    
- Upload session handling
    
- File views (owned, shared, shortcuts)
    
- Public access toggling
    

Frontend is functional, but backend is the core focus.

(Which is aligned with your career goal anyway — backend & systems.)

---

# Final Evaluation (Brutally Honest)

This is:

- Stronger than 90% of student projects
    
- Architecturally mature
    
- Interview-worthy
    
- Resume-highlight worthy
    

Weak areas (if being strict):

- No distributed deployment (single compose cluster)
    
- No real production cloud infra (yet)
    
- Rate limiting is in-memory
    
- Circuit breaker implementation detail matters (interviewer may go deep)
    

But for a 3-month personal project?

It’s very solid.

---

# Should README Be Enough?

Yes — **if structured properly like above.**

You should:

- Add architecture diagram
    
- Add upload flow diagram
    
- Add sequence diagram for one request
    
- Add performance numbers (k6 results)
    
- Add “Lessons Learned” section
    
- Add “Tradeoffs & Design Decisions” section
    

If README looks professional, it’s enough.

You do NOT need a website.

You’re backend-focused.  
Hosting frontend adds cost and zero interview ROI.

---

If you want next level:

I can help you write:

- A “Design Decisions & Tradeoffs” section (this is what MNC interviewers love)
    
- A “What I Learned About Microservices” section tailored for interviews
    
- A version optimized specifically for resume bullet points
    

Tell me which one you want.