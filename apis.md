
### **Gateway Entry Points (Port 9000)**

* `GET /health` – Check gateway status
* `POST /auth/signup` – Register new user
* `POST /auth/login` – Authenticate and get token


## Files: 
* `GET /api/files/{fileID}` – Get public file summary
* `GET /api/files/{fileID}/download` – Get public download link
* `GET /api/files/owned` – List user's files
* `GET /api/files/shared` – List shared files
* `GET /api/files/shortcuts` – List user's shortcuts
* `PATCH /api/files/{fileID}` – Rename owned file
* `DELETE /api/files/{fileID}` – Soft delete owned file
* `POST /api/files/{fileID}/copy` – Duplicate existing file
* `POST /api/files/{fileID}/shares` – Share with multiple users
* `GET /api/files/{fileID}/shares` – List file access permissions
* `PATCH /api/files/{fileID}/shares/{userID}` – Update user permissions
* `DELETE /api/files/{fileID}/shares/{userID}` – Revoke user access
* `POST /api/files/{fileID}/public` – Enable public access
* `DELETE /api/files/{fileID}/public` – Disable public access
* `POST /api/files/{fileID}/shortcut` – Create file bookmark
* `DELETE /api/files/{fileID}/shortcut` – Remove file bookmark

## Upload: 

* `POST /api/upload/session` – Initialize chunked upload
* `POST /api/upload/chunk` – Upload binary file part
* `GET /api/upload/status` – Check upload progress

---

### **Internal Service Routes (Backend Only)**

**1. Upload Service  Files Service**
* `POST /internal/file` — Register finished S3 upload.

**2. Files Service  Auth Service**
* `POST /internal/internal/resolve-users` — Bulk convert emails to UUIDs.
