```
/metadata
/files
/sharing
/shortcuts
```
Each group maps cleanly to **one responsibility**.

---

# 2️⃣ Endpoint notes (one sentence each)

## 🔹 Metadata
- `UpdateFileName` - Updates only the file’s display name; does not affect ownership, storage, or access.
- `FetchFileMetadata`
Returns immutable and mutable metadata (name, size, created_at, owner, etc.) after permission check.

---

## 🔹 Files

- `FetchSingleFile`
Fetches a single file if the actor has ownership, share access, or public access.

- `FetchFilesByUser`
Lists files **owned by the user** with pagination and soft-delete filtering.

- `FetchSharedFiles`
Lists files shared **with** the user, paginated, excluding owned files.

- `MakeFileCopy`
Creates a new file owned by the actor by copying the underlying blob and metadata, regardless of source being shared or owned.

> One method. Permission decides **whether copy is allowed**, not how copy works.

---

## 🔹 Sharing

- `AddFileShares`
Adds multiple recipients to a file in a single operation for efficiency (initial share).

- `DeleteFileShare`
Removes one recipient’s access by deleting the corresponding share row.

- `UpdateFileShare`
Updates permission level for an existing recipient.

- `ListFileShares`
Returns all recipients and their access levels for a given file.

- `AddPublicAccess`
Makes a file publicly accessible by creating a public access rule.

- `RemovePublicAccess`
Revokes public access by deleting the public access rule.

---

## 🔹 Shortcuts

- `CreateShortcut`
Creates a reference to an existing file for quick access without affecting ownership.

- `DeleteShortcut`
Removes the reference only; the underlying file is untouched.





FILES
GET     /{fileID}
PATCH   /{fileID}
DELETE  /{fileID}
POST    /{fileID}/copy

LISTS
GET     /moved
GET     /shared

SHARES
POST    /{fileID}/shares
PATCH   /{fileID}/shares/{userID}
DELETE  /{fileID}/shares/{userID}
GET     /{fileID}/shares

PUBLIC ACCESS
POST    /{fileID}/public
DELETE  /{fileID}/public

SHORTCUTS
POST    /{fileID}/shortcut
DELETE  /shortcuts/{shortcutID}





FILES
GET /{fileID}
→ GetSingleFileSummary(ctx, fileID, actorUserID)

PATCH /{fileID}
→ UpdateFileName(ctx, UpdateFileNameInput)

DELETE /{fileID}
→ SoftDeleteFile(ctx, fileID, actorUserID)

POST /{fileID}/copy
→ MakeFileCopy(ctx, MakeFileCopyInput)

---

LISTS
GET /moved
→ ListOwnedFiles(ctx, ListOwnedFilesInput)

GET /shared
→ ListSharedFiles(ctx, ListSharedFilesInput)

---

SHARES
POST /{fileID}/shares
→ AddFileShares(ctx, AddFileSharesInput)

PATCH /{fileID}/shares/{userID}
→ UpdateFileShare(ctx, UpdateFileShareInput)

DELETE /{fileID}/shares/{userID}
→ RemoveFileShare(ctx, fileID, actorUserID, recipientUserID)

GET /{fileID}/shares
→ ListFileShares(ctx, fileID, actorUserID)

---

PUBLIC ACCESS
POST /{fileID}/public
→ AddPublicAccess(ctx, AddPublicAccessInput)

DELETE /{fileID}/public
→ RemovePublicAccess(ctx, RemovePublicAccessInput)

---

SHORTCUTS
POST /{fileID}/shortcut
→ CreateShortcut(ctx, CreateShortcutInput)

DELETE /shortcuts/{shortcutID}
→ DeleteShortcut(ctx, DeleteShortcutInput)

---

