package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// BulkMove moves many issues to another queue in a single background operation.
func (c *Client) BulkMove(ctx context.Context, targetQueue string, issues []string, values map[string]any, moveAllFields, initialStatus, notify bool) (json.RawMessage, error) {
	body := map[string]any{
		"queue":         targetQueue,
		"issues":        issues,
		"moveAllFields": moveAllFields,
		"initialStatus": initialStatus,
	}
	if len(values) > 0 {
		body["values"] = values
	}
	q := url.Values{}
	q.Set("notify", strconv.FormatBool(notify))
	return c.do(ctx, http.MethodPost, "/v3/bulkchange/_move", q, body)
}

// BulkUpdate sets field values on many issues in a single background operation.
func (c *Client) BulkUpdate(ctx context.Context, issues []string, values map[string]any, notify bool) (json.RawMessage, error) {
	body := map[string]any{"issues": issues, "values": values}
	q := url.Values{}
	q.Set("notify", strconv.FormatBool(notify))
	return c.do(ctx, http.MethodPost, "/v3/bulkchange/_update", q, body)
}

// BulkTransition executes a workflow transition on many issues in a single background operation.
func (c *Client) BulkTransition(ctx context.Context, transition string, issues []string, values map[string]any, notify bool) (json.RawMessage, error) {
	body := map[string]any{"transition": transition, "issues": issues}
	if len(values) > 0 {
		body["values"] = values
	}
	q := url.Values{}
	q.Set("notify", strconv.FormatBool(notify))
	return c.do(ctx, http.MethodPost, "/v3/bulkchange/_transition", q, body)
}

// GetBulkChangeStatus retrieves the status of a background bulk-change operation.
func (c *Client) GetBulkChangeStatus(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/bulkchange/"+url.PathEscape(id), nil, nil)
}

// GetBulkChangeErrors retrieves the issues that failed within a bulk-change operation, with error details.
func (c *Client) GetBulkChangeErrors(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/bulkchange/"+url.PathEscape(id)+"/issues", nil, nil)
}
