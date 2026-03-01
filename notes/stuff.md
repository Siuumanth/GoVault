
# db sql 

# Go `database/sql` — Core DB Functions

## 1️⃣ `db.Exec()`
Used when:
> You don’t expect any rows back.
Typical use:
- INSERT (without RETURNING)
- UPDATE
- DELETE

Example:

```go
result, err := db.Exec(
    "UPDATE upload_sessions SET status=$1 WHERE id=$2",
    "completed", sessionID,
)
```

What you can get:

```go
rowsAffected, _ := result.RowsAffected()
```

You **cannot** get row data from `Exec`.

---
## 2️⃣ `db.QueryRow()`

Used when:

> You expect **exactly one row**.
Perfect for:
- Get by ID
- COUNT(*)
- INSERT ... RETURNING

Example:

```go
err := db.QueryRow(
    "SELECT total_chunks FROM upload_sessions WHERE id=$1",
    sessionID,
).Scan(&total)
```

If no row:
```
err == sql.ErrNoRows
```

---
## 3️⃣ `db.Query()`

Used when:
> You expect **multiple rows**.

Example:

```go
rows, err := db.Query(
    "SELECT chunk_index FROM upload_chunks WHERE session_id=$1",
    sessionID,
)
defer rows.Close()

for rows.Next() {
    var idx int
    rows.Scan(&idx)
}
```

Always `rows.Close()`.

---
## 4️⃣ `INSERT ... RETURNING`

This is extremely important in GoVault.
Instead of:

```go
db.Exec("INSERT INTO files ...")
```

You do:

```go
db.QueryRow(
    "INSERT INTO files (...) VALUES (...) RETURNING id, created_at",
    ...
).Scan(&file.ID, &file.CreatedAt)
```

This:
- writes
- reads generated columns
- is atomic
---
## 5️⃣ `sql.ErrNoRows`

When using `QueryRow`, if no row exists:

```go
err == sql.ErrNoRows
```

You must handle this — otherwise it becomes a 500 error.

---

## 6️⃣ `*sql.DB` is a pool

`sql.DB` is:
- thread safe
- connection pooled
- long lived

You should create **one per service**, not per repository.

---
## 7️⃣ Transactions (`db.Begin()`)

When you need multiple queries to be atomic:

```go
tx, _ := db.Begin()
tx.Exec(...)
tx.Exec(...)
tx.Commit()
```

Use this for:
- assemble + insert file + mark session completed
---
## Summary Table

|Function|When to use|
|---|---|
|`Exec`|No rows expected|
|`QueryRow`|Exactly 1 row expected|
|`Query`|Many rows expected|
|`RETURNING`|Get generated columns|
|`ErrNoRows`|Row not found|
|`Begin`|Multi-step atomic ops|

---
## The GoVault rule
If your DB call:
- creates or reads **one thing** → `QueryRow`
- creates without needing fields → `Exec`
- reads **many** → `Query`









---

In GoVault (and in your Auth service), the rule is:

```
HTTP Handler  ⇄  DTO  ⇄  Service  ⇄  Model  ⇄  Repository
```

DTOs exist **only** at the **edge**:

* to parse JSON
* to shape responses
* to version APIs

Once the data enters the **service layer**, it should be:

* domain models
* primitives
* rich objects


Done during Upload testing:

- Debugged env collisions on Windows
- Fixed DSN parsing
- Designed clean DTOs
- Implemented chunked uploads correctly
- Handled filesystem paths + mkdir
- Thought about indexing, latency, and schema semantics