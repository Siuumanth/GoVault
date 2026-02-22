package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

type FileClient struct {
	BaseURL string
	client  *http.Client
	cb      *gobreaker.CircuitBreaker
}

func NewFileClient(url string) *FileClient {
	// Configure the Circuit Breaker settings
	settings := gobreaker.Settings{
		Name:        "FilesService",
		MaxRequests: 5,                // Max requests allowed when "half-open"
		Interval:    10 * time.Second, // Clear counts periodically when "closed"
		Timeout:     30 * time.Second, // How long to stay "open" before trying again
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			// Trip if failure rate > 50% after at least 10 requests
			if counts.Requests < 10 {
				return false
			}
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return failureRatio >= 0.5
		},

		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.Printf("CircuitBreaker[%s] state change: %s → %s\n", name, from.String(), to.String())
		},
	}
	return &FileClient{
		BaseURL: url,
		client:  &http.Client{Timeout: 10 * time.Second},
		cb:      gobreaker.NewCircuitBreaker(settings),
	}

}

// functions needed:
// 1. insert file model
func (c *FileClient) AddFile(
	ctx context.Context,
	file *CreateFileRequest,
) error {
	// Wrap the logic in the circuit breaker
	// CB is ur proxy so execute thru it

	_, err := c.cb.Execute(func() (interface{}, error) {
		url := fmt.Sprintf("%s/internal/file", c.BaseURL)

		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(file); err != nil {
			return nil, err
		}

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			return nil, err // This counts as a failure for the breaker
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 500 {
			return nil, fmt.Errorf("server error %d", resp.StatusCode) // counts as failure
		}

		// < 500 means request success
		return nil, nil
	})

	return err
}
