package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CreateWorkflow creates a new workflow.
func (c *Client) CreateWorkflow(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/workflows", nil, fields)
}

// ListWorkflows lists all workflows in the organization (excluding deleted ones).
func (c *Client) ListWorkflows(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/workflows", nil, nil)
}

// GetQueueWorkflows retrieves the workflow assigned to each issue type in a queue.
func (c *Client) GetQueueWorkflows(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/workflows", nil, nil)
}

// GetWorkflow retrieves a single workflow by id.
func (c *Client) GetWorkflow(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/workflows/"+url.PathEscape(id), nil, nil)
}

// PatchWorkflow edits a workflow; version must match the workflow's current version.
func (c *Client) PatchWorkflow(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/workflows/"+url.PathEscape(id), q, fields)
}

// PatchWorkflowAction edits a single action of a workflow step; version must match the workflow's current version.
func (c *Client) PatchWorkflowAction(ctx context.Context, id, status, actionID string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	path := "/v3/workflows/" + url.PathEscape(id) + "/steps/" + url.PathEscape(status) + "/actions/" + url.PathEscape(actionID)
	return c.do(ctx, http.MethodPatch, path, q, fields)
}

// DeleteWorkflow deletes a workflow.
func (c *Client) DeleteWorkflow(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/workflows/"+url.PathEscape(id), nil, nil)
}
