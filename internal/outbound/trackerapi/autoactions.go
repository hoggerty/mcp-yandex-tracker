package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// CreateAutoAction creates a new auto action on a queue.
func (c *Client) CreateAutoAction(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/autoactions", nil, fields)
}

// GetAutoAction retrieves the parameters of an auto action.
func (c *Client) GetAutoAction(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/autoactions/"+url.PathEscape(autoActionID), nil, nil)
}

// ListAutoActionLogs lists the run history of an auto action.
func (c *Client) ListAutoActionLogs(ctx context.Context, queueKey, autoActionID string) (json.RawMessage, error) {
	path := "/v3/queues/" + url.PathEscape(queueKey) + "/autoactions/" + url.PathEscape(autoActionID) + "/logs"
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// GetAutoActionLog retrieves a single auto action run, including the issues it touched.
func (c *Client) GetAutoActionLog(ctx context.Context, queueKey, autoActionID, runID string) (json.RawMessage, error) {
	path := "/v3/queues/" + url.PathEscape(queueKey) + "/autoactions/" + url.PathEscape(autoActionID) + "/logs/" + url.PathEscape(runID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}
