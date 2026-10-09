package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// ListFields retrieves all global fields of the organization.
func (c *Client) ListFields(ctx context.Context) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/fields", nil, nil)
}

// CreateField creates a new global issue field.
func (c *Client) CreateField(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/fields", nil, fields)
}

// GetField retrieves the parameters of a single global field.
func (c *Client) GetField(ctx context.Context, id string) (json.RawMessage, error) {
	return c.do(ctx, http.MethodGet, "/v3/fields/"+url.PathEscape(id), nil, nil)
}

// PatchField edits a global field's name and/or other parameters (category, order,
// description, visibility, possible values); version must match the field's current version.
func (c *Client) PatchField(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/fields/"+url.PathEscape(id), q, fields)
}

// CreateFieldCategory creates a new category for grouping global fields.
func (c *Client) CreateFieldCategory(ctx context.Context, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/fields/categories", nil, fields)
}

// PatchFieldCategory edits a field category; version must match the category's current version.
func (c *Client) PatchFieldCategory(ctx context.Context, id string, version int, fields map[string]any) (json.RawMessage, error) {
	q := url.Values{}
	q.Set("version", strconv.Itoa(version))
	return c.do(ctx, http.MethodPatch, "/v3/fields/categories/"+url.PathEscape(id), q, fields)
}
