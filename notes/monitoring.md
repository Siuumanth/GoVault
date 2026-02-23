
When you define a `HistogramVec` in Go metrics.go:

```go
HttpRequestDuration = prometheus.NewHistogramVec(
    prometheus.HistogramOpts{
        Name:    "http_request_duration_seconds",
        Buckets: prometheus.DefBuckets,
    },
    []string{"method", "path"},
)
```

Prometheus automatically creates **3 metrics** from that one definition:

- `http_request_duration_seconds_bucket` — counts how many requests fell into each duration bucket (e.g. under 0.1s, under 0.5s, etc.)
- `http_request_duration_seconds_sum` — total sum of all durations
- `http_request_duration_seconds_count` — total number of observations

You never define these explicitly — Prometheus generates them for you under the hood when you call `.Observe()`.

`histogram_quantile(0.95, ...)` is then just a **PromQL function** that uses the bucket data to mathematically calculate the 95th percentile. Grafana/Prometheus does that math, not your Go code.

`prometheus.DefBuckets` specifically gives you these default buckets:
```
0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10 seconds
```

So as long as your request durations fall within those ranges (they will for normal HTTP), the percentile calculation works fine.