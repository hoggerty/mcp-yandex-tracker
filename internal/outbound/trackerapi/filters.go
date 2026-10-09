package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// CreateFilter saves a new saved issue filter.
func (c *Client) CreateFilter(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/filters", nil, fields)
}

// GetFilter retrieves a saved issue filter by id.
func (c *Client) GetFilter(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/filters/"+url.PathEscape(id), nil, nil)
}

// UpdateFilter replaces a saved issue filter's name/filter/sorts.
func (c *Client) UpdateFilter(ctx context.Context, id string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/filters/"+url.PathEscape(id), nil, fields)
}

// DeleteFilter deletes a saved issue filter. This endpoint is only documented under
// the legacy /v2 prefix; there is no /v3 equivalent.
func (c *Client) DeleteFilter(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v2/filters/"+url.PathEscape(id), nil, nil)
}
