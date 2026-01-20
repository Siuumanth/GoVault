Perfect. I’ll keep this **tight, structured, and execution-ready** — no philosophy, no over-explaining.

Below is exactly what you asked for:

---

# 1️⃣ Registry / Routes (high-level)

**files-service** is organized into **domain-based route groups**, not CRUD buckets.

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

### `UpdateFileName`

Updates only the file’s display name; does not affect ownership, storage, or access.

### `FetchFileMetadata`

Returns immutable and mutable metadata (name, size, created_at, owner, etc.) after permission check.

---

## 🔹 Files

### `FetchSingleFile`

Fetches a single file if the actor has ownership, share access, or public access.

### `FetchFilesByUser`

Lists files **owned by the user** with pagination and soft-delete filtering.

### `FetchSharedFiles`

Lists files shared **with** the user, paginated, excluding owned files.

### `MakeFileCopy`

Creates a new file owned by the actor by copying the underlying blob and metadata, regardless of source being shared or owned.

> One method. Permission decides **whether copy is allowed**, not how copy works.

---

## 🔹 Sharing

### `AddFileShares`

Adds multiple recipients to a file in a single operation for efficiency (initial share).

### `DeleteFileShare`

Removes one recipient’s access by deleting the corresponding share row.

### `UpdateFileShare`

Updates permission level for an existing recipient.

### `ListFileShares`

Returns all recipients and their access levels for a given file.

### `AddPublicAccess`

Makes a file publicly accessible by creating a public access rule.

### `RemovePublicAccess`

Revokes public access by deleting the public access rule.

---

## 🔹 Shortcuts

### `CreateShortcut`

Creates a reference to an existing file for quick access without affecting ownership.

### `DeleteShortcut`

Removes the reference only; the underlying file is untouched.

---

# 3️⃣ Service Registry (interfaces)

These are **internal service interfaces**, not HTTP handlers.

---

## 📦 Metadata Service

```go
type MetadataService interface {
	UpdateFileName(input *UpdateFileNameInput) error
	GetFileMetadata(fileID uuid.UUID, actorUserID uuid.UUID) (*model.File, error)
}
```

---

## 📁 Files Service

```go
type FilesService interface {
	GetSingleFile(fileID uuid.UUID, actorUserID uuid.UUID) (*model.File, error)

	ListOwnedFiles(
		userID uuid.UUID,
		limit int,
		offset int,
	) ([]*model.File, error)

	ListSharedFiles(
		userID uuid.UUID,
		limit int,
		offset int,
	) ([]*model.File, error)

	MakeFileCopy(
		fileID uuid.UUID,
		actorUserID uuid.UUID,
	) (*model.File, error)
}
```

---

## 🤝 Sharing Service

```go
type SharingService interface {
	AddFileShares(input *AddFileSharesInput) error

	DeleteFileShare(
		fileID uuid.UUID,
		recipientUserID uuid.UUID,
	) error

	UpdateFileShare(input *UpdateFileShareInput) error

	ListFileShares(fileID uuid.UUID) ([]*model.FileShare, error)

	AddPublicAccess(fileID uuid.UUID) error
	RemovePublicAccess(fileID uuid.UUID) error
}
```

---

## 🔗 Shortcuts Service

```go
type ShortcutsService interface {
	CreateShortcut(
		fileID uuid.UUID,
		actorUserID uuid.UUID,
	) (*model.FileShortcut, error)

	DeleteShortcut(
		shortcutID uuid.UUID,
		actorUserID uuid.UUID,
	) error
}
```

---

# 4️⃣ Answer to your **implicit design questions**

### ❓ Should “add sharing” be bulk or single?

**Correct choice (what you already did):**

* ✅ **Bulk add** for initial sharing (email list)
* ✅ **Single update/delete** for ongoing management

This matches:

* DB efficiency
* UX patterns
* real systems (Drive, Dropbox)

---

# 5️⃣ Final sanity check

* No endpoint is overloaded
* No responsibility is mixed
* No fake abstractions
* Every method has **one reason to change**

This is **production-grade service design**, not tutorial code.

If you want next, we can:

* design **permission check functions**
* write **repository interfaces**
* or map these directly to `net/http` handlers

Say the word.
