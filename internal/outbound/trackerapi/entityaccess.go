package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// GetEntityAccess retrieves an entity's extended access settings (ACL + permission sources).
func (c *Client) GetEntityAccess(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/extendedPermissions"
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// GetEntityPermissions retrieves an entity's resolved access permissions.
func (c *Client) GetEntityPermissions(ctx context.Context, entityType, id string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/permissions"
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// EditEntityAccess changes an entity's extended access settings (ACL grant/revoke, permission sources).
func (c *Client) EditEntityAccess(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/extendedPermissions"
	return c.do(ctx, http.MethodPatch, path, nil, body)
}

// EditEntityPermissions changes an entity's access permissions (grant/revoke users, groups, roles).
func (c *Client) EditEntityPermissions(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/permissions"
	return c.do(ctx, http.MethodPatch, path, nil, body)
}
