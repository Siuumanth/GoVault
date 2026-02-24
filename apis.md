
### **Gateway Entry Points (Port 9000)**

* `GET /health` – Check gateway status
* `POST /auth/signup` – Register new user
* `POST /auth/login` – Authenticate and get token


## Files: 
* `GET /api/files/f/{fileID}` – Get public file summary
* `GET /api/files/f/{fileID}/download` – Get public download link
* `GET /api/files/me/owned` – List user's files
* `GET /api/files/me/shared` – List shared files
* `GET /api/files/me/shortcuts` – List user's shortcuts
* `PATCH /api/files/f/{fileID}` – Rename owned file
* `DELETE /api/files/f/{fileID}` – Soft delete owned file
* `POST /api/files/f/{fileID}/copy` – Duplicate existing file
* `POST /api/files/f/{fileID}/shares` – Share with multiple users
* `GET /api/files/f/{fileID}/shares` – List file access permissions
* `PATCH /api/files/f/{fileID}/shares/{userID}` – Update user permissions
* `DELETE /api/files/f/{fileID}/shares/{userID}` – Revoke user access
* `POST /api/files/f/{fileID}/public` – Enable public access
* `DELETE /api/files/f/{fileID}/public` – Disable public access
* `POST /api/files/f/{fileID}/shortcut` – Create file bookmark
* `DELETE /api/files/f/{fileID}/shortcut` – Remove file bookmark

## Upload: 

* `POST /api/upload/proxy/session` – Initialize chunked upload
* `POST /api/upload/proxy/chunk` – Upload binary file part
* 
* `GET /api/upload/status` – Check upload progress
  
* `POST /api/upload/multipart/session` - Init multipart session
* `GET /api/upload/multipart/Presign` - Get al chunks presign urls
* `POST /api/upload/multipart/part` - Record ETag after a successful direct S3 upload

---

### **Internal Service Routes (Backend Only)**

**1. Upload Service  Files Service**
* `POST /internal/file` — Register finished S3 upload.

**2. Files Service  Auth Service**
* `POST /internal/internal/resolve-users` — Bulk convert emails to UUIDs.
