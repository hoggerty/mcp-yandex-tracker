package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CreateTrigger creates a new trigger on a queue.
func (c *Client) CreateTrigger(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/triggers", nil, fields)
}

// ListTriggers lists the triggers configured on a queue.
func (c *Client) ListTriggers(ctx context.Context, queueKey string, perPage int, lastID string) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if lastID != "" {
		q.Set("id", lastID)
	}
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/triggers", q, nil)
}

// GetTrigger retrieves the parameters of a single trigger.
func (c *Client) GetTrigger(ctx context.Context, queueKey, triggerID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/triggers/"+url.PathEscape(triggerID), nil, nil)
}

// PatchTrigger edits a trigger; version must match the trigger's current version.
func (c *Client) PatchTrigger(ctx context.Context, queueKey, triggerID string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/queues/"+url.PathEscape(queueKey)+"/triggers/"+url.PathEscape(triggerID), q, fields)
}

// ListTriggerLogs lists the webhook invocation log of a trigger.
func (c *Client) ListTriggerLogs(ctx context.Context, queueKey, triggerID, issueID string, limit int, from, to string) (json.RawMessage, error) {
	q := url.Values{}
	if issueID != "" {
		q.Set("issueId", issueID)
	}
	if limit > 0 {
		q.Set("limit", strconv.Itoa(limit))
	}
	if from != "" {
		q.Set("from", from)
	}
	if to != "" {
		q.Set("to", to)
	}
	path := "/v3/queues/" + url.PathEscape(queueKey) + "/triggers/" + url.PathEscape(triggerID) + "/webhooks/log"
	return c.do(ctx, http.MethodGet, path, q, nil)
}
