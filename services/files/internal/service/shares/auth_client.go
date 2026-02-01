package shares

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type AuthClient struct {
	BaseURL string
	client  *http.Client
}

func NewAuthClient(url string) *AuthClient {
	return &AuthClient{
		BaseURL: url,
		client:  &http.Client{Timeout: 2 * time.Second},
	}
}

// TODO: add correct URL
func (c *AuthClient) ResolveEmails(
	ctx context.Context,
	emails []string,
) (map[string]uuid.UUID, error) {

	url := fmt.Sprintf("%s/internal/resolve-users", c.BaseURL)

	// encode request body
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(emails); err != nil {
		return nil, fmt.Errorf("encode emails: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		&buf,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("auth request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(
			"auth service error %d: %s",
			resp.StatusCode,
			string(b),
		)
	}

	var result map[string]uuid.UUID
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return result, nil
}
