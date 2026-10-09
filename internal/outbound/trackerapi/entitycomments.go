package trackerapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// AddEntityComment adds a comment to an entity.
func (c *Client) AddEntityComment(ctx context.Context, entityType, id string, fields map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/comments"
	return c.do(ctx, http.MethodPost, path, nil, fields)
}

// EditEntityComment edits an entity comment.
func (c *Client) EditEntityComment(ctx context.Context, entityType, id, commentID string, fields map[string]any) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/comments/" + url.PathEscape(commentID)
	return c.do(ctx, http.MethodPatch, path, nil, fields)
}

// ListEntityComments lists the comments of an entity. expand is optional (e.g. "all", "html", "attachments", "reactions").
func (c *Client) ListEntityComments(ctx context.Context, entityType, id, expand string) (json.RawMessage, error) {
	q := url.Values{}
	if expand != "" {
		q.Set("expand", expand)
	}
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/comments"
	return c.do(ctx, http.MethodGet, path, q, nil)
}

// GetEntityComment retrieves a single entity comment.
func (c *Client) GetEntityComment(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/comments/" + url.PathEscape(commentID)
	return c.do(ctx, http.MethodGet, path, nil, nil)
}

// DeleteEntityComment deletes an entity comment.
func (c *Client) DeleteEntityComment(ctx context.Context, entityType, id, commentID string) (json.RawMessage, error) {
	path := "/v3/entities/" + url.PathEscape(entityType) + "/" + url.PathEscape(id) + "/comments/" + url.PathEscape(commentID)
	return c.do(ctx, http.MethodDelete, path, nil, nil)
}
