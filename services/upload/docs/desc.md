Short and clean.
I’ll give **only what’s needed**.

---

## 1️⃣ Valid HTTP endpoints (final MVP)

### 🔹 Create upload session

```
POST /api/upload/session
```

**Request body**

```json
{
  "file_name": "video.mp4",
  "file_size_bytes": 734003200
}
```

**Headers**

```
X-User-ID: <uuid>   // injected by gateway
```

**Response**

```json
{
  "upload_uuid": "550e8400-e29b-41d4-a716-446655440000",
  "total_chunks": 112
}
```

---

### 🔹 Upload a chunk

```
POST /api/upload/chunk
```

**Headers**

```
X-User-ID: <uuid>
Content-Type: application/octet-stream
X-Upload-UUID: <uuid>
X-Chunk-Index: <int>
X-Checksum: <sha256>
```

**Body**

```
raw binary chunk bytes
```

**Response**

```
200 OK
```

(You don’t need JSON here.)

---

### 🔹 Get upload status

```
GET /api/upload/status?upload_uuid=<uuid>
```

**Headers**

```
X-User-ID: <uuid>
```

**Response**

```json
{
  "upload_uuid": "550e8400-e29b-41d4-a716-446655440000",
  "status": "pending",
  "total_chunks": 112,
  "uploaded_count": 57
}
```
