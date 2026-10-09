package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListComponents lists all components in the organization.
func (c *Client) ListComponents(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/components", nil, nil)
}

// ListQueueComponents lists the components of a queue.
func (c *Client) ListQueueComponents(ctx context.Context, queueKey string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/queues/"+url.PathEscape(queueKey)+"/components", nil, nil)
}

// GetComponent retrieves the parameters of a single component.
func (c *Client) GetComponent(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/components/"+url.PathEscape(id), nil, nil)
}

// CreateComponent creates a new component on a queue.
func (c *Client) CreateComponent(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/components", nil, fields)
}

// EditComponent edits a component; version must match the component's current version.
func (c *Client) EditComponent(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/components/"+url.PathEscape(id), q, fields)
}

// DeleteComponent deletes a component.
func (c *Client) DeleteComponent(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodDelete, "/v3/components/"+url.PathEscape(id), nil, nil)
}

// GetComponentUserAccess retrieves a user's permissions for a component.
func (c *Client) GetComponentUserAccess(ctx context.Context, componentID, userID string) (json.RawMessage, error) {
	path := "/v3/components/" + url.PathEscape(componentID) + "/permissions/users/" + url.PathEscape(userID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// GetComponentGroupAccess retrieves a group's permissions for a component.
func (c *Client) GetComponentGroupAccess(ctx context.Context, componentID, groupID string) (json.RawMessage, error) {
	path := "/v3/components/" + url.PathEscape(componentID) + "/permissions/groups/" + url.PathEscape(groupID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}
