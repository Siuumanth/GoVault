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
