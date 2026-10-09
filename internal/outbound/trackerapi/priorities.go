package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListPriorities lists the issue priorities available in the organization.
// If localized is false, names are returned in all supported languages.
func (c *Client) ListPriorities(ctx context.Context, localized bool) (json.RawMessage, error) {
	q := url.Values{}
	if !localized {
		q.Set("localized", "false")
	}
	return c.do(ctx, http.MethodGet, "/v3/priorities", q, nil)
}

// CreatePriority creates a new issue priority.
func (c *Client) CreatePriority(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/priorities", nil, fields)
}

// EditPriority edits an issue priority; version must match the priority's current version.
func (c *Client) EditPriority(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/priorities/"+url.PathEscape(id), q, fields)
}
