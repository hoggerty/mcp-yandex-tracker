package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// ListMacros lists the macros configured on a queue.
func (c *Client) ListMacros(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/macros", nil, nil)
}

// GetMacro retrieves the parameters of a single macro.
func (c *Client) GetMacro(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/macros/"+url.PathEscape(macroID), nil, nil)
}

// CreateMacro creates a new macro on a queue.
func (c *Client) CreateMacro(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/macros", nil, fields)
}

// PatchMacro edits a macro.
func (c *Client) PatchMacro(ctx context.Context, queueKey, macroID string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/queues/"+url.PathEscape(queueKey)+"/macros/"+url.PathEscape(macroID), nil, fields)
}

// DeleteMacro deletes a macro from a queue.
func (c *Client) DeleteMacro(ctx context.Context, queueKey, macroID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/queues/"+url.PathEscape(queueKey)+"/macros/"+url.PathEscape(macroID), nil, nil)
}
