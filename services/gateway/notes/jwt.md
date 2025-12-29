
---
# JWT (JSON Web Token) — Notes
## 1️⃣ What is a JWT
A **JWT** is a **signed string** that proves **who the user is** and **some claims about them**.
It is:
- **Stateless** (server doesn’t store session)
- **Self-contained** (identity + metadata inside token)
- **Tamper-proof** (signature verification)
JWT is just a **string** sent with every request.

---
## 2️⃣ Structure of a JWT

A JWT has **3 parts**, separated by dots:

```
header.payload.signature
```
### Header
- Algorithm used (`alg`: HS256)
- Token type (`typ`: JWT)
### Payload (Claims)
- User data (`user_id`, `email`)
- Standard claims:
    - `exp` (expiry)
    - `iat` (issued at)
    - `nbf` (not before)
### Signature
- Cryptographic proof that:
    - Header + payload were not modified
    - Token was issued by a trusted party
---
## 3️⃣ Where JWT is sent (your final decision)

You are using:

```
Authorization: Bearer <JWT>
```
So:
- JWT is **not a cookie**
- Browser does **not auto-send** it
- No CSRF issue
---
## 4️⃣ High-level JWT verification flow

What your **JWT middleware** does:

1. Extract JWT from `Authorization` header
2. Parse JWT string
3. Verify signature using secret key
4. Validate standard claims (`exp`, etc.)
5. Extract user claims
6. Attach claims to `context.Context`
7. Forward request downstream

JWT is verified **once**, in the **gateway**.

---
## 5️⃣ Extracting the Bearer token (code)

```go
authHeader := r.Header.Get("Authorization")
if authHeader == "" {
    return
}

const prefix = "Bearer "
if !strings.HasPrefix(authHeader, prefix) {
    return
}

tokenString := strings.TrimPrefix(authHeader, prefix)
```

Notes:
- Headers are always `string`
- `Bearer` **always lives in Authorization header**
---
## 6️⃣ What is `jwt.Parse`

Signature:

```go
func jwt.Parse(tokenString string, keyFunc jwt.Keyfunc) (*jwt.Token, error)
```

What it does:
- Decodes JWT (header + payload)
- Verifies signature
- Returns a **parsed token object** (`*jwt.Token`)
It **does not know** which key to trust — that’s why `keyFunc` exists.
---
## 7️⃣ What is `jwt.Token`

`jwt.Token` is an **in-memory parsed representation** of the JWT.

It contains:
- Header
- Claims
- Signing method (`alg`)
- `Valid` flag (true if verification passed)

JWT string → `jwt.Token` object

---
## 8️⃣ What is the “key”

The **key** is used to verify the signature.
For your case (HS256):
- Key = **JWT secret**
- Type = `[]byte`

Same secret is used to:
- Sign JWT
- Verify JWT
---
## 9️⃣ What is `keyFunc`
Type:

```go
type Keyfunc func(*jwt.Token) (interface{}, error)
```
### What it is
- A **callback function**
- Passed to `jwt.Parse`
- Called **by the JWT library**, not by you
### Why it exists
Because:
- Key choice depends on token metadata
- Library cannot decide which key is trusted
- Algorithm must be checked at runtime
---
## 10️⃣ What is a callback (in this context)

A **callback** is:
- A function passed as an argument
- Executed later by the library
In `jwt.Parse`:
- Library parses token
- Then calls your function:
    
    ```
    key := keyFunc(token)
    ```
- Uses returned key to verify signature
---
## 11️⃣ What happens inside `jwt.Parse` (exact order)
1. Decode JWT header + payload
2. Read algorithm from header
3. Call `keyFunc(token)`
4. Get verification key
5. Verify signature
6. Set `token.Valid`

Key is chosen **at verification time**, not before.

---
## 12️⃣ Why you can’t “just pass the key”

Because:
- Key selection happens **after parsing**
- Algorithm must be validated first
- Supports key rotation & multiple algorithms

So the library **requires a function**, not a raw key.

---
## 13️⃣ Extracting claims from parsed token (code)

### Using `MapClaims`

```go
token, err := jwt.Parse(tokenString, keyFunc)
if err != nil || !token.Valid {
    return
}

claims, ok := token.Claims.(jwt.MapClaims)
if !ok {
    return
}

userID := claims["user_id"]
email  := claims["email"]
```

---

## 14️⃣ What is this syntax?

```go
claims, ok := token.Claims.(jwt.MapClaims)
```

This is a **type assertion**.
Meaning:
- `token.Claims` is `interface{}`
- You are asserting it is actually `jwt.MapClaims`
- `ok` prevents panic if wrong type

Interfaces hide concrete types — assertions reveal them.

---
## 15️⃣ Final JWT middleware responsibility (locked)

JWT middleware must:
- Verify JWT once
- Convert JWT → Go data
- Store claims in `context.Context`
- Let downstream trust context, not token

---
## 16️⃣ Core mental model (final)

> **JWT is a signed string  
> Parsed once → verified once  
> Claims become trusted Go data  
> Passed downstream via context**




---



# Code:

Below are **clean, structured notes** exactly for **your understanding + interview recall**, following the format you prefer:

