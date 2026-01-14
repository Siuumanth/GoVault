1. POST upload/session - CreateUploadSessionHandler
2. POST upload/chunk?idx=4 - UploadChunkHandler
3. GET /upload/status?upload_uuid=...  - GetUploadStatusHandler



1. POST upload/session
   Whenever a user wants to upload, they first cal this URL to create a session, then the chunks are sent 
   - handler: validate fields, call service
   - service:
       - calculate total chunks 
       - insert session row to uploadSession table 
       - get session ID , make folder for tat session
       - return UUID, 200, saying that session is created

```yaml
  /upload/session:
    post:
      summary: Create an upload session
      requestBody:
        required: true
        content:
          application/json:
            schema:
              type: object
              required:
                - file_name
                - file_size_bytes
              properties:
                file_name:
                  type: string
                file_size_bytes:
                  type: integer
                  format: int64
      responses:
        "200":
          description: Upload session created
          content:
            application/json:
              schema:
                type: object
                properties:
                  upload_uuid:
                    type: string
                    format: uuid
                  total_chunks:
                    type: integer

```
  

2. POST upload/chunk?idx=4
   After the session is created, frontend starts sending chunks in this format
   - handler: validate fields, bytes exist, checksum exists, sessionID
   - service:
       - calculate checksum of the bytes sent
       - compare with the checksum sent in req body
           - if equal, continue
           - if not equal, return 400/409
           - the frontend reads it to check what chunk to send again
       - save chunk in session folder
       - Saving data in DB:
           - Get upload session object
           - validate userID and get total chunks (u)
           - next, get total number of chunks uploaded in chunks table (v)
           - if u==v 
                - upload session complete 
                - trigger next s3 upload session 
                - update status 
                - return 200
           - if u!=v
                - just return 200 with status field     
           - else:
                - set status field to fail
                - delete chunk stored in tmp
                - send upload failed response   

### POST /upload

- validate
- compute totalChunks
- INSERT upload_sessions
- return `upload_uuid`

1. Resolve session by upload_uuid
2. If status != uploading → reject
3. Stream bytes → temp file
4. Compute checksum while streaming
5. If checksum mismatch → reject
6. INSERT upload_chunks (unique constraint)
7. COUNT(upload_chunks)
8. If count == total_chunks:
       - set status = assembling
       - assemble file
       - upload to S3
       - create files row
       - set status = completed
9. Return success
