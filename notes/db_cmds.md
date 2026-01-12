## Starting PSQL container:
- docker exec -it govault psql -U postgres

## PSQL cmds:

---

### 1️⃣ Create database

```sql
CREATE DATABASE govault;
```
---

### 2️⃣ List all databases

```sql
\l
```
or
```sql
\list
```
---
### 3️⃣ Connect / use a database

```sql
\c govault
```\

---

### 4️⃣ Show all tables (in current DB)

```sql
\dt
```

---

### 5️⃣ Show tables with schema (optional)

```sql
\dt+
```

---

### 6️⃣ Describe a table

```sql
\d table_name
```

---

### 7️⃣ Show current database

```sql
SELECT current_database();
```

---

### 8️⃣ Exit psql

```sql
\q
```

That’s all you need.
