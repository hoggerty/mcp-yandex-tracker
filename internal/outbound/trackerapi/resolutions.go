package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListResolutions lists the issue resolutions available in the organization.
func (c *Client) ListResolutions(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/resolutions", nil, nil)
}

// CreateResolution creates a new issue resolution.
func (c *Client) CreateResolution(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/resolutions", nil, fields)
}

// EditResolution edits an issue resolution; version must match the resolution's current version.
func (c *Client) EditResolution(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/resolutions/"+url.PathEscape(id), q, fields)
}
