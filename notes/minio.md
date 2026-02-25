```go

func getS3Clients() (*s3.Client, *s3.Client) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// internal client for PutObject, CreateMultipartUpload etc
	// below lines for MinIO: if AWS_ENDP isnt theres its s3, if ter its miniIO
	// if endpoint is set, use it (MinIO in dev), otherwise use real S3

	internalEndpoint := os.Getenv("AWS_ENDPOINT")
	internalClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &internalEndpoint
		o.UsePathStyle = true
	})

	publicEndpoint := os.Getenv("AWS_PUBLIC_ENDPOINT")
	if publicEndpoint == "" {
		publicEndpoint = internalEndpoint
	}
	presignBaseClient := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = &publicEndpoint
		o.UsePathStyle = true
	})

	return internalClient, presignBaseClient
}

```

**Why two endpoints:**

MinIO runs inside Docker. Two different callers need to reach it:

- **Your Go services** (inside Docker) → use `http://minio:9000` — internal Docker hostname, works because they're on the same network
- **Your browser** (outside Docker) → use `http://localhost:9100` — host machine port mapping, works because Docker exposes MinIO on 9100

---

**Why it matters for presigned URLs specifically:**

A presigned URL is a signed URL — the signature is calculated using the endpoint hostname. So:

- If signed with `minio:9000` → browser can't resolve `minio` → fails
- If signed with `localhost:9100` → browser resolves it fine → works

The signature must match the host you actually hit. So the presign client must use the public endpoint from the start.

---

**Why we need two separate S3 clients:**

- `internalClient` → uses `minio:9000` → for backend ops like `PutObject`, `CreateMultipartUpload`, `CompleteMultipartUpload`
- `presignBaseClient` → uses `localhost:9100` → only for generating presigned URLs that the browser will use

In production `AWS_PUBLIC_ENDPOINT` is empty so it falls back to `AWS_ENDPOINT` which is your real S3 domain — resolves fine everywhere.