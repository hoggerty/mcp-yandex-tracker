package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// CreateVersion creates a new version on a queue.
func (c *Client) CreateVersion(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/versions", nil, fields)
}

// ListQueueVersions lists the versions of a queue.
func (c *Client) ListQueueVersions(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/versions", nil, nil)
}

// GetVersion retrieves the parameters of a single version.
func (c *Client) GetVersion(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/versions/"+url.PathEscape(id), nil, nil)
}

// EditVersion edits a version.
func (c *Client) EditVersion(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/versions/"+url.PathEscape(id), nil, fields)
}

// DeleteVersion deletes a version.
func (c *Client) DeleteVersion(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/versions/"+url.PathEscape(id), nil, nil)
}
