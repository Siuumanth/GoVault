package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sony/gobreaker"
)

type AuthClient struct {
	BaseURL string
	client  *http.Client
	cb      *gobreaker.CircuitBreaker
}

func NewAuthClient(url string) *AuthClient {
	// Standard configuration for Auth:
	// If 5 requests fail in a row, or failure rate is high, stop trying for 30s.
	settings := gobreaker.Settings{
		Name:        "AuthService",
		MaxRequests: 3,               // Number of requests allowed in Half-Open state
		Interval:    5 * time.Minute, // How often to reset the counts in Closed state
		Timeout:     30 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip if we have more than 5 consecutive failures
			return counts.ConsecutiveFailures > 5
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf(
				"[CB] %s state change: %s → %s",
				name,
				from.String(),
				to.String(),
			)
		},
	}

	return &AuthClient{
		BaseURL: url,
		client:  &http.Client{Timeout: 2 * time.Second},
		cb:      gobreaker.NewCircuitBreaker(settings),
	}
}

func (c *AuthClient) ResolveEmails(
	ctx context.Context,
	emails []string,
) (map[string]uuid.UUID, error) {

	// Wrap the call in the Circuit Breaker
	body, err := c.cb.Execute(func() (interface{}, error) {
		url := fmt.Sprintf("%s/internal/resolve-users", c.BaseURL)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(emails); err != nil {
			return nil, err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			b, _ := io.ReadAll(resp.Body)
			return nil, fmt.Errorf("server error %d: %s", resp.StatusCode, string(b))
		}

		var result map[string]uuid.UUID
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, err
		}

		return result, nil
	})

	if err != nil {
		return nil, err
	}

	// Type assertion to get the data back from interface{}
	return body.(map[string]uuid.UUID), nil
}
