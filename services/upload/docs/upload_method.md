
### Frontend

* User selects file
* File is **manually sliced into fixed-size chunks**
* Each chunk is sent in a **separate HTTP request**
* Request body = **raw bytes**
* Metadata sent via headers:

  * `uploadId`
  * `chunkIndex`
* Failed chunks are retried
* Final “complete upload” request is sent

### Backend

* Create an `uploadId`
* Accept **one chunk per request**
* Stream chunk bytes directly to storage (disk/S3)
* Track uploaded chunks in DB
* Allow idempotent re-uploads
* On completion:

  * Validate all chunks
  * Assemble into final file
  * Clean up temp data

**One-line summary:**
> Chunked raw-body uploads with one request per chunk and resumable backend assembly.