> **Overview → Explanation → Code (with comments) → Key takeaways**

---
# 🔐 JWT Middleware in Go (API Gateway)

## 1️⃣ Overview (What this code does)

This JWT middleware runs **inside the API Gateway** and is responsible for:

* Extracting the JWT from the `Authorization` header
* Verifying the token signature and expiry
* Validating and normalizing JWT claims
* Attaching authenticated user information to `context.Context`
* Forwarding the request to downstream handlers

The gateway is **stateless**, so:

* No database
* No session storage
* JWT is verified **once** at the edge

---

## 2️⃣ Why JWT Middleware belongs in the Gateway

* Gateway is the **single entry point**
* JWT verification is **CPU-only**, no I/O
* Downstream services should **trust identity**, not parse JWT again
* Identity is propagated using `context.Context`

This avoids:

* Duplicate JWT parsing
* Security bugs
* Tight coupling between services and auth logic

---

## 3️⃣ Context key design (important concept)

### Concept name: **Typed context keys**

```go
type contextKey string

const authCtxKey contextKey = "auth"
```

### Why this exists

* `context.Context` stores values as `interface{}`
* Using string keys can cause **collisions across packages**
* Typed keys prevent accidental overwrites

This is **production-safe** and **interview-relevant**.

---

## 4️⃣ JWT Middleware – Full Code (with detailed comments)

```go
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// 1. Read Authorization header
		// Expected format: "Authorization: Bearer <JWT>"
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 2. Validate Bearer scheme
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// 3. Extract raw JWT string
		// Strings are immutable in Go → must assign result
		tokenString := strings.TrimPrefix(authHeader, prefix)

		// 4. Load JWT secret
		// Shared secret used for HS256 verification
		jwtSecret := os.Getenv("JWT_SECRET")
		if jwtSecret == "" {
			http.Error(w, "Server misconfiguration", http.StatusInternalServerError)
			return
		}

		// 5. Parse and verify JWT
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

			// 5a. Validate signing algorithm
			// Prevents "alg: none" and algorithm confusion attacks
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// 5b. Return secret key for signature verification
			return []byte(jwtSecret), nil
		})

		// 6. Handle parsing / validation errors
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				http.Error(w, "Token expired", http.StatusUnauthorized)
			case errors.Is(err, jwt.ErrTokenMalformed):
				http.Error(w, "Malformed token", http.StatusUnauthorized)
			default:
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
			return
		}

		// 7. Ensure token is valid
		if !parsedToken.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// 8. Extract claims
		// Claims are stored as interface{} → type assertion needed
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// 9. Normalize claims into a typed auth object
		authCtx, err := utils.MapClaims(claims)
		if err != nil {
			http.Error(w, "Invalid token data", http.StatusUnauthorized)
			return
		}

		// 10. Attach auth context to request context
		ctx := context.WithValue(r.Context(), authCtxKey, authCtx)

		// 11. Forward request with enriched context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
```

---

## 5️⃣ Claim Normalization (MapClaims)

### Why normalization is needed

JWT claims:

* are `map[string]interface{}`
* contain untyped data
* cannot be trusted blindly

So we:

* validate required fields
* convert types
* produce a **clean, typed struct**

---

## 6️⃣ AuthContext struct

```go
type AuthContext struct {
	UserID   string
	Username string
	Role     string
	Expires  time.Time
}
```

This struct represents **verified identity**, not raw JWT data.

---

## 7️⃣ `MapClaims` implementation (with explanation)

```go
func MapClaims(claims jwt.MapClaims) (*AuthContext, error) {

	// Extract user ID
	uid, ok := claims["uid"].(string)
	if !ok {
		return nil, errors.New("invalid or missing uid")
	}

	// Extract username
	username, ok := claims["user"].(string)
	if !ok {
		return nil, errors.New("invalid or missing user")
	}

	// Extract role
	role, ok := claims["role"].(string)
	if !ok {
		return nil, errors.New("invalid or missing role")
	}

	// Extract expiry timestamp
	// JWT numeric dates are float64 (Unix seconds)
	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("invalid or missing exp")
	}

	return &AuthContext{
		UserID:   uid,
		Username: username,
		Role:     role,
		Expires:  time.Unix(int64(exp), 0),
	}, nil
}
```

---

## 8️⃣ How downstream services use this

```go
auth, ok := r.Context().Value(authCtxKey).(*AuthContext)
if !ok {
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return
}

fmt.Println(auth.UserID, auth.Role)
```

Downstream services:

* **Do not parse JWT**
* **Do not read headers**
* Trust gateway-provided identity

---

## 9️⃣ Key Takeaways (lock this in)

* `r.Context()` → per-request data + lifecycle
* JWT verified **once**, at the gateway
* Claims must be **validated and normalized**
* Context carries **identity**, not tokens
* Typed context keys prevent collisions
* This pattern is **production standard**

---

### ✅ Final mental model

> **JWT → verify → normalize → attach to context → forward request**

You now have **correct code**, **correct architecture**, and **correct mental model**.

If you want next:

* RBAC middleware using `auth.Role`
* how auth service should mint JWTs
* or how to unit-test this middleware properly

Just say.
