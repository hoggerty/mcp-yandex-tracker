package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListStatuses lists the issue statuses available in the organization.
func (c *Client) ListStatuses(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/statuses", nil, nil)
}

// CreateStatus creates a new issue status.
func (c *Client) CreateStatus(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/statuses", nil, fields)
}

// EditStatus edits an issue status; version must match the status's current version.
func (c *Client) EditStatus(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/statuses/"+url.PathEscape(id), q, fields)
}
