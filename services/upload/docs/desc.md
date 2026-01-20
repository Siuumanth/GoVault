upload chunk example format
{
    "upload_uuid": "1cb96d71-824e-4e13-bada-38358d55163b",
    "checksum": "eef6a32ab0c3851e385fb4292092c421641b9e73b1338472f3ec3f87f585220b",
    "chunk_bytes": "aGVsbG8gZ292YXVsdA=="

}
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




SHaring functions: 



Fileservice methods 

Registry or routes: 

MetaData(name)
Files
Sharing

Metadata:
- Update name 
- Fetch file metadata, like name and date upload n stuff 

Files:
- fetchSingleFile
- Fetchfiles by user id , applying pagination 
- Fetch shared files with user , applying pagination
- MakeFileCopy(file_id, actor_user_id)
   -  copy specific file from shared to own
   -  copy from own to own
  : both have same logic, take a file and make a copy of it , tats it 

Sharing:
- Add sharing email ids with their permissions
(initialy email ids are sent in a list)
but while updating or deleting ,its done individually, or shud adding also be done individually
- delete sharing recipient
- update sharing recipient
- fetch all shared recipients and access levels
- add file to public view
- delete file form public view

Shortcuts:
- Makeshortcut(appears in shortcuts tab)
- Delete shortcut, just one outcome
