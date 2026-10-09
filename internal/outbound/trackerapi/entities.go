package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// CreateEntity creates a new entity (project, portfolio, or goal).
func (c *Client) CreateEntity(ctx context.Context, entityType string, fields map[string]any) (json.RawMessage, error) {
	return c.do(ctx, http.MethodPost, "/v3/entities/"+url.PathEscape(entityType), nil, fields)
}

// GetEntity retrieves the parameters of a single entity. expand and fields are optional.
func (c *Client) GetEntity(ctx context.Context, entityType, id, expand, fields string) (json.RawMessage, error) {
	q := url.Values{}
	if expand != "" {
		q.Set("expand", expand)
	}
	if fields != "" {
		q.Set("fields", fields)
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id)
	return c.do(ctx, http.MethodGet, path, q, nil)
}

// UpdateEntity edits an entity.
func (c *Client) UpdateEntity(ctx context.Context, entityType, id string, body map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id)
	return c.do(ctx, http.MethodPatch, path, nil, body)
}

// DeleteEntity deletes an entity. If withBoard is true, its board is deleted too.
func (c *Client) DeleteEntity(ctx context.Context, entityType, id string, withBoard bool) (json.RawMessage, error) {
	q := url.Values{}
	if withBoard {
		q.Set("withBoard", "true")
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id)
	return c.do(ctx, http.MethodDelete, path, q, nil)
}

// SearchEntities lists entities of a given type matching a filter.
func (c *Client) SearchEntities(ctx context.Context, entityType string, body map[string]any, fields string, perPage, page int) (json.RawMessage, error) {
	q := url.Values{}
	if fields != "" {
		q.Set("fields", fields)
	}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/_search"
	return c.do(ctx, http.MethodPost, path, q, body)
}

// BulkUpdateEntities updates fields, comment, or links on multiple entities at once.
func (c *Client) BulkUpdateEntities(ctx context.Context, entityType string, body map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/bulkchange/_update"
	return c.do(ctx, http.MethodPost, path, nil, body)
}

// GetEntityEventsRelative retrieves an entity's change history, paginated relative to an event id.
func (c *Client) GetEntityEventsRelative(ctx context.Context, entityType, id string, perPage int, from, selected string) (json.RawMessage, error) {
	q := url.Values{}
	if perPage > 0 {
		q.Set("perPage", strconv.Itoa(perPage))
	}
	if from != "" {
		q.Set("from", from)
	}
	if selected != "" {
		q.Set("selected", selected)
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/events/_relative"
	return c.do(ctx, http.MethodGet, path, q, nil)
}
