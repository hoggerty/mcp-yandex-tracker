package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListQueues lists the queues visible to the current user.
func (c *Client) ListQueues(ctx context.Context, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	return c.do(ctx, http.MethodGet, "/v3/queues", q, nil)
}

func (c *Client) GetQueue(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey), nil, nil)
}

// CreateQueue creates a new queue.
func (c *Client) CreateQueue(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues", nil, fields)
}

// DeleteQueue deletes a queue.
func (c *Client) DeleteQueue(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/queues/"+url.PathEscape(queueKey), nil, nil)
}

// RestoreQueue restores a previously deleted queue. Admin-only.
func (c *Client) RestoreQueue(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/_restore", nil, nil)
}

// GetQueueFields retrieves the required fields configured for a queue.
func (c *Client) GetQueueFields(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/fields", nil, nil)
}

// GetQueueTags retrieves the tags defined on a queue.
func (c *Client) GetQueueTags(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/tags", nil, nil)
}

// DeleteQueueTag removes a tag from a queue. Admin-only; the tag must be unused by any issue.
func (c *Client) DeleteQueueTag(ctx context.Context, queueKey, tag string) (json.RawMessage, error) {
	body := map[string]any{"tag": tag}
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/tags/_remove", nil, body)
}

// ManageQueueAccess grants or revokes queue-level permissions (create/write/read/grant) for users, groups, or roles.
func (c *Client) ManageQueueAccess(ctx context.Context, queueKey string, permissions map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/queues/"+url.PathEscape(queueKey)+"/permissions", nil, permissions)
}

// GetQueueUserAccess retrieves a user's permissions for a queue.
func (c *Client) GetQueueUserAccess(ctx context.Context, queueKey, userID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/permissions/users/"+url.PathEscape(userID), nil, nil)
}

// GetQueueGroupAccess retrieves a group's permissions for a queue.
func (c *Client) GetQueueGroupAccess(ctx context.Context, queueKey, groupID string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/permissions/groups/"+url.PathEscape(groupID), nil, nil)
}

// CreateLocalField creates a queue-local issue field.
func (c *Client) CreateLocalField(ctx context.Context, queueKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/queues/"+url.PathEscape(queueKey)+"/localFields", nil, fields)
}

// ListLocalFields lists a queue's local fields.
func (c *Client) ListLocalFields(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/localFields", nil, nil)
}

// GetLocalField retrieves a single local field of a queue.
func (c *Client) GetLocalField(ctx context.Context, queueKey, fieldKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/localFields/"+url.PathEscape(fieldKey), nil, nil)
}

// EditLocalField edits a queue-local issue field.
func (c *Client) EditLocalField(ctx context.Context, queueKey, fieldKey string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPatch, "/v3/queues/"+url.PathEscape(queueKey)+"/localFields/"+url.PathEscape(fieldKey), nil, fields)
}
