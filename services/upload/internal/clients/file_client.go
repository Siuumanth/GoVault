package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type FileClient struct {
	BaseURL string
	client  *http.Client
}

func NewAuthClient(url string) *FileClient {
	return &FileClient{
		BaseURL: url,
		// timeout 2 second
		client: &http.Client{Timeout: 2 * time.Second},
	}
}

// functions needed:
// 1. insert file model
func (c *FileClient) AddFile(
	ctx context.Context,
	file *CreateFileRequest,
) error {

	url := fmt.Sprintf("%s/internal/file", c.BaseURL)

	// encode request body
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		return fmt.Errorf("create file error: %w", err)
	}

	// post file
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		&buf,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf(
			"file service error %d: %s",
			resp.StatusCode,
			string(b),
		)
	}

	return nil
}
